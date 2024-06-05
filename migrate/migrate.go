package main

import (
	"github.com/wallacce135/gin-crud/initializers"
	"github.com/wallacce135/gin-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
