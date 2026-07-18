package backend

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Welcome to the scales API!"})
	})

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("Failed to start the backend.", "error", err)
	}
}