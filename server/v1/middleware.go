package v1

import (
	"net/http"
	"strings"

	"github.com/Gavazn/Gavazn/internal/user"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func checkSorts(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		sorts := strings.Split(ctx.QueryParam("sort"), ",")
		sort := bson.D{}
		for _, s := range sorts {
			if s == "" {
				continue
			}

			value := 1
			if strings.Contains(s, "-") {
				value = -1
			}

			key := strings.Trim(s, "-")
			sort = append(sort, bson.E{Key: key, Value: value})
		}

		ctx.Set("sort", sort)
		return next(ctx)
	}
}

func setUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		u, err := user.LoadByRequest(ctx.Request())
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "loading user from token " + err.Error()})
		}

		ctx.Set("user", u)
		return next(ctx)
	}
}
