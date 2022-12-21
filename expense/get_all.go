package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetExpenses(c echo.Context) error {
	expenses, err := h.db.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, expenses)
}
