package v1

import (
	"net/http"
	"strconv"

	"github.com/Gavazn/Gavazn/internal/media"
	"github.com/Gavazn/Gavazn/internal/user"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mediaToJSON(m media.Media) bson.M {
	u, _ := user.FindOne(bson.M{"_id": m.User})

	return bson.M{
		"id":         m.ID.Hex(),
		"user":       u,
		"name":       m.Name,
		"paths":      m.Paths,
		"type":       m.Type,
		"size":       m.Size,
		"created_at": m.CreatedAt,
	}
}

/**
 * @api {post} /api/v1/medias add media
 * @apiVersion 1.0.0
 * @apiName addMedia
 * @apiGroup Media
 *
 * @apiParam {File} file media file
 *
 * @apiSuccess {String} message success message.
 * @apiSuccess {Object} media media model
 *
 * @apiError {String} error api error message
 */
func addMedia(ctx echo.Context) error {
	u := ctx.Get("user").(*user.User)

	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	paths, err := media.UploadFile(file)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	m := &media.Media{
		User:  u.ID,
		Name:  file.Filename,
		Paths: *paths,
		Type:  file.Header.Get("Content-Type"),
		Size:  file.Size,
	}

	if err := m.Save(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "media created successfully",
		"media":   mediaToJSON(*m),
	})
}

/**
 * @api {get} /api/v1/medias/:id get a media
 * @apiVersion 1.0.0
 * @apiName getMedia
 * @apiGroup Media
 *
 * @apiSuccess {Object} media media model
 *
 * @apiError {String} error api error message
 */
func getMedia(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	m, err := media.FindOne(bson.M{"_id": id})
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"media": mediaToJSON(*m),
	})
}

/**
 * @api {delete} /api/v1/medias/:id remove media
 * @apiVersion 1.0.0
 * @apiName removeMedia
 * @apiGroup Media
 *
 * @apiSuccess {String} message success message.
 *
 * @apiError {String} error api error message
 */
func removeMedia(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	m, err := media.FindOne(bson.M{"_id": id})
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	if err := m.Delete(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// remove file from system
	if err := m.Paths.DeleteFile(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "media removed successfully",
	})
}

/**
 * @api {get} /api/v1/medias list of medias
 * @apiVersion 1.0.0
 * @apiName listMedias
 * @apiGroup Media
 *
 * @apiParam {Number} page list page
 * @apiParam {Number} limit list limit
 * @apiParam {String} sort sort list example -created,...
 *
 * @apiSuccess {Number} page page number
 * @apiSuccess {Number} total_count total number of results
 * @apiSuccess {Object[]} medias array of media model
 *
 * @apiError {String} error api error message
 */
func listMedias(ctx echo.Context) error {
	filter := bson.M{}

	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	medias := media.Find(filter, page, limit, ctx.Get("sort").(bson.D)...)
	count := media.Count(filter)

	response := []bson.M{}
	for _, m := range medias {
		response = append(response, mediaToJSON(m))
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"medias":      response,
		"page":        page,
		"total_count": count,
	})
}
