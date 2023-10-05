package main

import (
	"github.com/devdotatb/go-jwt.git/controllers"
	"github.com/devdotatb/go-jwt.git/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToGoLangDb()
}

func main() {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)

	r.POST("/user", controllers.UserCreate)
	r.GET("/user", controllers.UsersRead)
	r.GET("/user/:id", controllers.UserRead)
	r.PUT("/user/:id", controllers.UserUpdate)
	r.DELETE("/user/:id", controllers.UserDelete)
	r.Run()
}
