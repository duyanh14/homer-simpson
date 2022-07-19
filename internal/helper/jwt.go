package helper

import (
	"context"
	"fmt"
	"simpson/internal/dto"

	"github.com/dgrijalva/jwt-go"
)

func GeneratorToken(ctx context.Context, req dto.JwtReq, cfg dto.JwtConfig) (string, error) {
	var (
		token string
		err   error
	)
	jwtSignMethod := jwt.GetSigningMethod(cfg.SigningMethod)
	jwtToken := jwt.New(jwtSignMethod)

	jwtClaim := dto.JwtClaim{
		Username: req.Username,
		UserID:   req.UserID,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: cfg.ExpiresAt,
			Issuer:    cfg.Issuer,
		},
	}
	jwtToken.Claims = jwtClaim

	fmt.Println(jwtToken)
	return token, err
}

func VerifyToken(ctx context.Context, token, tokenRefesh string) (*dto.JwtClaim, error) {
	var (
		resp = &dto.JwtClaim{}
		err  error
	)
	return resp, err
}
