package service

import (
	"github.com/hmarf/sample_layer/domain/model"
	"github.com/hmarf/sample_layer/domain/repository"
)

type userServiceStruct struct {
	userRepo repository.UserRepository
}

type UserServiceInterface interface {
	GetUserService(string) (model.User, error)
	InsertUserService(string) (model.User, error)
	DeleteUserService(string) error
}

func NewUserService(u repository.UserRepository) UserServiceInterface {
	return &userServiceStruct{userRepo: u}
}
