package usecase

import (
	"fmt"

	"github.com/Fuuma0000/manetabi_api/infrastructure"
	"github.com/Fuuma0000/manetabi_api/model"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) error
	CheckDuplicateEmail(email string) (bool, string, error)
}

type userUsecase struct {
	ui infrastructure.IUserInfrastructer
}

func NewUserUsecase(ui infrastructure.IUserInfrastructer) IUserUsecase {
	return &userUsecase{ui}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	// TODO: メールアドレスのバリデーションをしないと空文字を許容してしまう
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

func (uu *userUsecase) Login(user model.User) error {
	storedUser := model.User{}
	if err := uu.ui.GetUserByEmail(&storedUser, user.Email); err != nil {
		return err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return err
	}
	return nil
}

func (uu *userUsecase) CheckDuplicateEmail(email string) (bool, string, error) {
	b, str, err := uu.ui.CheckDuplicateEmail(email)
	if err != nil {
		fmt.Println(err)
		return b, str, err
	}
	return b, str, nil
}
