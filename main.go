package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MEDALIALPHA331/reservation/database"
	"github.com/MEDALIALPHA331/reservation/server"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load enviroment variables file, %+v", err)
	}

	var (
		PORT      = os.Getenv("PORT")
		MONGO_URI = os.Getenv("MONGO_URI")
	)

	mongodb, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatalf("Could not connect to mongo db, %+v", err)
	}

	handler := server.NewUserHandler(database.NewMongoUserStore(mongodb))

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiv1 := e.Group("/apiv1")
	apiv1.GET("/", server.HelloHandler)
	apiv1.GET("/users", handler.GetAllUsersHandler)
	apiv1.GET("/users/:id", handler.GetUserByIdHandler)
	apiv1.POST("/users", handler.CreateUserHandler)
	apiv1.PUT("/users/:id", handler.UpdateUserHandler)
	apiv1.DELETE("/users/:id", handler.DeleteUserHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%+v", PORT)))
}
