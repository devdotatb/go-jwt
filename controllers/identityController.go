package controllers

import "github.com/gin-gonic/gin"

func FuckingGet() {

}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
