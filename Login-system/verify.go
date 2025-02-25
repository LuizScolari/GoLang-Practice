package main

import "fmt"

func getLoginByEmail(Email string) (bool, error) {
	database := ConnectDB()

	var login Login2
	result := database.Where("email = ?", Email).First(&login)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return false, nil
		}
		fmt.Printf("Error to find email", result.Error)
		return false, result.Error
	}

	return true, nil
}
