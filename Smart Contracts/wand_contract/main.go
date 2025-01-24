package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Material representa um material utilizado na criação de varinhas
type Material struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity_material"`
	Origin   string `json:"origin"`
}

// Wand representa uma varinha feita de materiais
type Wand struct {
	ID        string     `json:"id"`
	Materials []Material `json:"materials"`
	Quantity  int        `json:"quantity_wand"`
}

// StudioContract define o contrato do chaincode
type StudioContract struct {
	contractapi.Contract
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(StudioContract))
	if err != nil {
		fmt.Printf("Error creating StudioContract chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting StudioContract chaincode: %s", err.Error())
	}
}

// Init inicializa o ledger com dados opcionais ou configurações iniciais
func (sc *StudioContract) Init(ctx contractapi.TransactionContextInterface) error {
	// A lógica de inicialização pode ser personalizada conforme necessário
	fmt.Println("Chaincode StudioContract foi inicializado")
	return nil
}

// CheckAccess verifica se a transação é realizada pela organização apropriada
func (sc *StudioContract) CheckAccess(ctx contractapi.TransactionContextInterface, requiredOrg string) error {
	if requiredOrg == "" {
		return fmt.Errorf("required organization is not specified")
	}

	clientOrg, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get client organization: %v", err)
	}

	if clientOrg != requiredOrg {
		return fmt.Errorf("access denied: client organization %s does not match required organization %s", clientOrg, requiredOrg)
	}

	return nil
}

// Invoke roteia a função chamada para a transação correta
func (sc *StudioContract) Invoke(ctx contractapi.TransactionContextInterface) error {
	// Obtém o nome da função sendo chamada
	fn, args := ctx.GetStub().GetFunctionAndParameters()

	// Roteia a função chamada para a transação correta
	switch fn {
	case "AddMaterial":
		if len(args) < 1 {
			return fmt.Errorf("AddMaterial requires materialJSON arguments")
		}
		return sc.AddMaterial(ctx, args[0])

	case "CreateWand":
		if len(args) < 1 {
			return fmt.Errorf("CreateWand requires wandJSON arguments")
		}
		return sc.CreateWand(ctx, args[0])

	case "QueryMaterial":
		if len(args) < 1 {
			return fmt.Errorf("QueryMaterial requires materialID argument")
		}
		material, err := sc.QueryMaterial(ctx, args[0])
		if err != nil {
			return err
		}
		// Converte o material para JSON e retorna o resultado
		materialJSON, err := json.Marshal(material)
		if err != nil {
			return fmt.Errorf("failed to marshal material: %v", err)
		}
		return ctx.GetStub().SetEvent("QueryMaterialResult", materialJSON)

	case "QueryWand":
		if len(args) < 1 {
			return fmt.Errorf("QueryWand requires wandID argument")
		}
		wand, err := sc.QueryWand(ctx, args[0])
		if err != nil {
			return err
		}
		// Converte a varinha para JSON e retorna o resultado
		wandJSON, err := json.Marshal(wand)
		if err != nil {
			return fmt.Errorf("failed to marshal wand: %v", err)
		}
		return ctx.GetStub().SetEvent("QueryWandResult", wandJSON)
	
	case "GetAllMaterials":
		materials, err := sc.GetAllMaterials(ctx)
		if err != nil {
			return fmt.Errorf("Failed to get materials: %v", err)
		}
	
		// Converte os materiais para JSON e retorna o resultado
		materialsJSON, err := json.Marshal(materials)
		if err != nil {
			return fmt.Errorf("Failed to marshal materials: %v", err)
		}
		return ctx.GetStub().SetEvent("GetAllMaterialsResult", materialsJSON)
	
	case "GetAllWands":
		wands, err := sc.GetAllWands(ctx)
		if err != nil {
			return fmt.Errorf("Failed to get wands: %v", err)
		}
	
		// Converte as varinhas para JSON e retorna o resultado
		wandsJSON, err := json.Marshal(wands)
		if err != nil {
			return fmt.Errorf("Failed to marshal wands: %v", err)
		}
		return ctx.GetStub().SetEvent("GetAllWandsResult", wandsJSON)	

	default:
		return fmt.Errorf("Unknown function: %s", fn)
	}
}

