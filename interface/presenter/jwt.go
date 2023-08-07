package presenter

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type JWTHandler struct {
	secretKey []byte
}

func NewJWTHandler(secretKey []byte) *JWTHandler {
	return &JWTHandler{
		secretKey: secretKey,
	}
}

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// JWTトークンの生成
func (jh *JWTHandler) GenerateJWTToken(userID uint, expiration time.Duration) (string, error) {
	// JWTのペイロード部分
	claims := JWTClaims{
		// ユーザーID
		UserID: userID,
		// 有効期限
		StandardClaims: jwt.StandardClaims{
			// 有効期限の設定
			ExpiresAt: jwt.At(time.Now().Add(expiration)),
		},
	}
	// JWTの署名部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 署名
	tokenString, err := token.SignedString(jh.secretKey)
	// 署名に失敗した場合
	if err != nil {
		return "", err
	}

	// 署名に成功した場合は、トークンを返す
	return tokenString, nil
}

// JWTトークンの検証
func (jh *JWTHandler) VerifyJWTToken(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	// トークンを検証
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jh.secretKey, nil
	})
	// 検証に失敗した場合
	if err != nil {
		return nil, err
	}

	// 検証に成功したが、トークンが無効の場合
	if !token.Valid {
		return nil, err
	}

	// 検証に成功した場合は、クレームを返す
	return claims, nil
}
