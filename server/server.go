package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Start web server
func Start(port string) error {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.CORS())

	return e.Start(fmt.Sprintf(":%v", port))
}
