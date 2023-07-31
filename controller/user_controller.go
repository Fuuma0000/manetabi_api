package controller

import (
	"github.com/Fuuma0000/manetabi_api/usecase"
)

type IUserController interface {
	// SignUp(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// func SignUp(c echo.Context) error {
// 	// ユーザー情報を格納する構造体
// 	user := model.User{}
// 	// リクエストボディを構造体にバインド
// 	if err := c.Bind(&user); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	// メールアドレスの重複チェック
// 	b, message := checkDuplicateEmail(user.Email)
// 	if b { // 重複あり
// 		return c.JSON(http.StatusBadRequest, message)
// 	}

// 	// パスワードのハッシュ化
// 	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	// ユーザー情報をDBに登録
// 	newUser := model.User{
// 		UserName:         user.UserName,
// 		Email:            user.Email,
// 		Password:         string(hash),
// 		ProfileImagePath: user.ProfileImagePath,
// 	}
// 	q := `INSERT INTO users (user_name, email, password, profile_image_path) VALUES (?, ?, ?, ?)`
// 	_, err = db.Exec(q, newUser.UserName, newUser.Email, newUser.Password, newUser.ProfileImagePath)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	return c.String(http.StatusOK, "Hello, World!")
// }

// // メールアドレスの重複チェック
// func checkDuplicateEmail(email string) (bool, string) {
// 	q := `SELECT COUNT(*) FROM users WHERE email = ?`
// 	row := db.QueryRow(q, email)

// 	var count int
// 	err := row.Scan(&count)
// 	if err != nil {
// 		fmt.Println(err)
// 		return true, "メールアドレスの重複チェックに失敗しました"
// 	}

// 	if count > 0 {
// 		return true, "メールアドレスが重複しています"
// 	}

// 	return false, ""
// }
