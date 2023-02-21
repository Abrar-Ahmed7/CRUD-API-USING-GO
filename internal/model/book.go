package model

import "time"

type Book struct {
	ID          int
	UserId      int
	Title       string
	AuthorName  string
	Publication string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
