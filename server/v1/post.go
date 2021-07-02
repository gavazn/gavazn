package v1

import (
	"net/http"
	"strconv"

	"github.com/Gavazn/Gavazn/internal/category"
	"github.com/Gavazn/Gavazn/internal/media"
	"github.com/Gavazn/Gavazn/internal/post"
	"github.com/Gavazn/Gavazn/internal/user"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type postForm struct {
	Title      string               `json:"title" form:"title"`
	Content    string               `json:"content" form:"content"`
	Categories []primitive.ObjectID `json:"categories" form:"categories"`
	Tags       []string             `json:"tags" form:"tags"`
	Thumbnail  primitive.ObjectID   `json:"thumbnail" form:"thumbnail"`
}

func postToJSON(p post.Post) bson.M {
	u, err := user.FindOne(bson.M{"_id": p.User})
	if err != nil {
		u = new(user.User)
	}

	categories := category.Find(bson.M{"_id": bson.M{"$in": p.Categories}}, 1, -1)
	t, _ := media.FindOne(bson.M{"_id": p.Thumbnail})

	return bson.M{
		"id":         p.ID.Hex(),
		"user":       userToJSON(*u),
		"title":      p.Title,
		"content":    p.Content,
		"categories": categories,
		"tags":       p.Tags,
		"thumbnail":  t,
		"created_at": p.CreatedAt,
	}
}

/**
 * @api {post} /api/v1/posts add post
 * @apiVersion 1.0.0
 * @apiName addPost
 * @apiGroup Post
 *
 * @apiParam {String} title post title
 * @apiParam {String} content post content
 * @apiParam {String[]} categories list of category id
 * @apiParam {String[]} tags list of tag
 * @apiParam {String} thumbnail thumbnail id
 *
 * @apiSuccess {String} message success message.
 * @apiSuccess {Object} post post model
 *
 * @apiError {String} error api error message
 */
func addPost(ctx echo.Context) error {
	u := ctx.Get("user").(*user.User)

	form := new(postForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	p := &post.Post{
		User:       u.ID,
		Title:      form.Title,
		Content:    form.Content,
		Categories: form.Categories,
		Tags:       form.Tags,
		Thumbnail:  form.Thumbnail,
	}

	if err := p.Save(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "post created successfully",
		"post":    postToJSON(*p),
	})
}

/**
 * @api {put} /api/v1/posts/:id edit post
 * @apiVersion 1.0.0
 * @apiName editPost
 * @apiGroup Post
 *
 * @apiParam {String} title post title
 * @apiParam {String} content post content
 * @apiParam {String[]} categories list of category id
 * @apiParam {String[]} tags list of tag
 * @apiParam {String} thumbnail thumbnail id
 *
 * @apiSuccess {String} message success message.
 * @apiSuccess {Object} post post model
 *
 * @apiError {String} error api error message
 */
func editPost(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	form := new(postForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	p, err := post.FindOne(bson.M{"_id": id})
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	p.Title = form.Title
	p.Content = form.Content
	p.Categories = form.Categories
	p.Tags = form.Tags
	p.Thumbnail = form.Thumbnail

	if err := p.Save(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "post updated successfully",
		"post":    postToJSON(*p),
	})
}

/**
 * @api {get} /api/v1/posts/:id get a post
 * @apiVersion 1.0.0
 * @apiName getPost
 * @apiGroup Post
 *
 * @apiSuccess {Object} post post model
 *
 * @apiError {String} error api error message
 */
func getPost(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	p, err := post.FindOne(bson.M{"_id": id})
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"post": postToJSON(*p),
	})
}

/**
 * @api {delete} /api/v1/posts/:id remove post
 * @apiVersion 1.0.0
 * @apiName removePost
 * @apiGroup Post
 *
 * @apiSuccess {String} message success message.
 *
 * @apiError {String} error api error message
 */
func removePost(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	p, err := post.FindOne(bson.M{"_id": id})
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	if err := p.Delete(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "post removed successfully",
	})
}

/**
 * @api {get} /api/v1/posts list of posts
 * @apiVersion 1.0.0
 * @apiName listPosts
 * @apiGroup Post
 *
 * @apiParam {String} q query search
 * @apiParam {String} user user id
 * @apiParam {String} category category id
 * @apiParam {Number} page list page
 * @apiParam {Number} limit list limit
 * @apiParam {String} sort sort list example -created,title,...
 *
 * @apiSuccess {Number} page page number
 * @apiSuccess {Number} total_count total number of results
 * @apiSuccess {Object[]} posts array of post model
 *
 * @apiError {String} error api error message
 */
func listPosts(ctx echo.Context) error {
	filter := bson.M{}

	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	if q := ctx.QueryParam("q"); q != "" {
		filter["$text"] = bson.M{"$search": q}
	}

	if userID, _ := primitive.ObjectIDFromHex(ctx.QueryParam("user")); !userID.IsZero() {
		filter["_user"] = userID
	}

	if categoryID, _ := primitive.ObjectIDFromHex(ctx.QueryParam("category")); !categoryID.IsZero() {
		filter["_categories"] = categoryID
	}

	posts := post.Find(filter, page, limit, ctx.Get("sort").(bson.D)...)
	count := post.Count(filter)

	response := []bson.M{}
	for _, p := range posts {
		response = append(response, postToJSON(p))
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"posts":       response,
		"page":        page,
		"total_count": count,
	})
}
