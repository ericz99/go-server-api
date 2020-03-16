package models

// Book Struct (Model)
type Book struct {
	BookID uint   `gorm:"primary_key" json:"bookId"`
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author" gorm:"foreignkey:AuthorID"`
}

// TableName return name of database table
func (b *Book) TableName() string {
	return "books"
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
