package main

import (
	"github.com/Fuuma0000/manetabi_api/controller"
	"github.com/Fuuma0000/manetabi_api/db"
	"github.com/Fuuma0000/manetabi_api/infrastructure"
	"github.com/Fuuma0000/manetabi_api/interface/handler"
	"github.com/Fuuma0000/manetabi_api/usecase"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// DB接続
	db := db.NewDB()
	userInfrastructure := infrastructure.NewUserInfrastructer(db)
	planInfrastructure := infrastructure.NewPlanInfrastructer(db)
	userUsecase := usecase.NewUserUsecase(userInfrastructure)
	planUsecase := usecase.NewPlanUsecase(planInfrastructure)
	userController := controller.NewUserController(userUsecase)
	planController := controller.NewPlanController(planUsecase)

	// サーバーを開始
	e := handler.NewRouter(userController, planController)
	e.Logger.Fatal(e.Start(":8080"))
}
