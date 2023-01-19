package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-crud/initializers"
	"github.com/go-crud/models"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))

	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	c.Bind(&body)

	post.Title = body.Title
	post.Body = body.Body

	initializers.DB.Save(&post)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	var post models.Post
	initializers.DB.Delete(&post, c.Param("id"))

	c.Status(200)
}
