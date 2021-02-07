package main

import (
	"GoSnowball/handler/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"os"
)

func runHttpserver() {
	e := echo.New()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if _, ok := err.(*echo.HTTPError); !ok {
			logrus.WithError(err).Errorln("[echo.HTTPErrorHandler]", c.Request().Method, c.Request().URL)
		}
		e.DefaultHTTPErrorHandler(err, c)
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{},
		AllowMethods:     []string{echo.GET, echo.OPTIONS, echo.POST},
		AllowCredentials: true,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: os.Stdout}))
	e.GET("/health", func(c echo.Context) error {
		c.NoContent(200)
		return nil
	})

	http.NewHttpHandler(e)
	e.Start(":8080")

}
func main() {
	runHttpserver()
}
