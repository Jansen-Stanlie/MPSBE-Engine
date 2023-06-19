package models

import "github.com/dgrijalva/jwt-go"

type JWTClaim struct {
	Signature string `json:"signature"`
	Timestamp string `json:"timestamp"`
	jwt.StandardClaims
}
