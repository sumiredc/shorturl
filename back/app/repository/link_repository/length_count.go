package link_repository

func (r *LinkRepository) LengthCount(length int) (int64, error) {
	var count int64
	err := r.Database.Table(r.Table).Where("length", length).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
