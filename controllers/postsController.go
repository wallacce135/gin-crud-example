package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wallacce135/gin-crud/initializers"
	"github.com/wallacce135/gin-crud/models"
)

func PostsCreate(context *gin.Context) {
	//Get data from request body

	var body struct {
		Title string
		Body  string
	}

	context.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		context.Status(400)
		return
	}

	context.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(context *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	context.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetOnePost(context *gin.Context) {

	id := context.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	context.JSON(200, gin.H{
		"posts": post,
	})
}
