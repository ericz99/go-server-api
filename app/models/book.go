package models

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

// TableName return name of database table
func (b *Book) TableName() string {
	return "book"
}

// # GET ID Method
func (b Book) getID() int {
	return b.BookID
}

// # GET ISBN Method
func (b Book) getISBN() string {
	return b.ISBN
}

// # GET Name Method
func (b Book) getTitle() string {
	return b.Title
}

// # SET Name Method
func (b *Book) setTitle(title string) {
	b.Title = title
}
