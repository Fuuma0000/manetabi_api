package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Fuuma0000/manetabi_api/interface/presenter"
	"github.com/labstack/echo"
)

// JWTMiddleware はJWTトークンを検証するミドルウェア
func JWTMiddleware(jwtHandler presenter.JWTHandler) echo.MiddlewareFunc {
	// ミドルウェア関数を返す
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		// ハンドラー関数を返す
		return func(c echo.Context) error {
			// Authorizationヘッダーからトークンを取得
			tokenString := c.Request().Header.Get("Authorization")
			tokenString = removeBearerPrefix(tokenString)
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, "Authorization header is missing")
			}

			// トークンを検証
			claims, err := jwtHandler.VerifyJWTToken(tokenString)
			if err != nil {
				fmt.Println(err)
				return c.JSON(http.StatusUnauthorized, "Invalid token")
			}

			// 検証に成功したら、クレーム情報をContextにセットしてハンドラー関数を実行
			c.Set("userID", claims.UserID)
			return next(c)
		}
	}
}

func removeBearerPrefix(tokenString string) string {
	// トークン文字列が "Bearer "で始まっているか確認
	if len(tokenString) > 7 && strings.ToLower(tokenString[0:7]) == "bearer " {
		// "Bearer "を除いたトークン文字列を返す
		return tokenString[7:]
	}
	return tokenString
}
