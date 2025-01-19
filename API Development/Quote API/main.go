package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Quote_dolar struct {
	Dolar_bid string `json:"bid"`
	Dolar_ask string `json:"ask"`
	Date      string `json:"date"`
}

func get_quote_dolar() (*Quote_dolar, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("resposta da api com código: %d", resp.StatusCode)
	}

	var result map[string]Quote_dolar
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("erro ao decodificar a resposar: %v", err)
	}

	quote := result["USDBRL"]
	return &quote, nil
}

func handlerOperation(w http.ResponseWriter, r *http.Request) {
	quote, err := get_quote_dolar()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error to get dollar quote: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"moeda":  "USD-BRL",
		"compra": quote.Dolar_bid,
		"venda":  quote.Dolar_ask,
		"data":   quote.Date,
	})
}

func main() {
	http.HandleFunc("/quote/dollar", handlerOperation)

	if err := http.ListenAndServe(":1000", nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
