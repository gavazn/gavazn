package v1

import (
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
func register(ctx echo.Context) error {
	form := new(registerForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := ctx.Validate(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if form.Password != form.RepeatPassword {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "password not equal with repeat password"})
	}

	email := strings.ToLower(form.Email)

	if _, err := user.LoadByEmail(email); err == nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "this email has already been registered"})
	}

	u := &user.User{
		Name:      form.Name,
		About:     "",
		Email:     email,
		Password:  passwd.Make(form.Password),
		SuperUser: false,
	}

	if err := u.Save(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	t, err := utils.CreateToken(u.ID.Hex())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "successfully registered",
		"token":   t,
		"user":    userToJSON(*u),
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
func login(ctx echo.Context) error {
	form := new(loginForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := ctx.Validate(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	u, err := user.AuthByEmail(strings.ToLower(form.Email), form.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	t, err := utils.CreateToken(u.ID.Hex())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "successfully login",
		"token":   t,
		"user":    userToJSON(*u),
	})
}
