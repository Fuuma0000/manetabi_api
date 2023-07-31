package usecase

import (
	"github.com/Fuuma0000/manetabi_api/infrastructure"
	"github.com/Fuuma0000/manetabi_api/model"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
}

type userUsecase struct {
	ui infrastructure.IUserInfrastructer
}

func NewUserUsecase(ui infrastructure.IUserInfrastructer) IUserUsecase {
	return &userUsecase{ui}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{
		UserName:         user.UserName,
		Email:            user.Email,
		Password:         string(hash),
		ProfileImagePath: user.ProfileImagePath,
	}
	if err := uu.ui.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:               newUser.ID,
		UserName:         newUser.UserName,
		Email:            newUser.Email,
		ProfileImagePath: newUser.ProfileImagePath,
	}
	return resUser, nil
}
