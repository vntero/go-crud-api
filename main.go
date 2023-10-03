package main

import (
	"go-crud-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() 

	models.ConnectDatabase()

	router.Run("localhost:3000")
}