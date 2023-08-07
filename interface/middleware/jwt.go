package middleware

import (
	"net/http"

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
			// トークンがなければエラー
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, "Authorization header is missing")
			}

			// トークンを検証
			claims, err := jwtHandler.VerifyJWTToken(tokenString)
			// 検証に失敗したらエラー
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "Invalid token")
			}

			// 検証に成功したら、クレーム情報をContextにセットしてハンドラー関数を実行
			c.Set("userID", claims.UserID)
			return next(c)
		}
	}
}
