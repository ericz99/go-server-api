package utils

import (
	"fmt"
	"go-server-api/app/models"
	"go-server-api/config"
	"log"

	// MYSQL
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Init Database Connection
func init() {
	config := config.GetConfig()

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	conn, err := gorm.Open("mysql", dbURI)

	if err != nil {
		log.Fatal("Could not connect to database")
		panic(err.Error())
	} else {
		fmt.Println("connected!")
	}

	// # reinitize instance
	db = conn

	// prining query
	db.LogMode(true)
	fmt.Println("Successfully Connected to MySQL Database")
	// Automatically create migration as per model
	db.Debug().AutoMigrate(&models.Book{}, &models.Author{}, &models.Response{}, &models.Test{})
}

// GetDB function return the instance of db
func GetDB() *gorm.DB {
	return db
}
