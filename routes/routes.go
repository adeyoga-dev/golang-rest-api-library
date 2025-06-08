package routes

import (
	"rest-api-library/handlers"

	"github.com/labstack/echo/v4"
)

func BookRoutes(e *echo.Echo) {
	api := e.Group("/api")
	bookGroup := api.Group("/books")
	bookGroup.GET("", handlers.GetBooks)
	bookGroup.GET("/:id", handlers.GetBookByID)
	bookGroup.POST("", handlers.CreateBook)
	bookGroup.PUT("/:id", handlers.UpdateBook)
	bookGroup.DELETE("/:id", handlers.DeleteBook)
}

func SetupRoutes(e *echo.Echo) {
	BookRoutes(e)
}
