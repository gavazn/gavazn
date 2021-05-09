package v1

import (
	"net/http"
	"strings"

	"github.com/Gavazn/Gavazn/internal/user"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func checkSorts(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sorts := strings.Split(c.QueryParam("sort"), ",")
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

		c.Set("sort", sort)
		return next(c)
	}
}

func setUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		u, err := user.LoadByRequest(c.Request())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "loading user from token " + err.Error()})
		}

		c.Set("user", u)
		return next(c)
	}
}
