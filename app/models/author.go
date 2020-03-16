package models

// Author Struct (Model)
type Author struct {
	AuthorID uint   `gorm:"primary_key" json:"authorId"`
	Name     string `json:"name"`
}

// TableName return name of database table
func (a *Author) TableName() string {
	return "authors"
}

// # GET Name Method
func (a Author) getName() string {
	return a.Name
}

// # SET Name Method
func (a *Author) setName(name string) {
	a.Name = name
}
