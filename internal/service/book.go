package service

import (
	"time"

	"github.com/Abrar-Ahmed7/rest-api-go/internal/model"
	"github.com/Abrar-Ahmed7/rest-api-go/internal/repo/book"
)

type bookService struct {
	bookRepo book.BookRepo
}

type BookService interface {
	Save(userId int, title, authorName, publication string) error
	GetBooks() ([]*model.Book, error)
	Update(bookId, userId int, title, authorName, publication string) error
	Delete(bookId, userId int) error
}

func NewBookService(br book.BookRepo) BookService {
	return &bookService{
		bookRepo: br,
	}
}

func (bs bookService) Save(userId int, title, authorName, publication string) error {
	b := model.Book{
		UserId:      userId,
		Title:       title,
		AuthorName:  authorName,
		Publication: publication,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := bs.bookRepo.Save(&b)
	if err != nil {
		return err
	}
	return nil
}

func (bs bookService) GetBooks() ([]*model.Book, error) {
	return bs.bookRepo.GetBooks()
}

func (bs bookService) Update(bookId, userId int, title, authorName, publication string) error {
	b := &model.Book{
		ID:          bookId,
		UserId:      userId,
		Title:       title,
		AuthorName:  authorName,
		Publication: publication,
		UpdatedAt:   time.Now(),
	}
	return bs.bookRepo.UpdateBook(b)
}

func (bs bookService) Delete(bookId, userId int) error {
	b := &model.Book{
		ID:     bookId,
		UserId: userId,
	}
	return bs.bookRepo.DeleteBook(b)
}
