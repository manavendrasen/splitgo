package util

import "github.com/golang-jwt/jwt/v5"

type AccessTokenClaim struct {
	ID          uint   `json:"ID"`
	DisplayName string `json:"DisplayName"`
	Email       string `json:"Email"`
	jwt.RegisteredClaims
}