package infrastructure

import (
	"database/sql"
	"log"
	"time"

	"github.com/hmarf/sample_layer/domain/model"

	"github.com/hmarf/sample_layer/domain/repository"
)

type userInfraStruct struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) repository.UserRepository {
	return &userInfraStruct{db: db}
}

func (u *userInfraStruct) Insert(userID string, name string, createdAt time.Time) (err error) {
	ins, err := u.db.Prepare("INSERT INTO user(user_id,name,createdAt) VALUES(?,?,?)")
	if err != nil {
		log.Println(err)
	}
	ins.Exec(userID, name, createdAt)
	return
}

func (u *userInfraStruct) Select(userID string) (user model.User, err error) {
	if err := u.db.QueryRow("SELECT * FROM user WHERE user_id = ?", userID).
		Scan(&user.UserID, &user.Name, &user.CreatedAt); err != nil {
	}
	return
}

func (u *userInfraStruct) Delete(userID string) (err error) {
	ins, err := u.db.Prepare("DELETE FROM user WHERE user_id=?")
	if err != nil {
		log.Println(err)
	}
	ins.Exec(userID)
	return
}
