package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler
func hello(c echo.Context) error {
	version := os.Getenv("VERSION")
	hostname, err := os.Hostname()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "get hostname error")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"version":  version,
		"hostname": hostname,
	})
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/hello", hello)
	e.GET("/world", hello)

	go func() {
		e.Logger.Fatal(e.Start(":8080"))
	}()

	// Start server
	e.Logger.Fatal(e.StartTLS(":8443", "./certs/simple.org.crt", "./certs/simple.org.key"))
}
