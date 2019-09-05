package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/hmarf/sample_layer/domain/model"
)

func (u *userServiceStruct) InsertUserService(name string) (user model.User, err error) {

	// userID　を作成する
	userID, err := uuid.NewRandom()
	if err != nil {
		return
	}

	nowTime := time.Now()
	user.UserID = userID.String()
	user.Name = name
	user.CreatedAt = nowTime.String()
	err = u.userRepo.Insert(user.UserID, user.Name, nowTime)
	return
}
