package main

import (
	"github.com/devdotatb/go-jwt.git/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func ping() (str string, hdl gin.HandlerFunc) {
	return "/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong2",
		})
	}
}

func main() {
	r := gin.Default()
	r.GET(ping())
	r.Run()
}
