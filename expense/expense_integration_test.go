// go:build integration
package expense

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type Response struct {
	*http.Response
	err error
}

func (r Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}
	return json.NewDecoder(r.Body).Decode(v)
}

func url(path ...string) string {
	host := "http://localhost:5000"
	if path == nil {
		return host
	}
	url := append([]string{host}, path...)
	return strings.Join(url, "/")
}

func request(method, url string, body io.Reader) *Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal("requeat error: ", err)
	}

	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}

	res, err := client.Do(req)
	return &Response{res, err}
}

func TestAddExpense(t *testing.T) {
	body := bytes.NewBufferString(`
{
"title": "strawberry smoothie",
"amount": 79,
"note": "night market promotion discount 10 bath",
"tags": ["food", "beverage"]
}
`)

	eh := echo.New()
	InitDB()
	go func(e *echo.Echo) {
		CreateRoute(e)
		e.Start(":5000")
	}(eh)

	for {
		conn, err := net.DialTimeout("tcp", "localhost:5000", 30*time.Second)
		if err != nil {
			log.Println(err)
		}

		if conn != nil {
			conn.Close()
			break
		}
	}

	var result Expense
	res := request(http.MethodPost, url("expenses"), body)
	err := res.Decode(&result)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, result.ID)
	assert.Equal(t, "strawberry smoothie", result.Title)
	assert.Equal(t, 79.0, result.Amount)
	assert.Equal(t, "night market promotion discount 10 bath", result.Note)
	assert.ElementsMatch(t, []string{"food", "beverage"}, result.Tags)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)

}

func TestGetExpenseByID(t *testing.T) {
	body := bytes.NewBufferString(`
	{
	"title": "strawberry smoothie 2",
	"amount": 7,
	"note": "night market promotion discount 100 bath", 
	"tags": ["food"]
}
	`)

	eh := echo.New()
	InitDB()
	go func(e *echo.Echo) {
		CreateRoute(e)
		e.Start(":5000")
	}(eh)

	for {
		conn, err := net.DialTimeout("tcp", "localhost:5000", 30*time.Second)
		if err != nil {
			log.Println(err)
		}

		if conn != nil {
			conn.Close()
			break
		}
	}

	var expected Expense
	res := request(http.MethodPost, url("expenses"), body)
	err := res.Decode(&expected)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var result Expense
	res = request(http.MethodGet, url("expenses", strconv.Itoa(expected.ID)), body)
	err = res.Decode(&result)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	assert.Equal(t, expected.ID, result.ID)
	assert.Equal(t, expected.Title, result.Title)
	assert.Equal(t, expected.Amount, result.Amount)
	assert.Equal(t, expected.Note, result.Note)
	assert.ElementsMatch(t, expected.Tags, result.Tags)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)

}

func TestUpdateExpense(t *testing.T) {
	body := bytes.NewBufferString(`
	{
	"title": "strawberry smoothie 2",
	"amount": 7,
	"note": "night market promotion discount 100 bath", 
	"tags": ["food"]
}
	`)

	updateBody := bytes.NewBufferString(`
{
    "title": "apple smoothie",
    "amount": 89,
    "note": "no discount",
    "tags": ["beverage"]
}
	`)

	eh := echo.New()
	InitDB()
	go func(e *echo.Echo) {
		CreateRoute(e)
		e.Start(":5000")
	}(eh)

	for {
		conn, err := net.DialTimeout("tcp", "localhost:5000", 30*time.Second)
		if err != nil {
			log.Println(err)
		}

		if conn != nil {
			conn.Close()
			break
		}
	}

	var expected Expense
	res := request(http.MethodPost, url("expenses"), body)
	err := res.Decode(&expected)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var result Expense
	res = request(http.MethodPut, url("expenses", strconv.Itoa(expected.ID)), updateBody)
	err = res.Decode(&result)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	assert.Equal(t, expected.ID, result.ID)
	assert.Equal(t, "apple smoothie", result.Title)
	assert.Equal(t, 89.0, result.Amount)
	assert.Equal(t, "no discount", result.Note)
	assert.ElementsMatch(t, []string{"beverage"}, result.Tags)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)

}
