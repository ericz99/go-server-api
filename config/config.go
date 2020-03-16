package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config Struct (Model)
type Config struct {
	DB *DBConfig
}

// DBConfig Struct (Model)
type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

// GetConfig Method
func GetConfig() *Config {

	e := godotenv.Load()

	if e != nil {
		fmt.Println(e)
	}

	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "127.0.0.1",
			Port:     3306,
			Username: os.Getenv("db_user"),
			Password: os.Getenv("db_pass"),
			Name:     os.Getenv("db_name"),
			Charset:  "utf8",
		},
	}
}
