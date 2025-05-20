package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func main() {
	// Загружаем переменные окружения из .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Проверка наличия переменных окружения
	requiredVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "SSL_MODE"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Missing required environment variable: %s\n", v)
		}
	}

	// Строка подключения
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"))

	// Открытие соединения с БД
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v\n", err)
	}
	defer db.Close()

	// Проверка подключения к БД
	if err := db.Ping(); err != nil {
		log.Fatalf("Cannot connect to database: %v\n", err)
	}

	// Инициализация роутера
	r := gin.Default()

	// Пример маршрута
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Пример маршрута для получения задач
	r.GET("/tasks", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, description, completed FROM tasks")
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch tasks"})
			return
		}
		defer rows.Close()

		var tasks []Task
		for rows.Next() {
			var task Task
			if err := rows.Scan(&task.ID, &task.Description, &task.Completed); err != nil {
				c.JSON(500, gin.H{"error": "Failed to scan task"})
				return
			}
			tasks = append(tasks, task)
		}

		c.JSON(200, tasks)
	})

	// Маршрут для добавления задачи
	r.POST("/tasks", func(c *gin.Context) {
		var newTask Task
		if err := c.ShouldBindJSON(&newTask); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		_, err := db.Exec("INSERT INTO tasks (description, completed) VALUES ($1, $2)", newTask.Description, newTask.Completed)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to insert task"})
			return
		}

		c.JSON(201, gin.H{"message": "Task created"})
	})

	// Запуск сервера на порту 8080 или указанном в переменной окружения PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Порт по умолчанию
	}
	r.Run(":" + port)
}
