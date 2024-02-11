package route

import (
	"github.com/labstack/echo/v4"
	"github.com/sumiredc/shorturl/app/interface/controller"
	"github.com/sumiredc/shorturl/app/interface/gateway/database/middleware"
)

func (_ *Route) Register(e *echo.Echo) {
	c := controller.New()

	// CORSの設定
	e.Use(middleware.Cors)

	// URLを定義
	e.GET("/url", c.UrlIndex)
	e.POST("/url", c.UrlStore)
	e.GET("/:slug", c.UrlRedirect)
}
