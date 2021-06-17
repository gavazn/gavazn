package v1

import (
	"net/http"
	"strconv"

	"github.com/Gavazn/Gavazn/internal/comment"
	"github.com/Gavazn/Gavazn/internal/post"
	"github.com/Gavazn/Gavazn/internal/user"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commentForm struct {
	Content string `json:"content" form:"content"`
}

func commentToJSON(c comment.Comment) bson.M {
	u, _ := user.FindOne(bson.M{"_id": c.User})
	p, _ := post.FindOne(bson.M{"_id": c.Post})

	return bson.M{
		"id":         c.ID.Hex(),
		"user":       u,
		"post":       p,
		"content":    c.Content,
		"created_at": c.CreatedAt,
	}
}

/**
 * @api {post} /api/v1/posts/:id/comments add comment
 * @apiVersion 1.0.0
 * @apiName addComment
 * @apiGroup Comment
 *
 * @apiParam {String} content content
 *
 * @apiSuccess {String} message success message.
 * @apiSuccess {Object} comment comment model
 *
 * @apiError {String} error api error message
 */
func addComment(ctx echo.Context) error {
	u := ctx.Get("user").(*user.User)

	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	form := new(commentForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	c := &comment.Comment{
		User:    u.ID,
		Post:    id,
		Content: form.Content,
	}

	if err := c.Save(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "comment created successfully",
		"comment": commentToJSON(*c),
	})
}

/**
 * @api {delete} /api/v1/commnets/:id remove commnet
 * @apiVersion 1.0.0
 * @apiName removeComment
 * @apiGroup Comment
 *
 * @apiSuccess {String} message success message.
 *
 * @apiError {String} error api error message
 */
func removeComment(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	c, err := comment.FindOne(bson.M{"_id": id})
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	if err := c.Delete(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "comment removed successfully",
	})
}

/**
 * @api {get} /api/v1/comments list of comments
 * @apiVersion 1.0.0
 * @apiName listComments
 * @apiGroup Comment
 *
 * @apiParam {String} q query search
 * @apiParam {String} post post object id
 * @apiParam {Number} page list page
 * @apiParam {Number} limit list limit
 * @apiParam {String} sort sort list example -created,title,...
 *
 * @apiSuccess {Number} page page number
 * @apiSuccess {Number} total_count total number of results
 * @apiSuccess {Object[]} comments array of comment model
 *
 * @apiError {String} error api error message
 */
func listComments(ctx echo.Context) error {
	filter := bson.M{}

	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	if q := ctx.QueryParam("q"); q != "" {
		filter["$text"] = bson.M{"$search": q}
	}

	if postID, _ := primitive.ObjectIDFromHex(ctx.QueryParam("post")); !postID.IsZero() {
		filter["_post"] = postID
	}

	comments := comment.Find(filter, page, limit, ctx.Get("sort").(bson.D)...)
	count := comment.Count(filter)

	response := []bson.M{}
	for _, c := range comments {
		response = append(response, commentToJSON(c))
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"comments":    response,
		"page":        page,
		"total_count": count,
	})
}
