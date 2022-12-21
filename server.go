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
	db := expense.InitDB()
	handler := expense.CreateHandler(db)

	e := echo.New()
	e.Use(expense.AuthMiddleware)

	expense.CreateRoute(e, handler)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	log.Printf("Server up by port %v\n", os.Getenv("PORT"))
	go func() {
		if err := e.Start(os.Getenv("PORT")); err != nil {
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
