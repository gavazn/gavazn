package user

import (
	"context"
	"time"

	"github.com/Gavazn/Gavazn/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User model
type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	About     string             `bson:"about" json:"about"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"`
	SuperUser bool               `bson:"super_user" json:"super_user"`
	Thumbnail primitive.ObjectID `bson:"thumbnail" json:"thumbnail"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

func (u *User) collection() *mongo.Collection {
	return database.Connection.Collection("users")
}

// Insert new user
func (u *User) Insert() error {
	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now()

	_, err := u.collection().InsertOne(context.Background(), database.Bson(u))
	return err
}

// UpdateOne a user
func (u *User) UpdateOne() error {
	_, err := u.collection().UpdateOne(context.Background(), bson.M{"_id": u.ID}, bson.M{"$set": database.Bson(u)})
	return err
}

// UpdateMany users
func UpdateMany(filter, update bson.M) error {
	u := new(User)
	_, err := u.collection().UpdateMany(context.Background(), filter, update)
	return err
}

// Save a user insert or update
func (u *User) Save() error {
	if u.ID.IsZero() {
		return u.Insert()
	}
	return u.UpdateOne()
}

// DeleteOne a user
func (u *User) Delete() error {
	_, err := u.collection().DeleteOne(context.Background(), bson.M{"_id": u.ID})
	return err
}

// DeleteMany user
func DeleteMany(filter bson.M) error {
	u := new(User)
	_, err := u.collection().DeleteMany(context.Background(), filter)
	return err
}

// Count users
func Count(filter bson.M) int {
	u := new(User)
	count, _ := u.collection().CountDocuments(context.Background(), filter)
	return int(count)
}

// SetIndex create index
func SetIndex(keys bson.M, unique bool) error {
	index := mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetUnique(unique),
	}

	u := new(User)
	_, err := u.collection().Indexes().CreateOne(context.Background(), index)
	return err
}

// Drop users collection
func Drop() error {
	u := new(User)
	return u.collection().Drop(context.Background())
}

// FindOne user
func FindOne(filter bson.M) (*User, error) {
	u := new(User)
	if err := u.collection().FindOne(context.Background(), filter).Decode(u); err != nil {
		return nil, err
	}
	return u, nil
}

// Find users
func Find(filter bson.M, page, limit int, sorts ...bson.E) (users []User) {
	var sort bson.D
	opt := options.Find()
	if len(sorts) < 1 {
		sort = append(sort, bson.E{Key: "created_at", Value: -1})
	}

	sort = append(sort, sorts...)
	opt.SetSort(sort)

	if limit > 0 {
		opt.SetLimit(int64(limit))
	}

	if page > 1 {
		opt.SetSkip(int64((page - 1) * limit))
	}

	u := new(User)
	cursor, err := u.collection().Find(context.Background(), filter, opt)
	if err != nil {
		return
	}

	for cursor.Next(context.Background()) {
		u := new(User)
		if err := cursor.Decode(u); err != nil {
			continue
		}

		users = append(users, *u)
	}

	return
}
