package main

import (
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/tonnac/learngo/scrapper"
)

// FILENAME saved file name
const FILENAME string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(FILENAME)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(FILENAME, FILENAME)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
