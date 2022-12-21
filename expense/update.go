package expense

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) UpdateExpense(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	var ex Expense
	if err := c.Bind(&ex); err != nil {
		return c.JSON(http.StatusInternalServerError, "can't bind data")
	}

	err = h.db.Update(id, ex)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ex)
}
