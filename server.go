package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Dsypasit/assessment/expense"
	"github.com/labstack/echo/v4"
)

func main() {
	db := expense.InitDBTemp()
	handler := expense.CreateHandler(db)

	e := echo.New()
	e.Use(expense.AuthMiddleware)

	expense.CreateRoute(e, handler)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	log.Println("Server up by port 2565")
	go func() {
		if err := e.Start(":2565"); err != nil {
			e.Logger.Fatal("shutdown server")
		}
	}()

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
