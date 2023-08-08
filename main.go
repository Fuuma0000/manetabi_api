package main

import (
	"os"

	"github.com/Fuuma0000/manetabi_api/controller"
	"github.com/Fuuma0000/manetabi_api/db"
	"github.com/Fuuma0000/manetabi_api/infrastructure"
	"github.com/Fuuma0000/manetabi_api/interface/handler"
	"github.com/Fuuma0000/manetabi_api/interface/presenter"
	"github.com/Fuuma0000/manetabi_api/usecase"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// DB接続
	db := db.NewDB()
	// JWTのシークレットキー
	secretKey := []byte(os.Getenv("SECRET"))
	// JWTHandlerのインスタンス作成
	jwtHandler := presenter.NewJWTHandler(secretKey)

	// インフラ層
	userInfrastructure := infrastructure.NewUserInfrastructer(db)
	planInfrastructure := infrastructure.NewPlanInfrastructer(db)
	// ユースケース層
	userUsecase := usecase.NewUserUsecase(userInfrastructure, *jwtHandler)
	planUsecase := usecase.NewPlanUsecase(planInfrastructure)
	// コントローラー層
	userController := controller.NewUserController(userUsecase)
	planController := controller.NewPlanController(planUsecase)

	// サーバーを開始
	e := handler.NewRouter(userController, planController, jwtHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
