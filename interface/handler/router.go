package handler

import (
	"github.com/Fuuma0000/manetabi_api/controller"
	"github.com/labstack/echo"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	// e.POST("/signup", uc.SignUp)
	return e
}
