package models

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	ID          uint
	Nickname    string
	AuthorityId uint
	jwt.StandardClaims
}
