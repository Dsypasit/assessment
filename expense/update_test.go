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

func (m MockDB) Update(id int, ex Expense) error {
	return nil
}

func TestUpdateExpense(t *testing.T) {
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
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(inputJson))
	req.Header.Add("Content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/expenses/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, h.UpdateExpense(c)) {
		var result Expense
		json.NewDecoder(rec.Body).Decode(&result)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, result)
	}
}
