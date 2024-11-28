package main

import (
    "log"
    "net/http"
    "os"
    
    "github.com/joho/godotenv"
    "github.com/Tabintel/invoice-system/internal/database"
    "github.com/Tabintel/invoice-system/internal/server"
    httpSwagger "github.com/swaggo/http-swagger"
    _ "github.com/Tabintel/invoice-system/internal/docs"
)

// @title Invoice System API
// @version 1.0
// @description A modern invoice management system API
// @host localhost:8080
// @BasePath /api
func main() {
    if err := godotenv.Load(); err != nil {
        log.Printf("No .env file found")
    }
    
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL environment variable is not set")
    }
    
    client := database.NewClient(dbURL)
    defer client.Close()
    
    srv := server.NewServer(client)
    
    // Add Swagger endpoint
    srv.Router().Get("/swagger/*", httpSwagger.Handler(
        httpSwagger.URL("/swagger/doc.json"),
    ))
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    log.Println("Successfully connected to NeonTech PostgreSQL database")
    log.Printf("Server starting on port %s", port)
    if err := http.ListenAndServe(":"+port, srv); err != nil {
        log.Fatal(err)
    }
}
