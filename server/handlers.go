package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/MEDALIALPHA331/reservation/database"
	"github.com/MEDALIALPHA331/reservation/models"
	"github.com/labstack/echo/v4"
)



type UserHandler struct {
	store database.UserStore
}


func NewUserHandler(store database.UserStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}



func (u *UserHandler) CreateUserHandler(c echo.Context) error {

	var user models.UserDTO


	err := c.Bind(&user); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 600 * time.Millisecond)
	defer cancel()

	res, err :=  u.store.CreateUser(ctx, models.CreateUserFromDTO(&user))

	if err != nil {
		log.Fatalln(err)
	}

	return c.JSON(http.StatusCreated, res)
}



func HelloHandler(e echo.Context) error {
	return e.JSON(200, map[string]string{
		"message": "it works! 😝",
	})
}