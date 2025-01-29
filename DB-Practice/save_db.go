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
	var database *gorm.DB = db.ConnectDB()

	// Cria ou atualiza a estrutura da tabela "users" com base no modelo User
	database.AutoMigrate(&Datas2{})

	// Insere um novo registro no banco de dados
	result := database.Create(&datas)

	return result.Error
}
