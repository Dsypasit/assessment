package expense

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func GetExpenseByID(c echo.Context) error {
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}
	st, err := db.Prepare("SELECT * FROM expenses WHERE id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "can't prepare db statement")
	}

	var expense Expense
	err = st.QueryRow(id).Scan(&expense.ID, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, expense)
}
