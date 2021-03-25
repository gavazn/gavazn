package user

import (
	"errors"

	"github.com/jeyem/passwd"
	"go.mongodb.org/mongo-driver/bson"
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
