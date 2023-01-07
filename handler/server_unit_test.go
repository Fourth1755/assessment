//go:build unit
// +build unit

package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {
	// Setup
	e := echo.New()
	expenseJSON := `{
		"title": "strawberry smoothie",
		"amount": 82,
		"note":"night market promotion discount 20 bath",
		"tags":["food", "beverage"]
	}`
	req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(expenseJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	newsMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(1, "strawberry smoothie", 82, "night market promotion discount 20 bath", `["food", "beverage"]`)
	db, mock, err := sqlmock.New()
	mock.ExpectQuery("SELECT (.+) FROM expenses").WillReturnRows(newsMockRows)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := &handler{db}
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.CreatedExpenseHandler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expenseJSON, rec.Body.String())
	}
}
