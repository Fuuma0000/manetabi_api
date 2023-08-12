package handler

import (
	"github.com/Fuuma0000/manetabi_api/controller"
	"github.com/Fuuma0000/manetabi_api/interface/middleware"
	"github.com/Fuuma0000/manetabi_api/interface/presenter"
	"github.com/labstack/echo"
)

func NewRouter(uc controller.IUserController, pc controller.IPlanController, jwtHandler *presenter.JWTHandler) *echo.Echo {
	e := echo.New()

	// ここにルーティングを書いていく
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	p := e.Group("/plans")
	// JWTMiddlewareを適用
	p.Use(middleware.JWTMiddleware(*jwtHandler))
	p.POST("", pc.CreatePlan)
	p.GET("", pc.GetPlansByUserID)
	p.GET("/:planId", pc.GetPlanByID)
	p.PUT("/:planId", pc.UpdatePlan)
	p.DELETE("/:planId", pc.DeletePlan)
	return e
}
