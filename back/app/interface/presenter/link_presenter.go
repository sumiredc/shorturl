package presenter

import (
	"fmt"
	"os"

	"github.com/sumiredc/shorturl/app/domain/entity"
)

type LinkResponse struct {
	Original string `json:"original"`
	Short    string `json:"short"`
}

func NewLinkResponse(link *entity.Link) *LinkResponse {
	url := os.Getenv("FRONT_URL")
	short := fmt.Sprintf("%s/%s", url, link.Slug)

	return &LinkResponse{
		Original: link.Original,
		Short:    short,
	}
}

func NewLinksResponse(links []*entity.Link) []*LinkResponse {
	var res []*LinkResponse
	url := os.Getenv("FRONT_URL")

	for _, link := range links {
		short := fmt.Sprintf("%s/%s", url, link.Slug)
		res = append(res, &LinkResponse{Original: link.Original, Short: short})
	}

	return res
}
