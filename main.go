package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id       string `json:"id"`
	Item     string `json:"item"`
	Complete bool   `json:"complete"`
}

var todos = []todo{
	{Id: "1", Item: "Clean Room", Complete: false},
	{Id: "2", Item: "Read Book", Complete: false},
	{Id: "3", Item: "Record Item", Complete: true},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	var err = context.BindJSON(&newTodo)
	if err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusOK, newTodo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)

	router.Run("localhost:9090")
}
