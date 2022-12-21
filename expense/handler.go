package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	db DB
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if len(c.Request().Header.Values("Authorization")) == 0 {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		auth := c.Request().Header.Values("Authorization")[0]
		if auth != "November 10, 2009" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		return next(c)
	}
}

func CreateHandler(db DB) Handler {
	return Handler{
		db: db,
	}
}

func CreateRoute(app *echo.Echo, handler Handler) {
	app.POST("/expenses", handler.AddExpense)
	app.GET("/expenses/:id", handler.GetExpenseByID)
	app.PUT("/expenses/:id", handler.UpdateExpense)
	app.GET("/expenses", handler.GetExpenses)
}
