package usecase

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"simpson/config"
	"simpson/internal/dto"

	"github.com/dgrijalva/jwt-go"
)

type JwtUsecase interface {
	GeneratorToken(ctx context.Context, req dto.JwtReq) (string, error)
	VerifyToken(ctx context.Context, tokenStr string) (dto.JwtClaim, error)
}

type jwtUsecase struct {
	cfg        *config.Config
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
	signMethod jwt.SigningMethod
}

func NewJwtUsecase(
	cfg *config.Config,
	pri *rsa.PrivateKey,
	pub *rsa.PublicKey,
	sign jwt.SigningMethod,
) JwtUsecase {
	return &jwtUsecase{
		cfg:        cfg,
		publicKey:  pub,
		privateKey: pri,
		signMethod: sign,
	}
}

func ParseKey(cfg *config.Config) (*rsa.PrivateKey, *rsa.PublicKey, jwt.SigningMethod, error) {
	var (
		private *rsa.PrivateKey
		public  *rsa.PublicKey
		sign    jwt.SigningMethod
		err     error
	)

	base64.StdEncoding.DecodeString(cfg.JWT.PrivateKey)
	privateByte, err := base64.StdEncoding.DecodeString(cfg.JWT.PrivateKey)
	if err != nil {
		return private, public, sign, err
	}
	private, err = jwt.ParseRSAPrivateKeyFromPEM(privateByte)
	if err != nil {
		return private, public, sign, err
	}
	publicByte, err := base64.StdEncoding.DecodeString(cfg.JWT.PublicKey)
	if err != nil {
		return private, public, sign, err
	}
	public, err = jwt.ParseRSAPublicKeyFromPEM(publicByte)
	if err != nil {
		return private, public, sign, err
	}
	return private, public, sign, err
}

func (j *jwtUsecase) GeneratorToken(ctx context.Context, req dto.JwtReq) (string, error) {
	var (
		tokenStr string
		err      error
	)
	jwtToken := jwt.New(j.signMethod)
	jwtClaim := dto.JwtClaim{
		Username: req.Username,
		UserID:   req.UserID,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: int64(j.cfg.JWT.ShortTokenExpireTime),
			Issuer:    j.cfg.JWT.Issuer,
		},
	}
	jwtToken.Claims = jwtClaim
	tokenStr, err = jwtToken.SignedString(j.privateKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
func (j *jwtUsecase) VerifyToken(ctx context.Context, tokenStr string) (dto.JwtClaim, error) {
	var (
		claims dto.JwtClaim
		err    error
	)
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return j.publicKey, nil
	}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, keyFunc)
	jwtErr, _ := err.(*jwt.ValidationError)
	if jwtErr != nil && jwtErr.Errors == jwt.ValidationErrorExpired {
		return claims, errors.New("token expired")
	}
	if err != nil || !token.Valid {
		return claims, errors.New("token invalid")
	}
	return claims, nil
}
