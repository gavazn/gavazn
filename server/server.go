package server

import (
	"fmt"

	v1 "github.com/Gavazn/Gavazn/server/v1"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type customValidator struct {
	validator *validator.Validate
}

// Validate forms
func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Start web server
func Start(port string) error {
	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
	e.Use(middleware.Logger(), middleware.CORS())
	e.Static("/uploads", "./uploads")

	v1.Register(e)

	return e.Start(fmt.Sprintf(":%v", port))
}
