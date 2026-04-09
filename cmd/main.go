package main

import (
	"log"
	"net/http"
	"os"

	"github.com/YEDASAVG/student_api/internal/db"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pool, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer pool.Close()
	router := gin.Default()

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	router.Run(":" + os.Getenv("PORT"))
}
