package database

import (
	"context"
	"fmt"

	"github.com/MEDALIALPHA331/reservation/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollection = "user"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *models.User) (map[string]string, error)
	// GetUserById(id string) (*models.User, error)
	//..
}

type MongoDbUserStore struct {
	client *mongo.Client
	coll   mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoDbUserStore {
	mongodb := MongoDbUserStore{
		client: client,
		coll:   *client.Database(DBNAME).Collection(userCollection),
	}
	return &mongodb
}

func (c *MongoDbUserStore) CreateUser(ctx context.Context, user *models.User) (map[string]string, error) {
	result, err := c.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"message": fmt.Sprintf("User %s created", result.InsertedID),
	}, nil
}
