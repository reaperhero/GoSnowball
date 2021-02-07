package http

import (
	"github.com/labstack/echo"
)


type handler struct {
}



func NewHttpHandler(e *echo.Echo) {
	h := handler{}
	e.GET("/api/index/quote", h.index)
	e.GET("/api/index/hotStock", h.hotStock)
	e.GET("/api/index/public", h.public)
	e.GET("/api/index/news", h.news)
}
