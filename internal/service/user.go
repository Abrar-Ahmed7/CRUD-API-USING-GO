package service

import (
	"time"

	"github.com/Abrar-Ahmed7/rest-api-go/internal/model"
	"github.com/Abrar-Ahmed7/rest-api-go/internal/repo/user"
)

type userService struct {
	userRepo user.UserRepo
}

type UserService interface {
	Save() string
}

func NewUserService(ur user.UserRepo) UserService {
	return &userService{
		userRepo: ur,
	}
}

func (us userService) Save() string {
	u := model.User{
		Name:      "Alice",
		Email:     "alice@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := us.userRepo.Save(&u)
	if err != nil {
		return ""
	}
	return "success"
}
