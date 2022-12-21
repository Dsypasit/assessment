//go:build unit

package expense

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockDB struct {
	DB
}

func (m MockDB) GetAll() ([]Expense, error) {
	return []Expense{
		{
			ID:     1,
			Amount: 89,
			Note:   "no discount",
			Tags:   []string{"beverage"},
			Title:  "apple smoothie",
		},
	}, nil
}

func initMockDB() DB {
	return MockDB{}
}

func TestGetExpenses(t *testing.T) {
	db := initMockDB()
	h := CreateHandler(db)
	expected := []Expense{
		{
			ID:     1,
			Amount: 89,
			Note:   "no discount",
			Tags:   []string{"beverage"},
			Title:  "apple smoothie",
		},
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.GetExpenses(c)) {
		var result []Expense
		json.NewDecoder(rec.Body).Decode(&result)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.ElementsMatch(t, expected, result)
	}

}
