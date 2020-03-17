package migrate

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// Book Struct (Model)
type Book struct {
	BookID uint   `gorm:"primary_key" json:"bookId"`
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author" gorm:"foreignkey:AuthorID"`
}

// Author Struct (Model)
type Author struct {
	AuthorID uint   `gorm:"primary_key" json:"authorId"`
	Name     string `json:"name"`
	// Books    []Book `json:"books"`
}

// Response Status Struct (Model)
type Response struct {
	gorm.Model
	Status  string `gorm:"type:varchar(50)" json:"status" validate:"required"`
	Message string `gorm:"type:varchar(50)" json:"message" validate:"required"`
}

// Test Struct (Model)
type Test struct {
	gorm.Model
	Msg string `gorm:"type:varchar(50)" json:"msg" validate:"required"`
}

// Claim Struct (Model)
type Claim struct {
	UserID uint
	Name   string
	Email  string
	*jwt.StandardClaims
}

// User Struct (Model)
type User struct {
	UserID   uint   `gorm:"primary_key" json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password"`
}

// DBMigrate will migrate all models
func DBMigrate(db *gorm.DB) *gorm.DB {

	e := godotenv.Load()

	if e != nil {
		fmt.Println(e)
		return nil
	}

	mode := os.Getenv("mode")

	// # only drop table if doesn't exist
	if mode == "development" {
		// # Create the database. This is a one-time step.
		// # Comment out if running multiple times - You may see an error otherwise
		db.Exec("DROP DATABASE testdb")
		db.Exec("CREATE DATABASE testdb")
		db.Exec("USE testdb")
	} else {
		// # ONLY CREATE DATABASE IF NOT EXIST IN PRODUCTION
		db.Exec("CREATE DATABASE IF NOT EXISTS testdb")
	}

	// # Automatically create migration as per model
	db.Debug().AutoMigrate(&Book{}, &Author{}, &Response{}, &Test{}, &User{}, &Claim{})

	// # return instance
	return db
}
