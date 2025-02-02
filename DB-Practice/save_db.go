package main

import (
	"DB-Practice/db"

	"gorm.io/gorm"
)

type Datas2 struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Gender  string `json:"gender"`
	Message string `json:"message"`
}

func saveDB(datas Datas) error {
	// Conecta ao banco de dados
	database := db.ConnectDB()

	// Cria a tabela "datas2" se não existir
	database.AutoMigrate(&Datas2{})

	// Converte "Datas" (do formulário) para "Datas2" (do banco de dados)
	dataToSave := Datas2{
		Name:    datas.Name,
		Email:   datas.Email,
		Gender:  datas.Gender,
		Message: datas.Message,
	}

	// Insere os dados no banco
	result := database.Create(&dataToSave)
	return result.Error
}
