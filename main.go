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
	router.GET("/posts", controllers.FindPosts)
	router.GET("/posts/:id", controllers.FindPost)
	router.PATCH("posts/:id", controllers.UpdatePost)
	router.DELETE("posts/:id", controllers.DeletePost)

	router.Run("localhost:8080")
}
