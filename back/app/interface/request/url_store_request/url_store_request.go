package url_store_request

import (
	"net/http"

	"example.com/shorturl/app/service/validator"
)

type UrlStoreRequest struct {
	Request  *http.Request
	Original string
}

func New(req *http.Request) (*UrlStoreRequest, error) {
	vali := validator.New()

	original, err := validateOriginal(req, vali)
	if err != nil {
		return nil, err
	}

	return &UrlStoreRequest{
		Request:  req,
		Original: original,
	}, nil
}

func validateOriginal(req *http.Request, vali *validator.Validator) (string, error) {
	value := req.FormValue("original")

	if err := vali.Required(value); err != nil {
		return "", err
	}

	if err := vali.StringMax(value, 2000); err != nil {
		return "", err
	}

	if err := vali.Url(value); err != nil {
		return "", err
	}

	return value, nil
}
