package v1

import (
	"net/http"

	"github.com/Gavazn/Gavazn/internal/user"
	"github.com/labstack/echo"
)

type userForm struct {
	Name      string `json:"name" form:"name"`
	About     string `json:"about" form:"about"`
	Thumbnail string `json:"thumbnail" form:"thumbnail"`
}

/**
 * @api {get} /api/v1/profile get current user
 * @apiVersion 1.0.0
 * @apiName getProfile
 * @apiGroup User
 *
 * @apiSuccess {Object} user user object
 *
 * @apiError {String} error api error message
 */
func getProfile(ctx echo.Context) error {
	u := ctx.Get("user").(*user.User)

	return ctx.JSON(http.StatusOK, echo.Map{
		"user": u,
	})
}

/**
 * @api {put} /api/v1/profile edit current user
 * @apiVersion 1.0.0
 * @apiName editProfile
 * @apiGroup User
 *
 * @apiParam {String} name name
 * @apiParam {String} about about
 * @apiParam {String} thumbnail thumbnail
 *
 * @apiSuccess {String} message success message.
 * @apiSuccess {Object} user user model
 *
 * @apiError {String} error api error message
 */
func editProfile(ctx echo.Context) error {
	u := ctx.Get("user").(*user.User)

	form := new(userForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	u.Name = form.Name
	u.About = form.About
	u.Thumbnail = form.Thumbnail

	if err := u.Save(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "user updated successfully",
		"user":    u,
	})
}
