package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) AddExpense(c echo.Context) error {
	var ex Expense
	if c.Request().Header.Get("Content-type") != "application/json" {
		return c.JSON(http.StatusBadRequest, "require json type")
	}
	if err := c.Bind(&ex); err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't bind data")
	}

	err = h.db.Create(&ex)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, ex)
}
