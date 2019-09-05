package service

import (
	"log"

	"github.com/hmarf/sample_layer/domain/model"
)

func (u *userServiceStruct) GetUserService(userID string) (model.User, error) {

	user, err := u.userRepo.Select(userID)
	if err != nil {
		log.Println(err)
	}
	return user, err
}
