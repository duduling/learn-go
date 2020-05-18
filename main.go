package main

import (
	"fmt"
	"strings"

	"github.com/duduling/learngo/scrapper"
	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.File("index.html")
}

func handleScrap(c echo.Context) error {
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	fmt.Println(term)
	return nil
}

func main() {
	// scrapper.Scrape("term")
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrap)

	e.Logger.Fatal(e.Start(":1234"))
}
