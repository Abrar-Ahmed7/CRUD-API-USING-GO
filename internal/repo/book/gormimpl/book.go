package gormimpl

import (
	"errors"
	"fmt"

	"github.com/Abrar-Ahmed7/rest-api-go/internal/db"
	"github.com/Abrar-Ahmed7/rest-api-go/internal/model"
	"github.com/Abrar-Ahmed7/rest-api-go/internal/repo/book"
	"gorm.io/gorm"
)

type GormBookRepo struct {
	DB *gorm.DB
}

func NewGormBookRepo(db *db.Connection) book.BookRepo {
	return &GormBookRepo{
		DB: db.DB,
	}
}

func (gbr *GormBookRepo) Save(b *model.Book) error {
	if err := gbr.DB.Save(b).Error; err != nil {
		return err
	}
	return nil
}

func (gbr *GormBookRepo) GetBooks() ([]*model.Book, error) {
	var books []*model.Book
	if err := gbr.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// TODO: in update and delete check if the book belong to that user.
// Also, check if the user is there
func (gbr *GormBookRepo) UpdateBook(b *model.Book) error {

	var id int
	tx := gbr.DB.Raw("SELECT id FROM book WHERE id = ? and user_id = ?;", b.ID, b.UserId).Scan(&id)
	fmt.Println(id)
	if tx.RowsAffected == 0 {
		fmt.Println("User Not Found")
		return errors.New("user or book not found")
	}
	err := gbr.DB.Model(&b).Where("id = ? and user_id = ?", b.ID, b.UserId).Updates(&b).Error
	fmt.Println("Err: ", err)
	return err
}

func (gbr *GormBookRepo) DeleteBook(b *model.Book) error {
	var id int
	tx := gbr.DB.Raw("SELECT id FROM book WHERE id = ? and user_id = ?;", b.ID, b.UserId).Scan(&id)
	fmt.Println(id)
	if tx.RowsAffected == 0 {
		fmt.Println("User Not Found")
		return errors.New("user or book not found")
	}
	err := gbr.DB.Model(&b).Where("id = ? and user_id = ?", b.ID, b.UserId).Delete(&b).Error
	// err := gbr.DB.Updates(&b).Error
	fmt.Println("Err: ", err)
	return err
}
