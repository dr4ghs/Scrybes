package assets

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed dist/*
var assets embed.FS

var contentHandler = echo.WrapHandler(http.FileServer(http.FS(assets)))
var contentRewrite = middleware.Rewrite(map[string]string{"/*": "/dist/$1"})

func Mount(r *echo.Echo) {
	r.GET("/*", contentHandler, contentRewrite)
}
