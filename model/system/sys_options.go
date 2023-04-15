package system

import "github.com/dgrijalva/jwt-go"

type BaseLogin struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
	Captcha  string `json:"captcha"`
	jwt.StandardClaims
}
