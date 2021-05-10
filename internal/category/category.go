package category

import (
	"context"
	"time"

	"github.com/Gavazn/Gavazn/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Category Model
type Category struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	User      primitive.ObjectID `bson:"_user" json:"user"`
	Parent    primitive.ObjectID `bson:"_parent" json:"parent"`
	Name      string             `bson:"name" json:"name"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

func (c *Category) collection() *mongo.Collection {
	return database.Connection.Collection("categories")
}

// Count categories
func Count(filter bson.M) int {
	c := new(Category)
	count, _ := c.collection().CountDocuments(context.Background(), filter)
	return int(count)
}

// Find categories
func Find(filter bson.M, page, limit int, sorts ...bson.E) []Category {
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

	c := new(Category)
	cursor, err := c.collection().Find(context.Background(), filter, opt)
	if err != nil {
		return nil
	}

	categories := []Category{}
	for cursor.Next(context.Background()) {
		c := new(Category)
		if err := cursor.Decode(c); err != nil {
			continue
		}

		categories = append(categories, *c)
	}

	return categories
}

// FindOne category
func FindOne(filter bson.M) (*Category, error) {
	c := new(Category)
	if err := c.collection().FindOne(context.Background(), filter).Decode(c); err != nil {
		return nil, err
	}

	return c, nil
}

// Insert new Category
func (c *Category) Insert() error {
	c.ID = primitive.NewObjectID()
	c.CreatedAt = time.Now()

	_, err := c.collection().InsertOne(context.Background(), database.Bson(c))
	return err
}

// Update a Category
func (c *Category) Update() error {
	_, err := c.collection().UpdateOne(context.Background(), bson.M{"_id": c.ID}, bson.M{"$set": database.Bson(c)})
	return err
}

// Save a Category insert or update
func (c *Category) Save() error {
	if c.ID.IsZero() {
		return c.Insert()
	}

	return c.Update()
}

// Delete a category
func (c *Category) Delete() error {
	_, err := c.collection().DeleteOne(context.Background(), bson.M{"_id": c.ID})
	return err
}
