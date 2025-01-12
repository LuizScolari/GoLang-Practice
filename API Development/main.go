package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todoList struct {
	ID          int    `json:"id"`
	Task        string `json:"task"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var tasks = []todoList{
	{ID: 0, Task: "Learn GoLang", Description: "Study and Practice GoLang", Completed: false},
}

func getTodoList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func postTodoList(c *gin.Context) {
	var newTask todoList

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTodoList)
	router.POST("/tasks", postTodoList)

	router.Run("localhost:5050")
}
