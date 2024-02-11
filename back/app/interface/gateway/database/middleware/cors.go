package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
)

func Cors(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		frontUrl := os.Getenv("FRONT_URL")
		c.Response().Header().Set("Access-Control-Allow-Origin", frontUrl)
		return next(c)
	}
}
