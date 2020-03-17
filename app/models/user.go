package models

// User Struct (Model)
type User struct {
	UserID   uint   `gorm:"primary_key" json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password"`
}

// TableName return name of database table
func (u *User) TableName() string {
	return "users"
}
