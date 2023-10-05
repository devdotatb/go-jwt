package initializers

import (
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToIdentityDb() {
	var err error
	// github.com/denisenkom/go-mssqldb
	dsn := os.Getenv("IDENTITY_DB")
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}
}

func ConnectToGoLangDb() {
	var err error
	// github.com/denisenkom/go-mssqldb
	dsn := os.Getenv("GoLang_DB")
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}
}
