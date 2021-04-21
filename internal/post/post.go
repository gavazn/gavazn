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
	Category  []primitive.ObjectID `bson:"_category" json:"category"`
	Tags      []string             `bson:"tags" json:"tags"`
	Thumbnail Thumbnail            `bson:"thumbnail" json:"thumbnail"`
	CreatedAt time.Time            `bson:"created_at" json:"created_at"`
}

// Thumbnail model
type Thumbnail struct {
	Original string `bson:"original" json:"original" form:"original"`
	Large    string `bson:"large" json:"large" form:"large"`
	Medium   string `bson:"medium" json:"medium" form:"medium"`
	Small    string `bson:"small" json:"small" form:"small"`
}

func (p *Post) collection() *mongo.Collection {
	return database.Connection.Collection("posts")
}

// Count posts
func Count(filter bson.M) int {
	p := new(Post)
	count, _ := p.collection().CountDocuments(context.Background(), filter)
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

	p := new(Post)
	cursor, err := p.collection().Find(context.Background(), filter, opt)
	if err != nil {
		return nil
	}

	posts := []Post{}
	for cursor.Next(context.Background()) {
		p := new(Post)
		if err := cursor.Decode(p); err != nil {
			continue
		}

		posts = append(posts, *p)
	}

	return posts
}

// FindOne post
func FindOne(filter bson.M) (*Post, error) {
	p := new(Post)
	if err := p.collection().FindOne(context.Background(), filter).Decode(p); err != nil {
		return nil, err
	}

	return p, nil
}

// Insert new Post
func (p *Post) Insert() error {
	p.ID = primitive.NewObjectID()
	p.CreatedAt = time.Now()

	_, err := p.collection().InsertOne(context.Background(), database.Bson(p))
	return err
}

// Update a Post
func (p *Post) Update() error {
	_, err := p.collection().UpdateOne(context.Background(), bson.M{"_id": p.ID}, bson.M{"$set": database.Bson(p)})
	return err
}

// Save a Post insert or update
func (p *Post) Save() error {
	if p.ID.IsZero() {
		return p.Insert()
	}

	return p.Update()
}

// Delete a post
func (p *Post) Delete() error {
	_, err := p.collection().DeleteOne(context.Background(), bson.M{"_id": p.ID})
	return err
}
