package server

import "github.com/labstack/echo/v4"



func HelloHandler(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"message": "it works! 😝",
	})
}