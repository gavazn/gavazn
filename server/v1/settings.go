package v1

import (
	"net/http"

	"github.com/Gavazn/Gavazn/internal/settings"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type settingForm struct {
	Title       string             `json:"title" form:"title"`
	Description string             `json:"description" form:"description"`
	Logo        primitive.ObjectID `json:"logo" form:"logo"`
}

/**
 * @api {get} /api/v1/settings get setting
 * @apiVersion 1.0.0
 * @apiName getSetting
 * @apiGroup Setting
 *
 * @apiSuccess {Object} settings settings model
 *
 * @apiError {String} error message
 */
func getSetting(ctx echo.Context) error {
	s, err := settings.Get()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"settings": s,
	})
}

/**
 * @api {put} /api/v1/settings set setting
 * @apiVersion 1.0.0
 * @apiName setSetting
 * @apiGroup Setting
 *
 * @apiSuccess {Object} settings setting model
 *
 * @apiError {String} error message
 */
func setSetting(ctx echo.Context) error {
	form := new(settingForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	s := &settings.Setting{
		Title:       form.Title,
		Description: form.Description,
		Logo:        form.Logo,
	}

	if err := s.Set(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"setting": s,
	})
}
