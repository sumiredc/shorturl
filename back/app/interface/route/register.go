package route

import (
	"os"

	"example.com/shorturl/app/interface/controller"
	"github.com/labstack/echo/v4"
)

func (_ *Route) Register(e *echo.Echo) {
	c := controller.New()

	// CORSの設定
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			frontUrl := os.Getenv("FRONT_URL")
			c.Response().Header().Set("Access-Control-Allow-Origin", frontUrl)
			return next(c)
		}
	})

	e.GET("/url", c.UrlIndex)
	e.POST("/url", c.UrlStore)
}
