package models

import (
	"github.com/jinzhu/gorm"
)

// Response Status Struct (Model)
type Response struct {
	gorm.Model
	Status  string `gorm:"type:varchar(50)" json:"status" validate:"required"`
	Message string `gorm:"type:varchar(50)" json:"message" validate:"required"`
}

// TableName return name of database table
func (r *Response) TableName() string {
	return "response"
}
