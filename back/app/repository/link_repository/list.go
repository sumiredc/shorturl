package link_repository

import "github.com/sumiredc/shorturl/app/domain/entity"

func (r *LinkRepository) List() ([]*entity.Link, error) {
	var links []*entity.Link

	err := r.Database.Where("deleted_at IS NULL").Find(&links).Error
	if err != nil {
		return nil, err
	}

	return links, nil
}
