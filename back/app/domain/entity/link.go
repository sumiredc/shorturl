package entity

import "time"

type Link struct {
	Id        int       `json:"id"`
	Original  string    `json:"original"`
	Slug      string    `json:"slug"`
	Length    int       `json:"length"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
