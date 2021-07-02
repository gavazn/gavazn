package v1

import (
	"net/http"
	"strconv"

	"github.com/Gavazn/Gavazn/internal/category"
	"github.com/Gavazn/Gavazn/internal/media"
	"github.com/Gavazn/Gavazn/internal/user"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type categoryForm struct {
	Parent    primitive.ObjectID `json:"parent" form:"parent"`
	Name      string             `json:"name" form:"name"`
	Thumbnail primitive.ObjectID `json:"thumbnail" form:"thumbnail"`
}

func categoryToJSON(c category.Category) bson.M {
	u, _ := user.FindOne(bson.M{"_id": c.User})
	parent, _ := category.FindOne(bson.M{"_id": c.Parent})
	t, _ := media.FindOne(bson.M{"_id": c.Thumbnail})

	return bson.M{
		"id":         c.ID.Hex(),
		"user":       u,
		"parent":     parent,
		"name":       c.Name,
		"thumbnail":  t,
		"created_at": c.CreatedAt,
	}
}

/**
 * @api {post} /api/v1/categories add category
 * @apiVersion 1.0.0
 * @apiName addCategory
 * @apiGroup Category
 *
 * @apiParam {String} parent parent
 * @apiParam {String} name name
 * @apiParam {String} thumbnail thumbnail id
 *
 * @apiSuccess {String} message success message.
 * @apiSuccess {Object} category category model
 *
 * @apiError {String} error api error message
 */
func addCategory(ctx echo.Context) error {
	u := ctx.Get("user").(*user.User)

	form := new(categoryForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	c := &category.Category{
		User:      u.ID,
		Parent:    form.Parent,
		Name:      form.Name,
		Thumbnail: form.Thumbnail,
	}

	if err := c.Save(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message":  "category created successfully",
		"category": categoryToJSON(*c),
	})
}

/**
 * @api {put} /api/v1/categories/:id edit category
 * @apiVersion 1.0.0
 * @apiName editCategory
 * @apiGroup Category
 *
 * @apiParam {String} parent parent
 * @apiParam {String} name name
 * @apiParam {String} thumbnail thumbnail id
 *
 * @apiSuccess {String} message success message.
 * @apiSuccess {Object} category category model
 *
 * @apiError {String} error api error message
 */
func editCategory(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	form := new(categoryForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	c, err := category.FindOne(bson.M{"_id": id})
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	c.Parent = form.Parent
	c.Name = form.Name
	c.Thumbnail = form.Thumbnail

	if err := c.Save(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message":  "category updated successfully",
		"category": categoryToJSON(*c),
	})
}

/**
 * @api {get} /api/v1/categories/:id get a category
 * @apiVersion 1.0.0
 * @apiName getCategory
 * @apiGroup Category
 *
 * @apiSuccess {Object} category category model
 *
 * @apiError {String} error api error message
 */
func getCategory(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	c, err := category.FindOne(bson.M{"_id": id})
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"category": categoryToJSON(*c),
	})
}

/**
 * @api {delete} /api/v1/categories/:id remove category
 * @apiVersion 1.0.0
 * @apiName removeCategory
 * @apiGroup Category
 *
 * @apiSuccess {String} message success message.
 *
 * @apiError {String} error api error message
 */
func removeCategory(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	c, err := category.FindOne(bson.M{"_id": id})
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	if err := c.Delete(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "category removed successfully",
	})
}

/**
 * @api {get} /api/v1/categories list of categories
 * @apiVersion 1.0.0
 * @apiName listCategories
 * @apiGroup Category
 *
 * @apiParam {String} q query search
 * @apiParam {Number} page list page
 * @apiParam {Number} limit list limit
 * @apiParam {String} sort sort list example -created,title,...
 *
 * @apiSuccess {Number} page page number
 * @apiSuccess {Number} total_count total number of results
 * @apiSuccess {Object[]} categories array of category model
 *
 * @apiError {String} error api error message
 */
func listCategories(ctx echo.Context) error {
	filter := bson.M{}

	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	if q := ctx.QueryParam("q"); q != "" {
		filter["$text"] = bson.M{"$search": q}
	}

	categories := category.Find(filter, page, limit, ctx.Get("sort").(bson.D)...)
	count := category.Count(filter)

	response := []bson.M{}
	for _, c := range categories {
		response = append(response, categoryToJSON(c))
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"categories":  response,
		"page":        page,
		"total_count": count,
	})
}
