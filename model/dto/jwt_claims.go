package dto

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	jwt.RegisteredClaims
	UserId int64 `json:"userId"`
}
