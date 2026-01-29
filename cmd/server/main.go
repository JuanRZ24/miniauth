package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"miniauth/internal/config"
	"miniauth/internal/db"
	"miniauth/internal/handlers"
	"miniauth/internal/middleware"
	"miniauth/internal/repositories"
	"miniauth/internal/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	repo := repositories.NewPostgresUserRepository(database)
	authService := services.NewAuthService(repo)
	authHandler := handlers.NewAuthHandler(authService)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	authGroup := r.Group("/api")
	authGroup.Use(middleware.AuthMiddleware())
	authGroup.GET("/me", authHandler.Me)

	r.Run(":" + cfg.AppPort)
}
