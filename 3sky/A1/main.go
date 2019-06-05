package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
)

/**
curl -F "name=Kuba" \
-F "surname=W" \
-F "email=test@test.com" \
-F "pesel=121113221819" localhost:5000/createNewCard
**/

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

//Client is created client with Card Number, UID and default PIN
type Client struct {
	Name    string `json:"name"`
	Surname string `json:"surename"`
	Email   string `json:"maile"`
	PESEL   string `json:"pesel"`
	Result  Account
}

func main() {
	e := echo.New()
	e.POST("/createNewCard", createNewCard)
	e.GET("/status", getStatus)
	e.Logger.Fatal(e.Start(":5000"))
}

// e.POST("/createNewCard", createNewCard)
func createNewCard(c echo.Context) error {
	name := c.FormValue("name")
	surname := c.FormValue("surname")
	email := c.FormValue("email")
	pesel := c.FormValue("pesel")
	if basicValidate(name, surname, email, pesel) {
		newAccount, err := sendDataToApp2(name, surname, email, pesel)
		if err != nil {
			return c.JSONPretty(http.StatusOK, &Msg{Message: " Server 2 is down - check connection"}, " ")
		}
		return c.JSONPretty(http.StatusOK, &Client{
			Name:    name,
			Surname: surname,
			Email:   email,
			PESEL:   pesel,
			Result:  *newAccount}, " ")

	}
	return c.JSONPretty(http.StatusNotAcceptable, &Msg{Message: "Validation Error"}, " ")
}

// e.GET("/users/:id", getStatus)
func getStatus(c echo.Context) error {
	return c.String(http.StatusOK, "OK")

}

func sendDataToApp2(n, sn, e, i string) (*Account, error) {

	var newAccount Account
	resp, err := http.PostForm("http://localhost:5001/createClient", url.Values{"name": {n}, "surename": {sn}, "email": {e}, "pesel": {i}})

	if err != nil {
		return &Account{}, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&newAccount)
	if err != nil {
		return &Account{}, err
	}
	return &newAccount, nil

}

func basicValidate(n, sn, e, i string) bool {
	if n == "" {
		return false
	}
	if sn == "" {
		return false
	}
	if e == "" {
		return false
	}
	if len(i) < 10 {
		return false
	}
	return true
}
