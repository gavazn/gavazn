package settings

import (
	"context"

	"github.com/Gavazn/Gavazn/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Setting model
type Setting struct {
	ID          primitive.ObjectID `bson:"_id" json:"-"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Logo        primitive.ObjectID `bson:"logo" json:"logo"`
}

func (s *Setting) collection() *mongo.Collection {
	return database.Connection.Collection("settings")
}

func (s *Setting) insert() error {
	s.ID = primitive.NewObjectID()
	_, err := s.collection().InsertOne(context.Background(), database.Bson(s))
	return err
}

func (s *Setting) update() error {
	_, err := s.collection().UpdateOne(context.Background(), bson.M{}, bson.M{"$set": database.Bson(s)})
	return err
}

func findOne() (*Setting, error) {
	s := new(Setting)
	if err := s.collection().FindOne(context.Background(), bson.M{}).Decode(s); err != nil {
		return nil, err
	}
	return s, nil
}

// Set save setting data
func (s *Setting) Set() error {
	set, err := findOne()
	if err != nil {
		return s.insert()
	}

	s.ID = set.ID

	return s.update()
}

// Get get setting data
func Get() (*Setting, error) {
	s, err := findOne()
	if err != nil {
		s = &Setting{}
		if err := s.insert(); err != nil {
			return nil, err
		}
	}

	return s, nil
}
