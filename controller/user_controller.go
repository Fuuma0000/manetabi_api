package controller

import (
	"net/http"
	"time"

	"github.com/Fuuma0000/manetabi_api/interface/presenter"
	"github.com/Fuuma0000/manetabi_api/model"
	"github.com/Fuuma0000/manetabi_api/usecase"
	"github.com/labstack/echo"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
}

type userController struct {
	uu  usecase.IUserUsecase
	jwt presenter.JWTHandler
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{
		uu: uu,
	}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	b, message, err := uc.uu.CheckDuplicateEmail(user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if b { // 重複あり
		return c.JSON(http.StatusBadRequest, message)
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := uc.uu.Login(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// JWTトークンを生成
	expiration := time.Hour * 24 // トークンの有効期限を1日に設定
	token, err := uc.jwt.GenerateJWTToken(user.ID, expiration)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// レスポンスにトークンを含めて返す
	response := map[string]string{
		"token": token,
	}
	return c.JSON(http.StatusOK, response)
}
