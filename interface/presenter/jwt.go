package presenter

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
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
			ExpiresAt: time.Now().Add(expiration).Unix(),
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

	// Token verification
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method is HS256 and return the secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return jh.secretKey, nil
		}
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	})

	// If verification fails
	if err != nil {
		return nil, err
	}

	// If token is invalid
	if !token.Valid {
		return nil, err
	}

	// If verification is successful, return the claims
	return claims, nil
}
