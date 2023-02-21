package user

import "github.com/Abrar-Ahmed7/rest-api-go/internal/model"

type UserRepo interface {
	Save(*model.User) error
}
