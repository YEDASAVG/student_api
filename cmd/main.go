package main

import (
	"log"
	"net/http"

	"github.com/YEDASAVG/student_api/internal/config"
	"github.com/YEDASAVG/student_api/internal/db"
	"github.com/YEDASAVG/student_api/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	database, err := db.ConnectDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal("Failed to get database instance", err)
	}
	defer sqlDB.Close()

	database.AutoMigrate(&models.Student{})

	router := gin.Default()

	router.GET("/healthcheck", func(ctx *gin.Context) { // Healthcheck endpoint.
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	router.Run(":" + cfg.Port)
}
