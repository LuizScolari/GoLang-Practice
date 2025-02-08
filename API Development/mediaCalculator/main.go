package main

import "github.com/gin-gonic/gin"

type Scores struct {
	Math    float64 `json:"math"`
	English float64 `json:"English"`
	Science float64 `json:"Science"`
}

type Result struct {
	Media float64 `json:"media"`
}

func calculateMedia(c *gin.Context) {

}

func main() {

	r := gin.Default()

	r.POST("results", calculateMedia)

	err := r.Run(".9090")
	if err != nil {
		return
	}
}
