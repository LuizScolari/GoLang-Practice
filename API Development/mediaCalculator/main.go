package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Scores struct {
	Math    float64 `json:"math"`
	English float64 `json:"English"`
	Science float64 `json:"Science"`
}

type Result struct {
	Media float64 `json:"media"`
}

func calculateMedia(c *gin.Context) {
	scores := Scores{
		Math:    9,
		English: 8,
		Science: 10,
	}
	result := Result{
		Media: (scores.Math + scores.English + scores.Science) / 3,
	}
	c.JSON(http.StatusOK, result)
}

func main() {

	r := gin.Default()

	r.GET("/results", calculateMedia)

	err := r.Run(":9090")
	if err != nil {
		return
	}
}
