package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	msg1       = "OK"
	validError = "{\n \"msg\": \" Server 3 is down - check connection\"\n}\n"
)

func TestCreateCard(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/createCard", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, createCard(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "pin")
	}
}
