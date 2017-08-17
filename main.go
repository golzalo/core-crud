package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", CreateTodo)
		v1.GET("/", FetchAllTodo)
		v1.GET("/:id", FetchSingleTodo)
		v1.DELETE("/:id", DeleteSingleTodo)
		v1.POST("/:id", UpdateSingleTodo)

	}
	router.Run()

}

func CreateTodo(c *gin.Context) {
	var json Todo
	if c.BindJSON(&json) == nil {
		todo := Todo{Title: json.Title, Completed: json.Completed}
		ID := add(&todo)
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": ID})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Something went wrong"})
	}

}

func FetchAllTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": getAll()})
}

func FetchSingleTodo(c *gin.Context) {
	var todo *Todo
	todoId, _ := strconv.Atoi(c.Param("id"))
	todo = get(todoId)
	if todo == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todo})
}

func DeleteSingleTodo(c *gin.Context) {
	todoId, _ := strconv.Atoi(c.Param("id"))
	remove(todoId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "deleted"})
}

func UpdateSingleTodo(c *gin.Context) {
	todoId, _ := strconv.Atoi(c.Param("id"))
	var json Todo
	if c.BindJSON(&json) == nil {
		todo := Todo{Title: json.Title, Completed: json.Completed}
		update(todoId, &todo)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated", "resourceId": todoId})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Something went wrong"})
	}
}
