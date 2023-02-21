package gormimpl

import (
	"github.com/Abrar-Ahmed7/rest-api-go/internal/db"
	"github.com/Abrar-Ahmed7/rest-api-go/internal/model"
	"github.com/Abrar-Ahmed7/rest-api-go/internal/repo/user"
	"gorm.io/gorm"
)

type GormUserRepo struct {
	DB *gorm.DB
}

func NewGormUserRepo(db *db.Connection) user.UserRepo {
	return &GormUserRepo{
		DB: db.DB,
	}
}

// Save implements user.UserRepo
func (gur *GormUserRepo) Save(u *model.User) error {
	if err := gur.DB.Save(u).Error; err != nil {
		return err
	}

	return nil
}
