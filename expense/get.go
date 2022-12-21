package expense

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetExpenseByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}
	expense, err := h.db.GetByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, expense)
}
