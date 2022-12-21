//go:build unit

package expense

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func (m MockDB) Create(ex *Expense) error {
	ex.ID = 1
	return nil
}

func TestAddExpense(t *testing.T) {
	db := initMockDB()
	h := CreateHandler(db)
	input := Expense{
		Amount: 89,
		Note:   "no discount",
		Tags:   []string{"beverage"},
		Title:  "apple smoothie",
	}
	inputJson, _ := json.Marshal(input)

	expected := Expense{
		ID:     1,
		Amount: 89,
		Note:   "no discount",
		Tags:   []string{"beverage"},
		Title:  "apple smoothie",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(inputJson))
	req.Header.Add("Content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.AddExpense(c)) {
		var result Expense
		json.NewDecoder(rec.Body).Decode(&result)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expected, result)
	}
}
