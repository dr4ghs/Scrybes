package shared

import "github.com/labstack/echo/v4"

func Mount(r *echo.Echo) {
	r.GET("/", func(c echo.Context) error {
		return Page("Scrybes").Render(c.Request().Context(), c.Response().Writer)
	})
}
