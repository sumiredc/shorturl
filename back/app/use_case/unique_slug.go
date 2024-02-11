package use_case

import (
	"errors"

	"example.com/shorturl/app/repository/link_repository"
	"example.com/shorturl/app/service/str"
	"gorm.io/gorm"
)

const (
	MAX_ATTEMPTS = 10
)

func UniqueSlug(db *gorm.DB, length int) (string, error) {
	var slug string
	var err error

	attempts := 0
	rep := link_repository.New(db)

	for true {
		if attempts > MAX_ATTEMPTS {
			return "", errors.New("Attempt limit reached.")
		}

		slug, err = str.Random(length)
		if err != nil {
			return "", err
		}

		exists, err := rep.Exists(slug)
		if err != nil {
			return "", err
		}

		if !exists {
			break
		}

		attempts++
	}

	return slug, nil
}
