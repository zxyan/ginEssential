package common

import (
	"ctjsoft/ginessential/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_crect") // 加密密钥

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// ReleaseToken 发放 token
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // token 有效期设置为 7 天
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // token 过期时间
			IssuedAt:  time.Now().Unix(),     // token 发放时间
			Issuer:    "ctjsoft.com",         // 发放 token 的人
			Subject:   "user token",          // 主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) // 生成 token
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 从 tokenString 中解析 Claims, 然后返回
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
