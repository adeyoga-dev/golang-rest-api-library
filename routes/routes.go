package routes

import (
	"rest-api-library/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/books", handlers.GetBooks)
	e.GET("/books/:id", handlers.GetBookByID)
	e.POST("/books", handlers.CreateBook)
	e.PUT("/books/:id", handlers.UpdateBook)
	e.DELETE("/books/:id", handlers.DeleteBook)
}
