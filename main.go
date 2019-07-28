package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newService() *Service {
	service := &Service{
		Server: echo.New(),
		Store:  createDB(),
	}

	return service
}

func main() {

	service := newService()

	// Middleware
	service.Server.Use(middleware.Logger())
	service.Server.Use(middleware.Recover())

	// Routes
	service.Server.GET("/", hello)
	service.Server.POST("/login", createUser)
	// e.POST("/login", login)

	// Start server
	e.Logger.Fatal(e.Start(":1546"))
	//-------------
}

// Handler
func hello(e echo.Context) error {
	return e.String(http.StatusOK, "Hellso, World!")
}
