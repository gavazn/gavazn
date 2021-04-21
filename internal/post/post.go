package post

import (
	"context"
	"time"

	"github.com/Gavazn/Gavazn/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Post Model
type Post struct {
	ID        primitive.ObjectID   `bson:"_id" json:"id"`
	Author    primitive.ObjectID   `bson:"_author" json:"author"`
	Title     string               `bson:"title" json:"title"`
	Content   string               `bson:"content" json:"content"`
	Category  []primitive.ObjectID `bson:"category" json:"category"`
	Tags      []string             `bson:"tags" json:"tags"`
	Thumbnail thumbnail            `bson:"thumbnail" json:"thumbnail"`
	CreatedAt time.Time            `bson:"created_at" json:"created_at"`
}

type thumbnail struct {
	Original string `bson:"original" json:"original"`
	Large    string `bson:"large" json:"large"`
	Medium   string `bson:"medium" json:"medium"`
	Small    string `bson:"small" json:"small"`
}

func (u *Post) collection() *mongo.Collection {
	return database.Connection.Collection("posts")
}

// Count posts
func Count(filter bson.M) int {
	u := new(Post)
	count, _ := u.collection().CountDocuments(context.Background(), filter)
	return int(count)
}

// Find posts
func Find(filter bson.M, page, limit int, sorts ...bson.E) []Post {
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

	u := new(Post)
	cursor, err := u.collection().Find(context.Background(), filter, opt)
	if err != nil {
		return nil
	}

	posts := []Post{}
	for cursor.Next(context.Background()) {
		u := new(Post)
		if err := cursor.Decode(u); err != nil {
			continue
		}

		posts = append(posts, *u)
	}

	return posts
}

// FindOne post
func FindOne(filter bson.M) (*Post, error) {
	u := new(Post)
	if err := u.collection().FindOne(context.Background(), filter).Decode(u); err != nil {
		return nil, err
	}

	return u, nil
}

// Insert new Post
func (u *Post) Insert() error {
	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now()

	_, err := u.collection().InsertOne(context.Background(), database.Bson(u))
	return err
}

// Update a Post
func (u *Post) Update() error {
	_, err := u.collection().UpdateOne(context.Background(), bson.M{"_id": u.ID}, bson.M{"$set": database.Bson(u)})
	return err
}

// Save a Post insert or update
func (u *Post) Save() error {
	if u.ID.IsZero() {
		return u.Insert()
	}

	return u.Update()
}

// Delete a post
func (u *Post) Delete() error {
	_, err := u.collection().DeleteOne(context.Background(), bson.M{"_id": u.ID})
	return err
}
