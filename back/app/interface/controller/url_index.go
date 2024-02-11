package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sumiredc/shorturl/app/interface/gateway/database"
	"github.com/sumiredc/shorturl/app/interface/presenter"
	"github.com/sumiredc/shorturl/app/repository/link_repository"
)

func (_ *Controller) UrlIndex(c echo.Context) error {
	// DBに接続
	db, err := database.New().Connect()
	if err != nil {
		return err
	}

	// 一覧を取得
	rep := link_repository.New(db)
	links, err := rep.List()
	if err != nil {
		return err
	}

	// レスポンス生成
	res := presenter.NewLinksResponse(links)

	return c.JSON(http.StatusOK, res)
}
