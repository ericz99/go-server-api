package auth

import (
	"encoding/json"
	"fmt"
	"go-server-api/app/models"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// ErrorResponse Struct (Model)
type ErrorResponse struct {
	Error error
}

// ErrorMessage Struct (Model)
type ErrorMessage struct {
	Message string
}

// Register Method - Will auto hash password using bcrypt
func Register(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)
	w.Header().Set("Content-Type", "application/json")

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	// # set password to hashed password
	user.Password = string(hashed)
	// then create user..
	if err = db.Create(&user).Error; err != nil {
		resp := map[string]interface{}{"status": false, "message": "Account already exist!"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	// return back user object if successfully created
	json.NewEncoder(w).Encode(user)
}

// ComparePassword METHOD - Will compare hash password and plain text
func ComparePassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Login Method - Will generate new jwt token
func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp := map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	// # generate token & return token
	resp := GenerateToken(user.Email, user.Password, db)
	json.NewEncoder(w).Encode(resp)
}

// GenerateToken Method - Will generate a valid JWT token
func GenerateToken(email, password string, db *gorm.DB) map[string]interface{} {
	e := godotenv.Load()

	if e != nil {
		fmt.Println(e)
	}

	user := &models.User{}

	// # find user email
	if err := db.Where("email = ?", email).Find(&user).Error; err != nil {
		resp := map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	isMatch := ComparePassword(user.Password, password)

	// # if password doesnt match
	if isMatch != true {
		resp := map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	// # create payload
	claim := &models.Claim{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	// # create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, error := token.SignedString([]byte(os.Getenv("secret_key")))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	// resp["user"] = user
	return resp
}

// DecodeToken Method - Decode JWT token
func DecodeToken(tokenStr string) (*models.Claim, error) {
	e := godotenv.Load()

	if e != nil {
		fmt.Println(e)
	}

	tk := &models.Claim{}
	secret := os.Getenv("secret_key")

	_, err := jwt.ParseWithClaims(tokenStr, tk, func(token *jwt.Token) (interface{}, error) {
		// # validate token after parsing it
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return tk, nil
}
