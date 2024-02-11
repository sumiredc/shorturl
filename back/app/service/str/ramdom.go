package str

import (
	"crypto/rand"
	"errors"
)

const LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Random(digit int) (string, error) {
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("Unexpected error.")
	}

	var result string
	for _, v := range b {
		result += string(LETTERS[int(v)%len(LETTERS)])
	}
	return result, nil
}
