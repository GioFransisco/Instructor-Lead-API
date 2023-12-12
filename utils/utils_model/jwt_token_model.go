package utilsmodel

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaimsToken struct {
	jwt.RegisteredClaims
	UserId string `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}
