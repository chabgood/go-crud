package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-crud/controllers"
	"github.com/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToPostgres()
}


func main() {
	r := gin.Default()
	r.POST("/post", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/post/:id", controllers.PostsShow)
	r.PUT("/post/:id", controllers.PostsUpdate)
	r.DELETE("/post/:id", controllers.PostsDelete)
	r.Run()

}
