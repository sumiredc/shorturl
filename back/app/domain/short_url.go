package domain

import (
	"fmt"
	"os"
)

func ShortUrl(slug string) string {
	url := os.Getenv("APP_URL")
	return fmt.Sprintf("%s/%s", url, slug)
}
