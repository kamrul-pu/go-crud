package controllers

import (
	"example/go-crud/initializers"
	"example/go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context){
	// Get data off req body
	var body struct {
		Body string
		Title string

	}
	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create
	if result.Error!=nil{
		c.Status(400)
		return
	}
	//Return it
	c.JSON(201, gin.H{"post": post,})
}
func PostList (c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{"data": posts,})
}
func GetPost(c *gin.Context) {
	// Get id of url
	id := c.Param("id")
	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)
	// Respond with the post
	c.JSON(200, gin.H{"data": post})
}

func PostUpdate(c *gin.Context){
	// Get the id of the url
	id := c.Param("id")
	// Get the data of request body
	var body struct {
		Body string
		Title string

	}
	c.Bind(&body)
	// Find the post we are updating
	var post models.Post
	initializers.DB.First(&post, id)
	// Updating the data
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})
	// Respond the data
	c.JSON(200, gin.H{"data": post})
}


func PostDelete(c *gin.Context){
	// Get the post id from the url
	id := c.Param("id")

	// Delete the posts
	initializers.DB.Delete(&models.Post{}, id)
	// Respond
	c.Status(200)
	c.JSON(200, gin.H{"message": "Post delete Success"})
}