// AddMaterial adiciona um novo material ao ledger
func (sc *StudioContract) AddMaterial(ctx contractapi.TransactionContextInterface, materialJSON string) error {

	// Verifica se o usuário tem permissão para adicionar materiais, baseando-se no nome da organização.
	if err := sc.CheckAccess(ctx, "orgprodutor-example-com"); err != nil {
		return err // Retorna erro se o acesso for negado
	}

	// Deserializa o JSON fornecido para um objeto Material.
	var material Material
	err := json.Unmarshal([]byte(materialJSON), &material)
	if err != nil {
		return fmt.Errorf("failed to unmarshal material data: %v", err) // Retorna erro se houver falha na deserialização
	}

	// Verifica se todos os campos obrigatórios estão preenchidos e válidos (ID, Name, Quantity > 0, Origin).
	if material.ID == "" || material.Name == "" || material.Quantity <= 0 || material.Origin == "" {
		return fmt.Errorf("invalid material data: all fields (ID, Name, Quantity > 0, Origin) must be filled correctly")
	}

	// Define o tipo do objeto (usado para criar uma chave composta)
	objectType := "material"
	compositeKey, err := ctx.GetStub().CreateCompositeKey(objectType, []string{material.ID})
	if err != nil {
		return fmt.Errorf("failed to create composite key: %v", err) // Retorna erro se a criação da chave composta falhar
	}

	// Verifica se o material com a mesma ID já existe no ledger
	existingMaterialBytes, err := ctx.GetStub().GetState(compositeKey)
	if err != nil {
		return fmt.Errorf("failed to get material from ledger: %v", err) // Retorna erro se não conseguir acessar o ledger
	}

	// Se o material já existe no ledger
	if existingMaterialBytes != nil {
		var existingMaterial Material
		err = json.Unmarshal(existingMaterialBytes, &existingMaterial)
		if err != nil {
			return fmt.Errorf("failed to unmarshal existing material data: %v", err) // Retorna erro se houver falha na deserialização dos dados existentes
		}

		// Impede a atualização se o ID já existir, mas o Nome ou a Origem forem diferentes
		if existingMaterial.Name != material.Name || existingMaterial.Origin != material.Origin {
			return fmt.Errorf("cannot add material: material with the same ID but different Name or Origin already exists")
		}

		// Incrementa a quantidade do material existente se ele já tiver o mesmo ID, Nome e Origem
		existingMaterial.Quantity += material.Quantity
		materialAsBytes, err := json.Marshal(existingMaterial)
		if err != nil {
			return fmt.Errorf("failed to marshal updated material data: %v", err) // Retorna erro se a serialização falhar
		}

		// Atualiza o estado do material no ledger com a nova quantidade
		return ctx.GetStub().PutState(compositeKey, materialAsBytes)
	} else {
		// Se o material não existe, serializa e insere o novo material no ledger
		materialAsBytes, err := json.Marshal(material)
		if err != nil {
			return fmt.Errorf("failed to marshal material data: %v", err) // Retorna erro se a serialização falhar
		}

		// Adiciona o novo material ao ledger
		return ctx.GetStub().PutState(compositeKey, materialAsBytes)
	}
}

