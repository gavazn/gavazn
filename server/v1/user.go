package v1

import (
	"net/http"

	"github.com/Gavazn/Gavazn/internal/user"
	"github.com/jeyem/passwd"
	"github.com/labstack/echo"
)

type userForm struct {
	Name      string `json:"name" form:"name"`
	About     string `json:"about" form:"about"`
	Thumbnail string `json:"thumbnail" form:"thumbnail"`
}

type changePasswordForm struct {
	OldPassword    string `json:"old_password" form:"old_password" validate:"required"`
	NewPassword    string `json:"new_password" form:"new_password" validate:"required"`
	RepeatPassword string `json:"repeat_password" form:"repeat_password" validate:"required"`
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

/**
 * @api {patch} /api/v1/profile/change-password change password for current user
 * @apiVersion 1.0.0
 * @apiName changePassword
 * @apiGroup User
 *
 * @apiParam {String} old_password old password
 * @apiParam {String} new_password new password
 * @apiParam {String} repeat_password repeat password
 *
 * @apiSuccess {String} message success message.
 *
 * @apiError {String} error api error message
 */
func changePassword(ctx echo.Context) error {
	u := ctx.Get("user").(*user.User)

	form := new(changePasswordForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := ctx.Validate(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if ok := passwd.Check(form.OldPassword, u.Password); !ok {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "old password not matched"})
	}

	if form.NewPassword != form.RepeatPassword {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "repeat password not matched with new password"})
	}

	u.Password = passwd.Make(form.NewPassword)

	if err := u.Save(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "password changed successfully",
	})
}
