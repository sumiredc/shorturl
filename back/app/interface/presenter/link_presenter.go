package presenter

import (
	"fmt"
	"os"

	"example.com/shorturl/app/domain/entity"
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