// CreateWand cria uma nova varinha com base em materiais
func (sc *StudioContract) CreateWand(ctx contractapi.TransactionContextInterface, wandJSON string) error {
	// Verifica se o chamador pertence à organização OrgOlivaras
	if err := sc.CheckAccess(ctx, "orgolivaras-example-com"); err != nil {
		return err
	}

	// Deserializa a varinha a partir do JSON fornecido
	var wand Wand
	err := json.Unmarshal([]byte(wandJSON), &wand)
	if err != nil {
		return fmt.Errorf("failed to unmarshal wand data: %v", err)
	}

	// Validação dos dados obrigatórios da varinha
	if wand.ID == "" || wand.Quantity <= 0 || len(wand.Materials) == 0 {
		return fmt.Errorf("invalid wand data: all fields (ID, Quantity > 0, Materials) must be filled correctly")
	}

	// Verifica se os materiais necessários estão disponíveis e subtrai a quantidade
	for _, materialUsed := range wand.Materials {
		// Cria a chave composta para o material
		objectType := "material"
		compositeKey, err := ctx.GetStub().CreateCompositeKey(objectType, []string{materialUsed.ID})
		if err != nil {
			return fmt.Errorf("failed to create composite key for material %s: %v", materialUsed.ID, err)
		}

		// Obtém o material do ledger usando a chave composta
		materialBytes, err := ctx.GetStub().GetState(compositeKey)
		if err != nil {
			return fmt.Errorf("failed to get material %s from ledger: %v", materialUsed.ID, err)
		}

		if materialBytes == nil {
			return fmt.Errorf("material %s not found", materialUsed.ID)
		}

		// Desserializa o material
		var material Material
		err = json.Unmarshal(materialBytes, &material)
		if err != nil {
			return fmt.Errorf("failed to unmarshal material data: %v", err)
		}

		// Multiplica a quantidade necessária pelo número de varinhas a serem criadas
		requiredQuantity := materialUsed.Quantity * wand.Quantity

		// Verifica se há quantidade suficiente do material
		if material.Quantity < requiredQuantity {
			return fmt.Errorf("insufficient quantity of material %s: required %d, available %d", material.ID, requiredQuantity, material.Quantity)
		}

		// Subtrai a quantidade do material
		material.Quantity -= requiredQuantity

		// Atualiza o material no ledger
		materialAsBytes, err := json.Marshal(material)
		if err != nil {
			return fmt.Errorf("failed to marshal updated material data: %v", err)
		}

		err = ctx.GetStub().PutState(compositeKey, materialAsBytes)
		if err != nil {
			return fmt.Errorf("failed to update material %s in ledger: %v", material.ID, err)
		}
	}

	// Cria uma chave composta usando o tipo "wand" e o ID da varinha
	objectType := "wand"
	compositeKey, err := ctx.GetStub().CreateCompositeKey(objectType, []string{wand.ID})
	if err != nil {
		return fmt.Errorf("failed to create composite key: %v", err)
	}

	// Tenta obter o estado existente da varinha no ledger
	existingWandBytes, err := ctx.GetStub().GetState(compositeKey)
	if err != nil {
		return fmt.Errorf("failed to get wand from ledger: %v", err)
	}

	if existingWandBytes != nil {
		// Se a varinha já existir, desserializa os dados existentes
		var existingWand Wand
		err = json.Unmarshal(existingWandBytes, &existingWand)
		if err != nil {
			return fmt.Errorf("failed to unmarshal existing wand data: %v", err)
		}

		// Verifica se os materiais da nova varinha são os mesmos da varinha existente
		if len(existingWand.Materials) == len(wand.Materials) {
			matches := true
			for i := range existingWand.Materials {
				if existingWand.Materials[i].ID != wand.Materials[i].ID ||
					existingWand.Materials[i].Name != wand.Materials[i].Name ||
					existingWand.Materials[i].Origin != wand.Materials[i].Origin ||
					existingWand.Materials[i].Quantity != wand.Materials[i].Quantity {
					matches = false
					break
				}
			}

			if matches {
				// Adiciona a nova quantidade à varinha existente
				existingWand.Quantity += wand.Quantity

				// Serializa novamente para armazenar no ledger
				wandAsBytes, err := json.Marshal(existingWand)
				if err != nil {
					return fmt.Errorf("failed to marshal updated wand data: %v", err)
				}

				// Atualiza o estado da varinha no ledger
				return ctx.GetStub().PutState(compositeKey, wandAsBytes)
			}
		}

		return fmt.Errorf("cannot create wand: existing wand has different materials")
	} else {
		// Se a varinha não existir, serializa e adiciona normalmente
		wandAsBytes, err := json.Marshal(wand)
		if err != nil {
			return fmt.Errorf("failed to marshal wand data: %v", err)
		}

		if (wand.Quantity != 1){
			return fmt.Errorf("The quantity_wand need to be 1, this wand doesn`t exist yet, you tried to create: %v", wand.Quantity)
		}

		// Armazena a varinha no ledger usando a chave composta
		return ctx.GetStub().PutState(compositeKey, wandAsBytes)
	}
}

