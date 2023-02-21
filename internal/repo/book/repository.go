package book

import "github.com/Abrar-Ahmed7/rest-api-go/internal/model"

type BookRepo interface {
	Save(*model.Book) error
	GetBooks() ([]*model.Book, error)
	UpdateBook(*model.Book) error
	DeleteBook(*model.Book) error
}
