package jwt

import (
	"fmt"
	"os"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

func GetClaimsFromToken(token string) (*JwtClaims, error) {
	//? kembalikan token tadi hingga menjadi entity jwt
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, utils.UnauthenticatedResponse()
	}

	if !tokenClaims.Valid {
		return nil, utils.UnauthenticatedResponse()
	}

	if claims, ok := tokenClaims.Claims.(*JwtClaims); !ok {
		return claims, nil
	}

	return nil, utils.UnauthenticatedResponse()
}
