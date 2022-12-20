package expense

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetExpenseByID(c echo.Context) error {
	id := c.Param("id")
	log.Println(id)
	st, err := db.Prepare("SELECT * FROM expenses WHERE id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "can't prepare db statement")
	}

	var expense Expense
	err = st.QueryRow(id).Scan(&expense.ID, &expense.Title, &expense.Amount, &expense.Note, &expense.Tags)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, expense)
}
