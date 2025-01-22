package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Quote_dollar struct {
	Dolar_bid string `json:"bid"`
	Dolar_ask string `json:"ask"`
	Date      string `json:"create_date"`
}

type Response struct {
	Currency  string `json:"currency"`
	Dolar_bid string `json:"bid"`
	Dolar_ask string `json:"ask"`
	Date      string `json:"create_date"`
}

func get_quote_dolar() (*Quote_dollar, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error in the request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api answer code: %d", resp.StatusCode)
	}

	var result map[string]Quote_dollar
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error in the codification: %v", err)
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

	response := Response{
		Currency:  "USD-BRL",
		Dolar_bid: quote.Dolar_bid,
		Dolar_ask: quote.Dolar_ask,
		Date:      quote.Date,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"error":"Error to get dollar quote"}`, http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/quote/dollar", handlerOperation)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
