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

	postGroup := r.Group("posts")
	postGroup.GET("", listPosts)
	postGroup.POST("", addPost)
	postGroup.PUT("/:id", editPost)
	postGroup.GET("/:id", getPost)
}
