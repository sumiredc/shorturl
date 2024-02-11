package link_repository

import "gorm.io/gorm"

type LinkRepository struct {
	Table    string
	Database *gorm.DB
}

func New(db *gorm.DB) *LinkRepository {
	return &LinkRepository{Table: "links", Database: db}
}
