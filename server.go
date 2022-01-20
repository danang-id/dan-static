package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	if err := c.NoContent(code); err != nil {
		c.Logger().Error(err)
	}
}

func main() {
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.IPExtractor = echo.ExtractIPFromXFFHeader()

	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Browse: false,
		HTML5:  false,
		Root:   "content",
	}))

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "5000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
