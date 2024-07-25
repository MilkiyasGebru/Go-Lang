package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"taskAPI/Data"
	"taskAPI/Model"
)

func GetAllTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Data.Tasks)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	taskId, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameter ID given"})
	}

	for _, task := range Data.Tasks {
		if task.ID == taskId {
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	return
}

func CreateTask(c *gin.Context) {
	var newTask Model.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON Body"})
		return
	}

	Data.Tasks = append(Data.Tasks, newTask)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Item successfully added"})
	return
}

func UpdateTask(c *gin.Context) {
	var updatedTask Model.Task

	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": " Invalid JSON Body "})
		return
	}

	for _, task := range Data.Tasks {
		if task.ID == updatedTask.ID {
			task.Status = updatedTask.Status
			//task.Deadline = updatedTask.Deadline
			task.Description = updatedTask.Description
			task.Title = updatedTask.Title
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Task not found"})
	return
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": " Invalid Param Id"})
	}
	for i, task := range Data.Tasks {
		if task.ID == taskId {
			Data.Tasks = append(Data.Tasks[:i], Data.Tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Task has been successfully removed"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task to be deleted do not exists"})
	return
}
