package v1

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Gavazn/Gavazn/internal/media"
	"github.com/Gavazn/Gavazn/internal/user"
	"github.com/jeyem/passwd"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userForm struct {
	Name      string             `json:"name" form:"name"`
	About     string             `json:"about" form:"about"`
	Thumbnail primitive.ObjectID `json:"thumbnail" form:"thumbnail"`
}

type newUserForm struct {
	Name      string             `json:"name" form:"name"`
	About     string             `json:"about" form:"about"`
	Email     string             `json:"email" form:"email" validate:"required,email"`
	Password  string             `json:"password" form:"password" validate:"required"`
	SuperUser bool               `json:"super_user" form:"super_user"`
	Thumbnail primitive.ObjectID `json:"thumbnail" form:"thumbnail"`
}

type changePasswordForm struct {
	OldPassword    string `json:"old_password" form:"old_password" validate:"required"`
	NewPassword    string `json:"new_password" form:"new_password" validate:"required"`
	RepeatPassword string `json:"repeat_password" form:"repeat_password" validate:"required"`
}

func userToJSON(u user.User) bson.M {
	t, _ := media.FindOne(bson.M{"_id": u.Thumbnail})

	return bson.M{
		"id":         u.ID.Hex(),
		"name":       u.Name,
		"about":      u.About,
		"email":      u.Email,
		"password":   u.Password,
		"super_user": u.SuperUser,
		"thumbnail":  t,
		"created_at": u.CreatedAt,
	}
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
		"user": userToJSON(*u),
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
		"user":    userToJSON(*u),
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

/**
 * @api {post} /api/v1/users add user
 * @apiVersion 1.0.0
 * @apiName addUser
 * @apiGroup User
 *
 * @apiParam {String} name name
 * @apiParam {String} about about
 * @apiParam {String} email email
 * @apiParam {String} password password
 * @apiParam {Boolean} super_user super user
 * @apiParam {String} thumbnail thumbnail
 *
 * @apiSuccess {String} message success message.
 * @apiSuccess {Object} user user model
 *
 * @apiError {String} error api error message
 */
func addUser(ctx echo.Context) error {
	form := new(newUserForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := ctx.Validate(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	email := strings.ToLower(form.Email)

	if _, err := user.LoadByEmail(email); err == nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "this email has already been registered"})
	}

	u := &user.User{
		Name:      form.Name,
		About:     form.About,
		Email:     email,
		Password:  passwd.Make(form.Password),
		SuperUser: form.SuperUser,
		Thumbnail: form.Thumbnail,
	}

	if err := u.Save(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "user created successfully",
		"user":    userToJSON(*u),
	})
}

/**
 * @api {put} /api/v1/users/:id edit user
 * @apiVersion 1.0.0
 * @apiName editUser
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
func editUser(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	form := new(userForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	u, err := user.FindOne(bson.M{"_id": id})
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	u.Name = form.Name
	u.About = form.About
	u.Thumbnail = form.Thumbnail

	if err := u.Save(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "user updated successfully",
		"user":    userToJSON(*u),
	})
}

/**
 * @api {get} /api/v1/users/:id get a user
 * @apiVersion 1.0.0
 * @apiName getUser
 * @apiGroup User
 *
 * @apiSuccess {Object} user user model
 *
 * @apiError {String} error api error message
 */
func getUser(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	u, err := user.FindOne(bson.M{"_id": id})
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"user": userToJSON(*u),
	})
}

/**
 * @api {delete} /api/v1/users/:id remove user
 * @apiVersion 1.0.0
 * @apiName removeUser
 * @apiGroup User
 *
 * @apiSuccess {String} message success message.
 *
 * @apiError {String} error api error message
 */
func removeUser(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	u, err := user.FindOne(bson.M{"_id": id})
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	if err := u.Delete(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "user removed successfully",
	})
}

/**
 * @api {get} /api/v1/users list of users
 * @apiVersion 1.0.0
 * @apiName listUsers
 * @apiGroup User
 *
 * @apiParam {String} q query search
 * @apiParam {Number} page list page
 * @apiParam {Number} limit list limit
 * @apiParam {String} sort sort list example -created,title,...
 *
 * @apiSuccess {Number} page page number
 * @apiSuccess {Number} total_count total number of results
 * @apiSuccess {Object[]} users array of user model
 *
 * @apiError {String} error api error message
 */
func listUsers(ctx echo.Context) error {
	filter := bson.M{}

	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	if q := ctx.QueryParam("q"); q != "" {
		filter["$text"] = bson.M{"$search": q}
	}

	users := user.Find(filter, page, limit, ctx.Get("sort").(bson.D)...)
	count := user.Count(filter)

	response := []bson.M{}
	for _, u := range users {
		response = append(response, userToJSON(u))
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"users":       response,
		"page":        page,
		"total_count": count,
	})
}
