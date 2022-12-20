package expense

import "github.com/labstack/echo/v4"

type Expense struct {
	ID     int      `json:"id"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
	Title  string   `json:"title"`
}

func CreateRoute(app *echo.Echo) {
	app.POST("/expenses", AddExpense)
	app.GET("/expenses/:id", GetExpenseByID)
}
