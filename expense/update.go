package expense

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func UpdateExpense(c echo.Context) error {
	id := c.Param("id")

	var ex Expense
	if err := c.Bind(&ex); err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't bind data")
	}

	st, err := db.Prepare("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't prepare statement")
	}

	result, err := st.Exec(id, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't update information")
	}

	if rowAffected, err := result.RowsAffected(); err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't get row affect")
	} else if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	ex.ID, _ = strconv.Atoi(id)

	return c.JSON(http.StatusOK, ex)
}
