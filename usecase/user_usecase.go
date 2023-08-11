package usecase

import (
	"fmt"
	"time"

	"github.com/Fuuma0000/manetabi_api/infrastructure"
	"github.com/Fuuma0000/manetabi_api/interface/presenter"
	"github.com/Fuuma0000/manetabi_api/model"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (map[string]string, error)
	CheckDuplicateEmail(email string) (bool, error)
}

type userUsecase struct {
	ui  infrastructure.IUserInfrastructer
	jwt presenter.JWTHandler
}

func NewUserUsecase(ui infrastructure.IUserInfrastructer, jwt presenter.JWTHandler) IUserUsecase {
	return &userUsecase{
		ui, jwt,
	}
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

func (uu *userUsecase) Login(user model.User) (map[string]string, error) {
	storedUser := model.User{}
	if err := uu.ui.GetUserByEmail(&storedUser, user.Email); err != nil {
		return nil, err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))

	if err != nil {
		return nil, err
	}
	// JWTトークンを生成
	expiration := time.Hour * 24 // トークンの有効期限を1日に設定
	token, err := uu.jwt.GenerateJWTToken(storedUser.ID, expiration)
	if err != nil {
		return nil, err
	}
	// レスポンスにトークンを含めて返す
	response := map[string]string{
		"token": token,
	}
	return response, nil
}

func (uu *userUsecase) CheckDuplicateEmail(email string) (bool, error) {
	b, err := uu.ui.CheckDuplicateEmail(email)
	if err != nil {
		fmt.Println(err)
		return b, err
	}
	return b, nil
}
