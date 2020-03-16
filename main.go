package main

import (
	"fmt"
	"go-server-api/app/routers"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	e := godotenv.Load()

	if e != nil {
		fmt.Println(e)
	}

	port := os.Getenv("port")

	if port == "" {
		port = "8080" // default port if not specify
	}

	// # initalize router + run server + database
	app := &routers.App{}
	app.Initialize()
	app.Run(":" + port)
}
