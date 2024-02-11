package use_case

import (
	"example.com/shorturl/app/repository/link_repository"
	"example.com/shorturl/app/service/num"
	"example.com/shorturl/app/service/str"
	"gorm.io/gorm"
)

const (
	SLUG_LEN = 5
)

func NewSlugLen(db *gorm.DB) (int, error) {
	length := SLUG_LEN
	letters_len := len(str.LETTERS)
	rep := link_repository.New(db)

	for true {

		p, err := num.Permutation(letters_len, length)
		if err != nil {
			return 0, err
		}

		count, err := rep.LengthCount(length)
		if err != nil {
			return 0, err
		}

		if count < int64(p/2) {
			break
		}
		length++
	}

	return length, nil
}
