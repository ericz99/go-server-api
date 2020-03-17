package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim Struct (Model)
type Claim struct {
	UserID uint
	Name   string
	Email  string
	*jwt.StandardClaims
}

// TableName return name of database table
func (c *Claim) TableName() string {
	return "claims"
}
