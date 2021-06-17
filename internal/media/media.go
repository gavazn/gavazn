package media

import (
	"context"
	"time"

	"github.com/Gavazn/Gavazn/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Media model
type Media struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	User      primitive.ObjectID `bson:"_user" json:"user"`
	Name      string             `bson:"name" json:"name"`
	Paths     Path               `bson:"paths" json:"paths"`
	Type      string             `bson:"type" json:"type"`
	Size      int64              `bson:"size" json:"size"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

func (m *Media) collection() *mongo.Collection {
	return database.Connection.Collection("medias")
}

// Count medias
func Count(filter bson.M) int {
	m := new(Media)
	count, _ := m.collection().CountDocuments(context.Background(), filter)
	return int(count)
}

// Find medias
func Find(filter bson.M, page, limit int, sorts ...bson.E) []Media {
	var sort bson.D
	if len(sorts) < 1 {
		sort = append(sort, bson.E{Key: "created_at", Value: -1})
	}

	sort = append(sort, sorts...)
	opt := options.Find()
	opt.SetSort(sort)

	if limit > 0 {
		opt.SetLimit(int64(limit))
	}

	if page > 1 {
		opt.SetSkip(int64((page - 1) * limit))
	}

	m := new(Media)
	cursor, err := m.collection().Find(context.Background(), filter, opt)
	if err != nil {
		return nil
	}

	medias := []Media{}
	for cursor.Next(context.Background()) {
		m := new(Media)
		if err := cursor.Decode(m); err != nil {
			continue
		}

		medias = append(medias, *m)
	}

	return medias
}

// FindOne media
func FindOne(filter bson.M) (*Media, error) {
	m := new(Media)
	if err := m.collection().FindOne(context.Background(), filter).Decode(m); err != nil {
		return nil, err
	}

	return m, nil
}

// Insert new Media
func (m *Media) Insert() error {
	m.ID = primitive.NewObjectID()
	m.CreatedAt = time.Now()

	_, err := m.collection().InsertOne(context.Background(), database.Bson(m))
	return err
}

// Update a Media
func (m *Media) Update() error {
	_, err := m.collection().UpdateOne(context.Background(), bson.M{"_id": m.ID}, bson.M{"$set": database.Bson(m)})
	return err
}

// Save a Media insert or update
func (m *Media) Save() error {
	if m.ID.IsZero() {
		return m.Insert()
	}

	return m.Update()
}

// Delete a media
func (m *Media) Delete() error {
	_, err := m.collection().DeleteOne(context.Background(), bson.M{"_id": m.ID})
	return err
}
