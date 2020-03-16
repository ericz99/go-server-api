package migrate

import "github.com/jinzhu/gorm"

// Book Struct (Model)
type Book struct {
	gorm.Model
	BookID int    `gorm:"column:book_id" json:"bookid"`
	ISBN   string `gorm:"column:isbn" json:"isbn"`
	Title  string `gorm:"column:title" json:"title"`
	// Author    Author `gorm:"foreignkey:BookRefer"` // use BookRefer as foreign key
	// BookRefer uint
}

// Author Struct (Model)
type Author struct {
	gorm.Model
	Name  string `gorm:"type:varchar(50)" json:"name"`
	Books []Book `gorm:"foreignkey:BookRefer"`
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

// DBMigrate will migrate all models
func DBMigrate(db *gorm.DB) *gorm.DB {
	// Automatically create migration as per model
	db.Debug().AutoMigrate(&Book{}, &Author{}, &Response{}, &Test{})
	return db
}
