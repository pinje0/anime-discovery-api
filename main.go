package main

import (
	"log"

	"anime-discovery-api/config"
	"anime-discovery-api/handlers"
	"anime-discovery-api/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}

	cfg := config.Load()

	services.Init(cfg)

	r := gin.New()

	r.GET("/health", handlers.HealthCheck)
	r.GET("/anime/top", handlers.GetTopAnime)
	r.GET("/anime/:id", handlers.GetAnimeByID)

	log.Printf("Server started on :%s", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