// QueryMaterial retorna os materiais
func (s *StudioContract) QueryMaterial(ctx contractapi.TransactionContextInterface, materialID string) (*Material, error) {
    // Define o tipo do objeto e cria a chave composta com o materialID fornecido
    objectType := "material"
    compositeKey, err := ctx.GetStub().CreateCompositeKey(objectType, []string{materialID})
    if err != nil {
        return nil, fmt.Errorf("failed to create composite key for material %s: %v", materialID, err)
    }

    // Busca o estado do material usando a chave composta
    materialJSON, err := ctx.GetStub().GetState(compositeKey)
    if err != nil {
        return nil, fmt.Errorf("failed to read from world state for material %s: %v", materialID, err)
    }
    if materialJSON == nil {
        return nil, fmt.Errorf("material %s does not exist", materialID)
    }

    // Deserializa o JSON para a struct Material
    var material Material
    err = json.Unmarshal(materialJSON, &material)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal material JSON for material %s: %v", materialID, err)
    }

    return &material, nil
}

// QueryWand retorna as varinhas
func (s *StudioContract) QueryWand(ctx contractapi.TransactionContextInterface, wandID string) (*Wand, error) {
    objectType := "wand"
    compositeKey, err := ctx.GetStub().CreateCompositeKey(objectType, []string{wandID})
    if err != nil {
        return nil, fmt.Errorf("failed to create composite key: %v", err)
    }

    wandJSON, err := ctx.GetStub().GetState(compositeKey)
    if err != nil {
        return nil, fmt.Errorf("failed to read from world state: %v", err)
    }
    if wandJSON == nil {
        return nil, fmt.Errorf("the wand %s does not exist", wandID)
    }

    var wand Wand
    err = json.Unmarshal(wandJSON, &wand)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal wand JSON: %v", err)
    }

    return &wand, nil
}

// GetAllMaterials retorna todos os materiais disponíveis no ledger
func (sc *StudioContract) GetAllMaterials(ctx contractapi.TransactionContextInterface) ([]Material, error) {
	// Utiliza o Composite para busar os materiais
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey("material", []string{})
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve materials: %v", err)
	}
	defer resultsIterator.Close()

	var materials []Material

	// Itera sobre os resultados e desserializa apenas materiais
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("Error retrieving next result: %v", err)
		}

		var material Material
		// Tenta desserializar o dado como um material
		err = json.Unmarshal(queryResponse.Value, &material)
		if err != nil {
			// Se não for possível desserializar como Material, ignora o registro
			continue
		}

		// Adiciona o novo material válido a lista
		materials = append(materials, material)
	}

	return materials, nil
}

// GetAllWands retorna todas as varinhas disponíveis no ledger
func (sc *StudioContract) GetAllWands(ctx contractapi.TransactionContextInterface) ([]Wand, error) {
	// Utiliza o Composite para busar as varinhas
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey("wand", []string{})
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve wands: %v", err)
	}
	defer resultsIterator.Close()

	var wands []Wand

	// Itera sobre os resultados e desserializa apenas varinhas
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("Error retrieving next result: %v", err)
		}

		var wand Wand
		// Tenta desserializar o dado como uma varinha
		err = json.Unmarshal(queryResponse.Value, &wand)
		if err != nil {
			// Se não for possível desserializar como Wand, ignora o registro
			continue
		}

		// Adiciona a varinha válida à lista
		wands = append(wands, wand)
	}

	return wands, nil
}