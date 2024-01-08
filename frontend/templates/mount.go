package templates

import (
	"github.com/dr4ghs/scrybes/frontend/templates/pages"
	"github.com/dr4ghs/scrybes/frontend/templates/shared"
	"github.com/labstack/echo/v4"
)

func Mount(r *echo.Echo) {
	pages.Mount(r)
	shared.Mount(r)
}
