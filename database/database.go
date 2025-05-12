package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
    // Загружаем .env файл
    err := godotenv.Load()
    if err != nil {
        log.Println("Не удалось загрузить .env файл")
    }

    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    sslmode := os.Getenv("DB_SSLMODE")

    if host == "" {
        host = "localhost"
    }
    if port == "" {
        port = "5432"
    }
    if sslmode == "" {
        sslmode = "disable"
    }

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    DB = db
    log.Println("Successfully connected to PostgreSQL!")
}