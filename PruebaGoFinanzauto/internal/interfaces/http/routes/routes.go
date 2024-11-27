package routes

import (
	"PruebaGoFinanzauto/internal/interfaces/http/handlers"
	"PruebaGoFinanzauto/internal/interfaces/http/middleware"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, studentHandler *handlers.StudentHandler) {
	api := e.Group("/api/v1")

	// valida headers api key
	api.Use(middleware.APIKeyMiddleware)

	students := api.Group("/students")
	students.POST("", studentHandler.Create)
	students.GET("", studentHandler.GetAll)
	students.GET("/:id", studentHandler.GetByID)
	students.PUT("/:id", studentHandler.Update)
	students.DELETE("/:id", studentHandler.Delete)
}
