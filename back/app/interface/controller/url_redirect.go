package controller

import (
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sumiredc/shorturl/app/interface/gateway/database"
	"github.com/sumiredc/shorturl/app/repository/link_repository"
	"gorm.io/gorm"
)

func (_ *Controller) UrlRedirect(c echo.Context) error {
	slug := c.Param("slug")

	// DBに接続
	db, err := database.New().Connect()
	if err != nil {
		return err
	}

	// レコード取得
	rep := link_repository.New(db)
	link, err := rep.Find(slug)

	// レコードが存在しなければ、Not Foundページへリダイレクト
	if errors.Is(err, gorm.ErrRecordNotFound) {
		url := os.Getenv("FRONT_NOT_FOUND_URL")
		return c.Redirect(http.StatusMovedPermanently, url)
	}

	if err != nil {
		return err
	}

	// リダイレクト
	return c.Redirect(http.StatusMovedPermanently, link.Original)
}
