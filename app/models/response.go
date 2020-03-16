package models

// Response Status Struct (Model)
type Response struct {
	Status  string `gorm:"type:varchar(50)" json:"status" validate:"required"`
	Message string `gorm:"type:varchar(50)" json:"message" validate:"required"`
}

// TableName return name of database table
func (r *Response) TableName() string {
	return "responses"
}
