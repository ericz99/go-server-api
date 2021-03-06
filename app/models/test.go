package models

// Test Struct (Model)
type Test struct {
	Msg string `gorm:"type:varchar(50)" json:"msg" validate:"required"`
}

// TableName return name of database table
func (t *Test) TableName() string {
	return "tests"
}
