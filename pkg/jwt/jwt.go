package jwt_client

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

type JWTClaim struct {
	UserID   uint64
	RoleID   uint64
	IsVerify bool
	jwt.RegisteredClaims
}
