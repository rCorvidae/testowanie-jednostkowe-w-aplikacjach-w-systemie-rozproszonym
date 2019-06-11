package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	msg1       = "OK"
	validError = "{\n \"msg\": \" Server 3 is down - check connection\"\n}\n"
)

func TestCheckUserExist(t *testing.T) {

	assert.Equal(t, true, checkUserExist("12345678912", "Kuba"))
	assert.Equal(t, false, checkUserExist("1234567", "Marek"))

}

func TestGetStatus(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, getStatus(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, msg1, rec.Body.String())
	}
}

func TestCreateNewCard(t *testing.T) {

	testCard, err := createNewCard()

	assert.Empty(t, &testCard)
	assert.NotNil(t, err)
}

func TestCreateClient(t *testing.T) {
	// Setup
	e := echo.New()
	f := make(url.Values)
	f.Set("name", "Test")
	f.Set("pesel", "221819")
	req := httptest.NewRequest(http.MethodPost, "/createClient", strings.NewReader(f.Encode()))

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, createClient(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, validError, rec.Body.String())
	}
}
