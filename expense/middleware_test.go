//go:build unit

package expense

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	f := func(c echo.Context) error {
		return c.JSON(http.StatusOK, "test")
	}
	h := AuthMiddleware(f)

	req.Header.Set("Authorization", "November 10, 2009")

	// Assertions
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestAuthMiddlewareError(t *testing.T) {
	input := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Empty authrization should be status unauthorization",
			input: "",
			want:  http.StatusUnauthorized,
		},
		{
			name:  "invalid authrization should be status unauthorization",
			input: "test na kub",
			want:  http.StatusUnauthorized,
		},
	}

	for _, tt := range input {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			f := func(c echo.Context) error {
				return c.JSON(http.StatusOK, "test")
			}
			h := AuthMiddleware(f)

			// header input
			req.Header.Set("Authorization", tt.input)

			// Assertions
			if assert.NoError(t, h(c)) {
				assert.Equal(t, tt.want, rec.Code)
			}

		})
	}

}
