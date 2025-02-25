package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Login2 struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ConnectDB() *gorm.DB {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error to connect in the datebase: %v", err)
	}

	return db
}

func saveDB(login Login) error {
	database := ConnectDB()

	database.AutoMigrate(&Login2{})

	saveLogin := Login2{
		Email:    login.Email,
		Password: login.Password,
	}

	result := database.Create(&saveLogin)
	return result.Error
}
