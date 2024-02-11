package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sumiredc/shorturl/app/interface/gateway/database"
	"github.com/sumiredc/shorturl/app/interface/presenter"
	"github.com/sumiredc/shorturl/app/interface/request/url_store_request"
	"github.com/sumiredc/shorturl/app/repository/link_repository"
	"github.com/sumiredc/shorturl/app/use_case"
)

func (_ *Controller) UrlStore(c echo.Context) error {
	// バリデーション
	req, err := url_store_request.New(c.Request())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// DBに接続
	db, err := database.New().Connect()
	if err != nil {
		return err
	}

	// slugの長さを判定
	length, err := use_case.NewSlugLen(db)
	if err != nil {
		return err
	}

	// ユニークなslugを生成
	slug, err := use_case.UniqueSlug(db, length)
	if err != nil {
		return err
	}

	// レコード作成
	rep := link_repository.New(db)

	link, err := rep.Create(req.Original, slug, length)
	if err != nil {
		return err
	}

	// レスポンス生成
	res := presenter.NewLinkResponse(link)
	return c.JSON(http.StatusOK, res)
}
