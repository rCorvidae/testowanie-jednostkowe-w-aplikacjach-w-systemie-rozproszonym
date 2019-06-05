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
	validError = "{\n \"msg\": \"Validation Error\"\n}\n"
)

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

func TestBasicValidate(t *testing.T) {

	assert.Equal(t, false, basicValidate("", "b", "c", "12345678912"))
	assert.Equal(t, false, basicValidate("a", "", "c", "12345678912"))
	assert.Equal(t, false, basicValidate("a", "b", "", "12345678912"))
	assert.Equal(t, false, basicValidate("a", "b", "c", "12345678"))
	assert.Equal(t, true, basicValidate("a", "b", "c", "12345678912"))
}

func TestSendDataToApp2(t *testing.T) {

	testAccount, err := sendDataToApp2("a", "b", "c", "12345678912")

	assert.Empty(t, &testAccount)
	assert.NotNil(t, err)
}

func TestCreateNewCard(t *testing.T) {

	e := echo.New()
	f := make(url.Values)
	f.Set("name", "Kuba")
	f.Set("surname", "Test")
	f.Set("email", "test@test.com")
	f.Set("pesel", "22115121819")
	req := httptest.NewRequest(http.MethodPost, "/createNewCard", strings.NewReader(f.Encode()))

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, createNewCard(c)) {
		assert.Equal(t, http.StatusNotAcceptable, rec.Code)
		assert.Equal(t, validError, rec.Body.String())
	}
}
