package http

import (
	"github.com/labstack/echo"
)

type handler struct {
}

func NewHttpHandler(e *echo.Echo) {
	h := handler{}
	apigroup := e.Group("/api")
	{
		apigroup.GET("/index/quote", h.index)
		apigroup.GET("/index/hotStock", h.hotStock)
		apigroup.GET("/index/news", h.news)
		apigroup.GET("/choose/industries", h.industries)
		apigroup.GET("/choose/areas", h.chooseAreas)
		apigroup.GET("/choose/tools", h.chooseTools)
	}
}
