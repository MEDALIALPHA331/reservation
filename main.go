package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MEDALIALPHA331/reservation/server"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


const (
	DBNAME="hotel_reservation"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load enviroment variables file, %+v", err)
	}

	var (
		PORT = os.Getenv("PORT")
		MONGO_URI = os.Getenv("MONGO_URL")
	)

	_, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatalf("Could not connect to mongo db, %+v", err)
	}


	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiv1 := e.Group("/apiv1")
	apiv1.GET("/", server.HelloHandler)

	

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%+v", PORT)))
}
