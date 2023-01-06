package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatedExpense(t *testing.T) {
	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 82,
		"note": "night market promotion discount 20 bath", 
		"tags": ["food", "beverage"]
	}`)
	var e Expense

	res := request(http.MethodPost, uri("expenses"), body)
	err := res.Decode(&e)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, e.ID)
	assert.Equal(t, "strawberry smoothie", e.Title)
	assert.Equal(t, 82, e.Amount)
	assert.Equal(t, "night market promotion discount 20 bath", e.Note)
	assert.Equal(t, []string([]string{"food", "beverage"}), e.Tags)
}
func TestGetExpenseByID(t *testing.T) {
	c := seedExpense(t)

	var latest Expense
	res := request(http.MethodGet, uri("expenses", strconv.Itoa(c.ID)), nil)
	err := res.Decode(&latest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, c.ID, latest.ID)
	assert.NotEmpty(t, latest.Title)
	assert.NotEmpty(t, latest.Amount)
	assert.NotEmpty(t, latest.Note)
	assert.NotEmpty(t, latest.Tags)
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}
func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}
func seedExpense(t *testing.T) Expense {
	var e Expense
	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 82,
		"note": "night market promotion discount 20 bath", 
		"tags": ["food", "beverage"]
	}`)
	err := request(http.MethodPost, uri("expenses"), body).Decode(&e)
	if err != nil {
		t.Fatal("can't create uomer:", err)
	}
	return e
}
