package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//Msg is basic string structure
type Msg struct {
	Message string `json:"msg"`
}

//Card it's fresh card metadata
type Card struct {
	PIN        string `json:"pin"`
	CardNumber string `json:"cardNumber"`
}

//Account ...
type Account struct {
	UID      string `json:"uid"`
	CardData Card
}

var users = map[string]string{
	"12345678912": "Kuba",
	"12345678913": "Adam",
}

func main() {
	e := echo.New()
	e.POST("/createClient", createClient)
	e.GET("/status", getStatus)
	e.Logger.Fatal(e.Start(":5001"))
}

//e.POST("/createClient", createClient)
func createClient(c echo.Context) error {

	name := c.FormValue("name")
	pesel := c.FormValue("pesel")
	if !checkUserExist(pesel, name) {
		users[pesel] = name
		uid := strconv.Itoa(rand.Intn(89999) + 1000)
		cardData, err := createNewCard()
		if err != nil {
			return c.JSONPretty(http.StatusOK, &Msg{Message: " Server 3 is down - check connection"}, " ")
		}

		return c.JSONPretty(http.StatusOK, &Account{
			UID:      uid,
			CardData: *cardData}, " ")

	}

	return c.JSONPretty(http.StatusNotAcceptable, &Msg{Message: "User exist"}, " ")
}

func checkUserExist(u, e string) bool {
	if _, ok := users[u]; !ok {
		return false
	}
	return true
}

func createNewCard() (*Card, error) {

	var newCard Card
	resp, err := http.Get("http://localhost:5002/createCard")
	if err != nil {
		return &Card{}, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&newCard)
	if err != nil {
		return &Card{}, err
	}
	return &newCard, nil

}

// e.GET("/status", getStatus)
func getStatus(c echo.Context) error {
	return c.String(http.StatusOK, "OK")

}
