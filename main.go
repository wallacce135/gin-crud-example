package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wallacce135/gin-crud/controllers"
	"github.com/wallacce135/gin-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {

	app := gin.Default()
	app.POST("/posts", controllers.PostsCreate)
	app.GET("/posts", controllers.GetPosts)
	app.GET("/posts/:id", controllers.GetOnePost)

	app.Run()

}
