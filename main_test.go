package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetQuotes(t *testing.T) {
	router := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/data/quotes", nil)
	rec := httptest.NewRecorder()
	ctx := router.NewContext(req, rec)

	err := getQuotes(ctx)
	if err != nil {
		t.Error(err)
	}
	if rec.Code != http.StatusOK {
		t.Error("Not OK")
	}

	got := rec.Body.String()
	exp := "{\"name\":\"allochi\",\"age\":45}\n"
	if got != exp {
		assert.Equal(t, exp, got)
	}
}
