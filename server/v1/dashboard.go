package v1

import (
	"github.com/Gavazn/Gavazn/internal/category"
	"github.com/Gavazn/Gavazn/internal/comment"
	"github.com/Gavazn/Gavazn/internal/post"
	"github.com/Gavazn/Gavazn/internal/user"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

/**
 * @api {get} /api/v2/dashboard get statistics
 * @apiVersion 1.0.0
 * @apiName getStatistics
 * @apiGroup Dashboard
 *
 * @apiSuccess {Number} total_posts
 * @apiSuccess {Number} total_categories
 * @apiSuccess {Number} total_comments
 * @apiSuccess {Number} total_users
 *
 * @apiError {String} error api error message
 */
func getStatistics(ctx echo.Context)error{
	return ctx.JSON(200, echo.Map{
		"total_posts":    post.Count(bson.M{}),
		"total_categories":  category.Count(bson.M{}),
		"total_comments":   comment.Count(bson.M{}),
		"total_users":   user.Count(bson.M{}),
	})
}