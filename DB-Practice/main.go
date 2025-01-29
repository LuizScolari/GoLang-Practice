package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Datas struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Gender  string `json:"gender"`
	Message string `json:"message"`
}

func handleFormSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}

	// Pega os dados do formulário
	datas, err := parseForm(r)
	if err != nil {
		http.Error(w, "error: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = saveDB(datas)
	if err != nil {
		fmt.Errorf("error to save in the db")
	}

	w.Header().Set("Content-Typer", "application/json")
	json.NewEncoder(w).Encode(datas)
}

func parseForm(r *http.Request) (Datas, error) {
	err := r.ParseForm()
	if err != nil {
		return Datas{}, err
	}

	return Datas{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Gender:  r.FormValue("gender"),
		Message: r.FormValue("message"),
	}, nil
}

func server() {
	http.HandleFunc("/submit", handleFormSubmission)
	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
