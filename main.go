package main

import (
	"go-crud-api/controllers"
	"go-crud-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    models.ConnectDatabase()

	router.POST("/posts", controllers.CreatePost)

    router.Run("localhost:8080")
}