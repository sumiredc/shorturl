package main

import (
	"log"

	"github.com/sumiredc/shorturl/app/interface/gateway/database"
	"github.com/sumiredc/shorturl/app/interface/route"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	_, err = database.New().Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database.")
		return
	}

	e := echo.New()
	r := route.New()
	r.Register(e)

	e.Start(":8080")
}
