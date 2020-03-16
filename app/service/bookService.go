package service

import (
	"go-server-api/app/models"

	"github.com/jinzhu/gorm"
)

// BookService (Service)
type BookService struct {
	Book models.Book
}

// FindAll METHOD
func (b *BookService) FindAll(db *gorm.DB, books *[]models.Book) (err error) {
	if err = db.Preload("Author").Find(&books).Error; err != nil {
		return err
	}

	return nil
}

// SaveBook METHOD
func (b *BookService) SaveBook(db *gorm.DB, book *models.Book) (err error) {
	if err = db.Create(&book).Error; err != nil {
		return err
	}

	return nil
}

// GetBookByID | GET BOOK BY ID
func (b *BookService) GetBookByID(db *gorm.DB, id int, book *models.Book) (err error) {
	if err = db.Preload("Author").Where("book_id = ?", id).First(&book).Error; err != nil {
		return err
	}

	return nil
}

// DeleteBook | DELETE BOOK
func (b *BookService) DeleteBook(db *gorm.DB, book *models.Book) (err error) {
	if err = db.Delete(&book).Error; err != nil {
		return err
	}

	return nil
}
