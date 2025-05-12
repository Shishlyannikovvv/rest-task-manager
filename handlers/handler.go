package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	// Здесь будет код для получения задач из базы данных
	c.JSON(http.StatusOK, gin.H{"message": "Список задач"})
}

func CreateTask(c *gin.Context) {
	// Здесь будет код для создания новой задачи в базе данных
	c.JSON(http.StatusCreated, gin.H{"message": "Задача создана"})
}
