package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MEDALIALPHA331/reservation/database"
	"github.com/MEDALIALPHA331/reservation/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	store database.UserStore
}

func NewUserHandler(store database.UserStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}



func (u *UserHandler) DeleteUserHandler(c echo.Context) error {

	id := c.Param("id");
	oid , err := primitive.ObjectIDFromHex(id);

	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	 ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	 defer cancel()

	u.store.DeleteUser(ctx, oid)

	return c.String(http.StatusNoContent, "Deleted")
} 


func (u *UserHandler) UpdateUserHandler(c echo.Context) error {

	id := c.Param("id");
	oid , err := primitive.ObjectIDFromHex(id);

	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}


	 ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	 defer cancel()


	var user *models.UpdateUserDTO

	err = c.Bind(&user)
	
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	update := bson.M{"firstName": user.FirstName, "lastName": user.LastName}

	u.store.UpdateUser(ctx, oid, update)

	return nil
} 


func (u *UserHandler) CreateUserHandler(c echo.Context) error {
	
	var user models.UserDTO
	// TODO: Add in validation for these fileds: min & max, password structure, email validation

	
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()

	res, err := u.store.CreateUser(ctx, models.CreateUserFromDTO(&user))

	if err != nil {
		log.Fatalln(err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (u *UserHandler) GetUserByIdHandler(c echo.Context) error {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()

	newid, err := primitive.ObjectIDFromHex(id) 

	if err != nil {
		fmt.Errorf("Could not convert id to object id, %+v", err)
		return c.String(http.StatusBadRequest, "bad request") 
	}

	user, err := u.store.GetUserById(ctx, newid)
	if err != nil {
		return c.String(http.StatusNotFound, "User Not Found") 
	}

	return c.JSON(http.StatusOK, user)
}


func (u *UserHandler) GetAllUsersHandler(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()

	users, err := u.store.GetAllUsers(ctx)
	if err != nil {
		return c.String(http.StatusNotFound, "User Not Found") 
	}

	return c.JSON(http.StatusOK, users)
}






func HelloHandler(e echo.Context) error {
	return e.JSON(200, map[string]string{
		"message": "it works! 😝",
	})
}
