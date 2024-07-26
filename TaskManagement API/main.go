package main

import (
	"github.com/gin-gonic/gin"
	"taskAPI/Controller"
)

func main() {
	router := gin.Default()

	router.GET("/", Controller.GetAllTasks)
	router.GET("/:id", Controller.GetTaskByID)
	router.POST("/", Controller.CreateTask)
	router.PUT("/", Controller.UpdateTask)
	router.DELETE("/:id", Controller.DeleteTask)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
