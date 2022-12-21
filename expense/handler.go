package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

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

func CreateRoute(app *echo.Echo) {
	app.POST("/expenses", AddExpense)
	app.GET("/expenses/:id", GetExpenseByID)
	app.PUT("/expenses/:id", UpdateExpense)
	app.GET("/expenses", GetExpenses)
}
