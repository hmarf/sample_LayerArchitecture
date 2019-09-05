package handler

import (
	"github.com/hmarf/sample_layer/domain/service"
	"github.com/hmarf/sample_layer/infrastructure"
)

var userService service.UserServiceInterface

func DI() {
	userRepo := infrastructure.NewUserDB(infrastructure.DB)
	userService = service.NewUserService(userRepo)
}
