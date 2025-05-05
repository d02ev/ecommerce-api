package adapters

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type TokenService struct {
	accessTokenSecretKey string
	refreshTokenSecretKey string
	accessTokenExpirationTime time.Time
	refreshTokenExpirationTime time.Time
}

type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func NewTokenService() *TokenService {
	return &TokenService{
		accessTokenSecretKey: viper.GetString("ACCESS_TOKEN_SECRET"),
		refreshTokenSecretKey: viper.GetString("REFRESH_TOKEN_SECRET"),
		accessTokenExpirationTime: time.Now().Add(viper.GetDuration("ACCESS_TOKEN_EXPIRATION") * time.Hour),
		refreshTokenExpirationTime: time.Now().Add(viper.GetDuration("REFRESH_TOKEN_EXPIRATION") * time.Hour),
	}
}

func (ts *TokenService) GenerateAccessToken(userId, role uint) (string, error) {
	claims := CustomClaims{
		Role: strconv.Itoa(int(role)),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ecommerce-api",
			Subject:   strconv.Itoa(int(userId)),
			ExpiresAt: jwt.NewNumericDate(ts.accessTokenExpirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(ts.accessTokenSecretKey))
}

func (ts *TokenService) GenerateRefreshToken(userId uint) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "ecommerce-api",
		Subject:   strconv.Itoa(int(userId)),
		ExpiresAt: jwt.NewNumericDate(ts.refreshTokenExpirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(ts.refreshTokenSecretKey))
}
