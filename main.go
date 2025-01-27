package main

import (
	"example/go-crud/controllers"
	"example/go-crud/initializers"

	"github.com/gin-gonic/gin"
)
func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.GET("/posts", controllers.PostList)
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
