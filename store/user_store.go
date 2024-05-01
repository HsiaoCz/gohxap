package store

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context) error
}

type UserMongoStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewUserMongoStore(client *mongo.Client, dbname string, coll string) *UserMongoStore {
	return &UserMongoStore{
		client: client,
		coll:   client.Database(dbname).Collection(coll),
	}
}

func (u *UserMongoStore) CreateUser(ctx context.Context) error {
	return nil
}
