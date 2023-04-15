package utils

import (
	"blogger/global"
	"blogger/model/system"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(account, password string) (string, error) {
	//token的有效期
	//token的内容
	now := time.Now()
	claims := system.BaseLogin{
		Account:  account,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Unix(),
			IssuedAt:  now.Add(24 * time.Hour).Unix(),
			Issuer:    "blogger",
		},
	}
	secret := []byte(global.Viper.GetString("jwt.secret"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(secret)
	return tokenStr, err
}
