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
	GetUserById(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, id primitive.ObjectID, update bson.M) error
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
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

func (c *MongoDbUserStore) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	_ , err := c.coll.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		log.Fatalf("User deletion failed, %+v", err)
		return err
	}

	return nil
}

func (c *MongoDbUserStore) UpdateUser(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	_ , err := c.coll.UpdateByID(ctx, id, bson.D{{"$set", update}})

	if err != nil {
		log.Fatalf("User update failed, %+v", err)
		return err
	}

	return nil
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
			// log.Fatalf("Could not decode user, %+v", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := cur.Err(); err!= nil {
		// log.Fatalf("Could not iterate cursor, %+v", err)
		return nil, err
	}


	if err := cur.Close(ctx); err!= nil {
		// log.Fatalf("Could not close cursor, %+v", err)
		return nil, err
	}

	return users, nil

}


func (c *MongoDbUserStore) GetUserById(ctx context.Context, id primitive.ObjectID) (*models.User, error) {

	cur := c.coll.FindOne(ctx, bson.M{"_id": id})
	
	var user models.User
	err := cur.Decode(&user)
	
	if err != nil {
		// log.Fatalf("Could not decode user, %+v", err)
		return nil, err
	}

	return &user, nil
}



