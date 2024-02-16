package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	msg := sayHello("Becky")
	fmt.Println(msg)
	echoTheGo()
}

// add sayHello function for greeting
func sayHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

// simple server. Echo is a performance-focused, extensible, open-source Go web application framework
func echoTheGo() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
