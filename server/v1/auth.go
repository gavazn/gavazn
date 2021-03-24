package v1

import (
	"errors"
	"net/http"

	"github.com/Gavazn/Gavazn/internal/user"
	"github.com/Gavazn/Gavazn/utils"
	"github.com/jeyem/passwd"
	"github.com/labstack/echo"
)

type registerForm struct {
	Name           string `json:"name" form:"name"`
	Email          string `json:"email" form:"email" validate:"required,email"`
	Password       string `json:"password" form:"password" validate:"required"`
	RepeatPassword string `json:"repeat_password" form:"repeat_password" validate:"required"`
}

func register(c echo.Context) error {
	form := new(registerForm)
	if err := c.Bind(form); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(form); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if form.Password != form.RepeatPassword {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": errors.New("password not equal with repeat password").Error()})
	}

	u := &user.User{
		Name:      form.Name,
		About:     "",
		Email:     form.Email,
		Password:  passwd.Make(form.Password),
		SuperUser: false,
		Thumbnail: "",
	}

	if err := u.Save(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	t, err := utils.CreateToken(u.ID.Hex())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "successfully registered",
		"token":   t,
		"user":    u,
	})
}
