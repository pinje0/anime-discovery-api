package routes

import (
	"anime-discovery-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
		api.GET("/anime/top", handlers.GetTopAnime)
		api.GET("/anime/:id", handlers.GetAnimeByID)
	}
}
