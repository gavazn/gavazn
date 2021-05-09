package user

import (
	"errors"
	"net/http"

	"github.com/Gavazn/Gavazn/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jeyem/passwd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LoadByEmail load by email
func LoadByEmail(email string) (*User, error) {
	return FindOne(bson.M{"email": email})
}

// AuthByEmail authenticate with email
func AuthByEmail(email, password string) (*User, error) {
	autherr := errors.New("email or password not matched")

	u, err := LoadByEmail(email)
	if err != nil {
		return nil, autherr
	}

	if ok := passwd.Check(password, u.Password); !ok {
		return nil, autherr
	}

	return u, nil
}

// LoadByRequest load user from reqeust
func LoadByRequest(req *http.Request) (*User, error) {
	token, err := utils.ParseToken(req)
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	id, err := primitive.ObjectIDFromHex(claims["jti"].(string))
	if err != nil {
		return nil, err
	}

	return FindOne(bson.M{"_id": id})
}
