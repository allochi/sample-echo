package main

import (
	. "github.com/allochi/sample-echo/handlers"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetQuotes(t *testing.T) {
	router := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/data/quotes", nil)
	rec := httptest.NewRecorder()
	ctx := router.NewContext(req, rec)

	err := GetQuotes(ctx)
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

func TestGetTenders(t *testing.T) {
	router := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/data/tenders", nil)
	rec := httptest.NewRecorder()
	ctx := router.NewContext(req, rec)

	err := GetTenders(ctx)
	if err != nil {
		t.Error(err)
	}
	if rec.Code != http.StatusOK {
		t.Error("Not OK")
	}

	got := rec.Body.String()
	exp := "These are tenders!"
	if got != exp {
		assert.Equal(t, exp, got)
	}
}
