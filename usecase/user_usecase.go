package usecase

import (
	infrastructer "github.com/Fuuma0000/manetabi_api/infrastructure"
)

type IUserUsecase interface {
	// SignUp(user model.User) (model.UserResponse, error)
}

type userUsecase struct {
	ui infrastructer.IUserInfrastructer
}

func NewUserUsecase(ui infrastructer.IUserInfrastructer) IUserUsecase {
	return &userUsecase{ui}
}
