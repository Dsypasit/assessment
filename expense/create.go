package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func AddExpense(c echo.Context) error {
	var ex Expense
	if c.Request().Header.Get("Content-type") != "application/json" {
		return c.JSON(http.StatusBadRequest, "require json type")
	}
	if err := c.Bind(&ex); err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't bind data")
	}

	st, err := db.Prepare("INSERT INTO expenses (title, amount, note, tags) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't prepare statement")
	}

	row := st.QueryRow(&ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))

	if err := row.Scan(&ex.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, ex)
}

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
