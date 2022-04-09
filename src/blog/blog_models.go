package blog

import "time"

type BlogModel struct {
	Id        int       `json:"id"`
	Title     string    `json:"title" gorm:"unique;not null"`
	Body      string    `json:"body" gorm:"not null"`
	Thumbnail string    `json:"thumbnail"`
	Autor     string    `json:"autor" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b BlogModel) validate() string {
	if b.Title == "" {
		return "title is required"
	}

	if b.Body == "" {
		return "body is required"
	}

	return ""
}
