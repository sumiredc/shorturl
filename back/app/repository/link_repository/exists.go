package link_repository

import (
	"errors"

	"example.com/shorturl/app/domain/entity"
	"gorm.io/gorm"
)

func (r *LinkRepository) Exists(slug string) (bool, error) {
	link := &entity.Link{}
	err := r.Database.Where("slug = ?", slug).First(link).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err == nil {
		return true, nil
	}

	return false, err
}
