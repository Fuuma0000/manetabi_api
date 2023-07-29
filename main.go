package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

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
	e.POST("/login", SignUp)

	// サーバーを開始
	e.Logger.Fatal(e.Start(":8080"))
}

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func SignUp(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	_, err = db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", newUser.Email, newUser.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Hello, World!")
}
