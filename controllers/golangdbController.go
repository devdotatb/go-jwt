package controllers

import (
	"github.com/devdotatb/go-jwt.git/initializers"
	"github.com/devdotatb/go-jwt.git/models"
	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	var body struct {
		Email string
	}
	c.Bind(&body)

	user := models.User{Email: body.Email}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersRead(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"user": users,
	})
}

func UserRead(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	initializers.DB.First(&user, id)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserUpdate(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Email string
	}
	c.Bind(&body)

	updated_user := models.User{Email: body.Email}

	var user models.User
	initializers.DB.First(&user, id)
	initializers.DB.Model(&user).Updates(&updated_user)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserDelete(c *gin.Context) {
	id := c.Param("id")

	var user_del models.User
	initializers.DB.Delete(&user_del, id)

	c.JSON(200, gin.H{
		"user": user_del,
	})
}
