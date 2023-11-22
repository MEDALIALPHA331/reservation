package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load enviroment variables file, %+v", err)
	}

	var (
		PORT = os.Getenv("PORT")
	)


	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "it works! 😝",
		})
	})


	e.Logger.Fatal(e.Start(fmt.Sprintf(":%+v", PORT)))

}
