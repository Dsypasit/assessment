package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func GetExpenses(c echo.Context) error {
	st, err := db.Prepare("SELECT * FROM expenses")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "can't prepare db statement")
	}

	rows, err := st.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "can't query information")
	}

	var expenses []Expense
	for rows.Next() {
		var ex Expense
		rows.Scan(&ex.ID, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Note))
		expenses = append(expenses, ex)
	}

	return c.JSON(http.StatusOK, expenses)
}
