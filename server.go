package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Dsypasit/assessment/expense"
	"github.com/labstack/echo/v4"
)

func main() {
	expense.InitDB()
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Values("Authorization")[0]
			if auth != "November 10, 2009" {
				return c.JSON(http.StatusUnauthorized, "Unauthorized")
			}
			fmt.Println("from middle ware")
			return next(c)
		}
	})

	expense.CreateRoute(e)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
