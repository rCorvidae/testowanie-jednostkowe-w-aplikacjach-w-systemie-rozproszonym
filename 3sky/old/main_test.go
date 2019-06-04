package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	msg1 = "User 1 exist\n"
	msg2 = "User 2 exist\n"
	msg3 = "User 3 exist\n"
	msg4 = "No such user\n"
	msg5 = "No such user\n"
)

func TestGetUser1(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, msg1, rec.Body.String())
	}
}
func TestGetUser2(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	// Assertions
	if assert.NoError(t, getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, msg2, rec.Body.String())
	}
}
func TestGetUser3(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("3")

	// Assertions
	if assert.NoError(t, getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, msg3, rec.Body.String())
	}
}

func TestGetUser4(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("4")

	// Assertions
	if assert.NoError(t, getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, msg4, rec.Body.String())
	}
}

func TestGetUser5(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("5")

	// Assertions
	if assert.NoError(t, getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, msg5, rec.Body.String())
	}
}
