package router

import (
	"github.com/gin-gonic/gin"
	"rest-task-manager/handlers" // Замените "your_project_name" на фактическое имя вашего проекта
   )
   
func SetupRouter() *gin.Engine {
	r := gin.Default()
   
	// Определяем маршруты и связываем их с обработчиками
	r.GET("/tasks", handlers.GetTasks)
	r.POST("/tasks", handlers.CreateTask)
   
	return r
}