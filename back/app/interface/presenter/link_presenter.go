package presenter

import (
	"github.com/sumiredc/shorturl/app/domain"
	"github.com/sumiredc/shorturl/app/domain/entity"
)

type LinkResponse struct {
	Original string `json:"original"`
	Short    string `json:"short"`
}

func NewLinkResponse(link *entity.Link) *LinkResponse {
	short := domain.ShortUrl(link.Slug)

	return &LinkResponse{
		Original: link.Original,
		Short:    short,
	}
}

func NewLinksResponse(links []*entity.Link) []*LinkResponse {
	res := []*LinkResponse{}

	for _, link := range links {
		short := domain.ShortUrl(link.Slug)
		res = append(res, &LinkResponse{Original: link.Original, Short: short})
	}

	return res
}
