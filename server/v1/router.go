package v1

import (
	"github.com/Gavazn/Gavazn/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var secretKey = []byte(config.Get("SECRET_KEY"))

// Register routes
func Register(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")
	authGroup.POST("/register", register)
	authGroup.POST("/login", login)

	r := v1.Group("/")
	r.Use(middleware.JWT(secretKey), checkSorts, setUser)

	profileGroup := r.Group("profile")
	profileGroup.GET("", getProfile)
	profileGroup.PUT("", editProfile)
	profileGroup.PATCH("/change-password", changePassword)

	userGroup := r.Group("users")
	userGroup.GET("", listUsers)
	userGroup.POST("", addUser)
	userGroup.PUT("/:id", editUser)
	userGroup.GET("/:id", getUser)
	userGroup.DELETE("/:id", removeUser)

	postGroup := r.Group("posts")
	postGroup.GET("", listPosts)
	postGroup.POST("", addPost)
	postGroup.PUT("/:id", editPost)
	postGroup.GET("/:id", getPost)
	postGroup.DELETE("/:id", removePost)
	postGroup.POST("/:id/comments", addComment)

	commentGroup := r.Group("comments")
	commentGroup.GET("", listComments)
	commentGroup.DELETE("/:id", removeComment)

	categoryGroup := r.Group("categories")
	categoryGroup.GET("", listCategories)
	categoryGroup.POST("", addCategory)
	categoryGroup.PUT("/:id", editCategory)
	categoryGroup.GET("/:id", getCategory)
	categoryGroup.DELETE("/:id", removeCategory)
}
