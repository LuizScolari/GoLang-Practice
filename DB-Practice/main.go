package main

import (
	"DB-Practice/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	// Conecta ao banco de dados
	var database *gorm.DB = db.ConnectDB()

	// Cria ou atualiza a estrutura da tabela "users" com base no modelo User
	database.AutoMigrate(&User{})

	// Insere um novo registro no banco de dados
	user := User{Name: "Luiz", Age: 19}
	database.Create(&user)
}
