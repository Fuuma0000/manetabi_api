package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Fuuma0000/manetabi_api/model"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // グローバル変数としてDB接続を定義

func main() {
	// データベース接続
	var err error
	db, err = sql.Open("mysql", "fuuma:password@tcp(127.0.0.1:3308)/manetabi_db")
	if err != nil {
		fmt.Println(err)
		return // データベース接続が失敗した場合にプログラムを終了する
	}
	defer db.Close()

	fmt.Println("Connected to DB!")

	// APIサーバー起動
	e := echo.New()
	e.POST("/signup", SignUp)

	// サーバーを開始
	e.Logger.Fatal(e.Start(":8080"))
}

func SignUp(c echo.Context) error {
	// ユーザー情報を格納する構造体
	user := model.User{}
	// リクエストボディを構造体にバインド
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// パスワードのハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// ユーザー情報をDBに登録
	newUser := model.User{
		UserName:         user.UserName,
		Email:            user.Email,
		Password:         string(hash),
		ProfileImagePath: user.ProfileImagePath,
	}
	q := `INSERT INTO users (user_name, email, password, profile_image_path) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(q, newUser.UserName, newUser.Email, newUser.Password, newUser.ProfileImagePath)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "Hello, World!")
}
