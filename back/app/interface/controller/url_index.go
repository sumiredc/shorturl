package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (_ *Controller) UrlIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
