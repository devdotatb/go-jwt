package main

import (
	"github.com/devdotatb/go-jwt.git/initializers"
	"github.com/devdotatb/go-jwt.git/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToGoLangDb()
}
func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
