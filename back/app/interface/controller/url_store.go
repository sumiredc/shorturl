package controller

import (
	"errors"
	"net/http"

	"example.com/shorturl/app/domain/entity"
	"example.com/shorturl/app/interface/gateway/database"
	"example.com/shorturl/app/interface/presenter"
	"example.com/shorturl/app/service/num"
	"example.com/shorturl/app/service/str"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	MAX_ATTEMPTS = 10
	SLUG_LEN     = 5
)

func (_ *Controller) UrlStore(c echo.Context) error {
	db, err := database.New().Connect()
	if err != nil {
		return err
	}

	// slugの長さを判定

	length := SLUG_LEN
	letters_len := len(str.LETTERS)

	for true {
		p, err := num.Permutation(letters_len, length)
		if err != nil {
			return err
		}

		var count int64
		err = db.Table("links").Where("length", length).Count(&count).Error

		if err != nil {
			return err
		}

		if count < int64(p/2) {
			break
		}
		length++
	}

	// ユニークなslugを生成

	var slug string
	attempts := 0

	for true {
		if attempts > MAX_ATTEMPTS {
			return errors.New("Attempt limit reached.")
		}

		slug, err = str.Random(5)
		if err != nil {
			return err
		}

		link := &entity.Link{}
		err = db.Where("slug = ?", slug).First(link).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			break
		}

		if err != nil {
			return err
		}

		attempts++
	}

	// レコード作成

	original := c.Request().FormValue("original")

	link := &entity.Link{
		Original: original,
		Slug:     slug,
		Length:   length,
	}

	err = db.Create(&link).Error
	if err != nil {
		return err
	}

	// レスポンス生成

	res := presenter.NewLinkResponse(link)

	return c.JSON(http.StatusOK, res)
}
