package database

import (
	"context"
	"fmt"
	"log"

	"github.com/MEDALIALPHA331/reservation/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollection = "user"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *models.User) (map[string]string, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
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



func (c *MongoDbUserStore) GetAllUsers(ctx context.Context) ([]models.User, error) {

	cur, err := c.coll.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Could not find user, %+v", err)
		return nil, err
	}

	var users []models.User
	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatalf("Could not decode user, %+v", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := cur.Err(); err!= nil {
		log.Fatalf("Could not iterate cursor, %+v", err)
		return nil, err
	}


	if err := cur.Close(ctx); err!= nil {
		log.Fatalf("Could not close cursor, %+v", err)
		return nil, err
	}

	return users, nil

}


func (c *MongoDbUserStore) GetUserById(ctx context.Context, id string) (*models.User, error) {

	newid, err := primitive.ObjectIDFromHex(id) 

	if err != nil {
		log.Fatalf("Could not convert id to object id, %+v", err)
	}

	cur := c.coll.FindOne(ctx, bson.M{"_id": newid})

	if err != nil {
		log.Fatalf("Could not find user with id %s", id)
		return nil, err
	}
	
	var user models.User
	cur.Decode(&user)
	
	if err != nil {
		log.Fatalf("Could not decode user, %+v", err)
		return nil, err
	}

	return &user, nil
}
