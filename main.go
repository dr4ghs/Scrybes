package main

import (
	"github.com/dr4ghs/scrybes/assets"
	"github.com/dr4ghs/scrybes/frontend/templates"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	assets.Mount(e)
	templates.Mount(e)

	e.Logger.Fatal(e.Start(":8080"))
}
