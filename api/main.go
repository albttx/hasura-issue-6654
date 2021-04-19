package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.Any("/login", login)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

// Handler
func login(c echo.Context) error {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	req := Request{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid credentials")
	}

	// NOTE: debug
	requestDump, err := httputil.DumpRequest(c.Request(), true)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(requestDump))

	// TODO: Check user password

	return c.JSON(200, echo.Map{
		"access_token": "xxx",
	})
}
