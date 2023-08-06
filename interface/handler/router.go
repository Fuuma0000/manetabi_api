package handler

import (
	"github.com/Fuuma0000/manetabi_api/controller"
	"github.com/labstack/echo"
)

func NewRouter(uc controller.IUserController, pc controller.IPlanController) *echo.Echo {
	e := echo.New()

	// ここにルーティングを書いていく
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	p := e.Group("/plans")
	p.POST("", pc.CreatePlan)
	p.GET("", pc.GetPlansByUserID)
	p.GET("/:id", pc.GetPlanByID)
	return e
}
