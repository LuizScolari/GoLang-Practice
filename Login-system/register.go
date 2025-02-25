package main

import "log"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(Email string, Password string) {

	login := Login{
		Email:    Email,
		Password: Password,
	}

	err := saveDB(login)
	if err != nil {
		log.Print("Error to save the login in DB", err)
	}
}
