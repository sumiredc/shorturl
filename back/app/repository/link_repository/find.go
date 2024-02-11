package link_repository

import "github.com/sumiredc/shorturl/app/domain/entity"

func (r *LinkRepository) Find(slug string) (*entity.Link, error) {
	link := &entity.Link{}
	err := r.Database.Where("slug", slug).First(&link).Error
	if err != nil {
		return nil, err
	}

	return link, nil
}
