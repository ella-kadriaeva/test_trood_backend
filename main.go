package main

import (
    "log"
    "os"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    db "github.com/troodinc/trood-front-hackathon/database"
    _ "github.com/troodinc/trood-front-hackathon/docs"
    "github.com/troodinc/trood-front-hackathon/handlers"
    "github.com/gin-contrib/cors"
    "time"
)

// @title Trood Front Hackathon API
// @version 1.0
// @description This is the API documentation for the Trood Front Hackathon. Welcome to hell.
// @host localhost:8080
// @BasePath /

func main() {
    db.InitDatabase()
    handlers.InitProjects()

    r := gin.Default()
    
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.GET("/projects", handlers.GetProjects)
    r.GET("/projects/:id", handlers.GetProjectByID)
    r.POST("/projects", handlers.CreateProject)
    r.PUT("/projects/:id", handlers.EditProject)
    r.DELETE("/projects/:id", handlers.DeleteProject)

    r.GET("/projects/:id/vacancies", handlers.GetVacancies)
    r.POST("/projects/:id/vacancies", handlers.CreateVacancy)
    r.PUT("/vacancies/:id", handlers.EditVacancy)
    r.DELETE("/vacancies/:id", handlers.DeleteVacancy)

    // Use the PORT environment variable or default to 8080 for local development
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default to 8080 for local development
    }

    log.Println("Server running on port " + port)
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}