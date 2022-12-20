package main

import (
	"os"

	"github.com/Dsypasit/assessment/expense"
	"github.com/labstack/echo/v4"
)

func main() {
	expense.InitDB()
	e := echo.New()

	expense.CreateRoute(e)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
