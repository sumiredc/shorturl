package link_repository

import "github.com/sumiredc/shorturl/app/domain/entity"

func (r *LinkRepository) Create(original string, slug string, length int) (*entity.Link, error) {
	link := &entity.Link{
		Original: original,
		Slug:     slug,
		Length:   length,
	}

	err := r.Database.Create(&link).Error
	if err != nil {
		return nil, err
	}

	return link, nil
}
