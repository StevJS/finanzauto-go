package main

import (
	"PruebaGoFinanzauto/internal/infrastructure/database"
	"PruebaGoFinanzauto/internal/interfaces/http/handlers"
	"PruebaGoFinanzauto/internal/interfaces/http/routes"
	"PruebaGoFinanzauto/internal/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {
	// Database init
	db, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo := database.NewPostgresRepository(db)

	studentUseCase := usecases.NewStudentUseCase(repo)
	studentHandler := handlers.NewStudentHandler(studentUseCase)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	routes.SetupRoutes(e, studentHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
