package pages

import "github.com/labstack/echo/v4"

func Mount(r *echo.Echo) {
	r.GET("/home", func(c echo.Context) error {
		return Home().Render(c.Request().Context(), c.Response().Writer)
	})
}
