package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//Card it's fresh card metadata
type Card struct {
	PIN        string `json:"pin"`
	CardNumber string `json:"cardNumber"`
}

func main() {
	e := echo.New()
	e.GET("/createCard", createCard)
	e.Logger.Fatal(e.Start(":5002"))
}

// e.GET("/createCard", createCard)
func createCard(c echo.Context) error {
	cardNumber := strconv.Itoa(rand.Intn(2000000) + 1000000)
	pin := strconv.Itoa(rand.Intn(8999) + 1000)
	return c.JSONPretty(http.StatusOK, &Card{PIN: pin, CardNumber: cardNumber}, " ")
}
