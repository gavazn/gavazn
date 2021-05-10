package comment

import (
	"context"
	"time"

	"github.com/Gavazn/Gavazn/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Comment Model
type Comment struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	User      primitive.ObjectID `bson:"_user" json:"user"`
	Post      primitive.ObjectID `bson:"_post" json:"post"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

func (c *Comment) collection() *mongo.Collection {
	return database.Connection.Collection("comments")
}

// Count comments
func Count(filter bson.M) int {
	c := new(Comment)
	count, _ := c.collection().CountDocuments(context.Background(), filter)
	return int(count)
}

// Find comments
func Find(filter bson.M, page, limit int, sorts ...bson.E) []Comment {
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

	c := new(Comment)
	cursor, err := c.collection().Find(context.Background(), filter, opt)
	if err != nil {
		return nil
	}

	comments := []Comment{}
	for cursor.Next(context.Background()) {
		c := new(Comment)
		if err := cursor.Decode(c); err != nil {
			continue
		}

		comments = append(comments, *c)
	}

	return comments
}

// FindOne comment
func FindOne(filter bson.M) (*Comment, error) {
	c := new(Comment)
	if err := c.collection().FindOne(context.Background(), filter).Decode(c); err != nil {
		return nil, err
	}

	return c, nil
}

// Insert new Comment
func (c *Comment) Insert() error {
	c.ID = primitive.NewObjectID()
	c.CreatedAt = time.Now()

	_, err := c.collection().InsertOne(context.Background(), database.Bson(c))
	return err
}

// Update a Comment
func (c *Comment) Update() error {
	_, err := c.collection().UpdateOne(context.Background(), bson.M{"_id": c.ID}, bson.M{"$set": database.Bson(c)})
	return err
}

// Save a Comment insert or update
func (c *Comment) Save() error {
	if c.ID.IsZero() {
		return c.Insert()
	}

	return c.Update()
}

// Delete a comment
func (c *Comment) Delete() error {
	_, err := c.collection().DeleteOne(context.Background(), bson.M{"_id": c.ID})
	return err
}
