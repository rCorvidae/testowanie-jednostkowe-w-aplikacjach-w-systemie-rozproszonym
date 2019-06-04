package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/users/:id", getUser)
	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	id := c.Param("id")
	if id == "4" || id == "5" {
		return c.String(http.StatusOK, "No such user\n")
	} else {
		return c.String(http.StatusOK, fmt.Sprintf("User %s exist\n", id))
	}

}
