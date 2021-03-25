package v1

import (
	"errors"
	"net/http"
	"strings"

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

type loginForm struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

/**
 * @api {post} /api/v1/auth/register register user
 * @apiVersion 1.0.0
 * @apiName register
 * @apiGroup Auth
 *
 * @apiParam {String} name name
 * @apiParam {String} email email address
 * @apiParam {String} password password
 * @apiParam {String} repeat_password repeat password
 *
 * @apiSuccess {String} message message
 * @apiSuccess {String} token user token
 * @apiSuccess {Object} user user model
 *
 * @apiError {String} error message
 */
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

	email := strings.ToLower(form.Email)

	if _, err := user.LoadByEmail(email); err == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": errors.New("this email has already been registered").Error()})
	}

	u := &user.User{
		Name:      form.Name,
		About:     "",
		Email:     email,
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

/**
 * @api {post} /api/v1/auth/login login user
 * @apiVersion 1.0.0
 * @apiName login
 * @apiGroup Auth
 *
 * @apiParam {String} email email address
 * @apiParam {String} password password
 *
 * @apiSuccess {String} message message
 * @apiSuccess {String} token user token
 * @apiSuccess {Object} user user model
 *
 * @apiError {String} error message
 */
func login(c echo.Context) error {
	form := new(loginForm)
	if err := c.Bind(form); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(form); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	u, err := user.AuthByEmail(strings.ToLower(form.Email), form.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	t, err := utils.CreateToken(u.ID.Hex())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "successfully login",
		"token":   t,
		"user":    u,
	})
}
