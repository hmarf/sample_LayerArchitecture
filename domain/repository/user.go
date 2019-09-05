package repository

import (
	"time"

	"github.com/hmarf/sample_layer/domain/model"
)

// UserRepository はUserTable
type UserRepository interface {
	Insert(string, string, time.Time) error
	Select(string) (model.User, error)
	Delete(string) error
}
