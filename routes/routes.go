package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitRoutes() {
	router := gin.Default()

	router.POST("/jobs", recieveJobs)
	router.GET("/job/:id", getJob)
	godotenv.Load()
	port := os.Getenv("PORT")

	router.Run(":" + port)
}