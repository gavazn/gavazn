package utils

import (
	"net/http"
	"strings"
	"time"

	"github.com/Gavazn/Gavazn/config"
	"github.com/dgrijalva/jwt-go"
)

const expireDuration = time.Hour * 72

// CreateToken creating token
func CreateToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["jti"] = id
	claims["exp"] = time.Now().Add(expireDuration).Unix()

	return token.SignedString([]byte(config.Get("SECRET_KEY")))
}

// GetToken get user token
func GetToken(req *http.Request) string {
	cleared := strings.Replace(req.Header.Get("Authorization"), " ", "", -1)
	return strings.Replace(cleared, "Bearer", "", -1)
}
