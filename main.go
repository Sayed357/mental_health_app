package main

import (
	"log"
	"mental_health_app/database"
	"mental_health_app/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDatabase()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.SetupRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "2313"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
