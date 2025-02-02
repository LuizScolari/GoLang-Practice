package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Estrutura para os dados do formulário (deve ser igual ao frontend)
type Datas struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Gender  string `json:"gender" form:"gender"`
	Message string `json:"message" form:"message"`
}

func main() {
	router := gin.Default()

	// Middleware para permitir CORS (compatível com o frontend)
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	// Rota para receber o formulário
	router.POST("/submit", handleFormSubmission)

	// Inicia o servidor
	log.Println("Servidor rodando na porta 8080...")
	router.Run(":8080")
}

func handleFormSubmission(c *gin.Context) {
	var datas Datas

	// Extrai os dados do formulário HTML
	if err := c.ShouldBind(&datas); err != nil {
		c.JSON(400, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Salva no banco de dados
	if err := saveDB(datas); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao salvar no banco de dados: " + err.Error()})
		return
	}

	// Retorna os dados salvos como confirmação
	c.JSON(200, datas)
}
