package service

import (
	"log"
)

func (u *userServiceStruct) DeleteUserService(userID string) (err error) {

	err = u.userRepo.Delete(userID)
	if err != nil {
		log.Println(err)
	}
	return
}
