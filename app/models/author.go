package models

import (
	"github.com/jinzhu/gorm"
)

// Author Struct (Model)
type Author struct {
	gorm.Model
	Name  string `gorm:"type:varchar(50)" json:"name"`
	Books []Book `gorm:"foreignkey:BookRefer"`
}

// TableName return name of database table
func (a *Author) TableName() string {
	return "author"
}

// # GET Name Method
func (a Author) getName() string {
	return a.Name
}

// # SET Name Method
func (a *Author) setName(name string) {
	a.Name = name
}
