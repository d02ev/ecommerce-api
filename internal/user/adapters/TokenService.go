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
	Admin bool `json:"admin"`
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
	var isAdmin bool = false;
	if role == 1 { isAdmin = true; }

	claims := CustomClaims{
		Admin: isAdmin,
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

func (ts *TokenService) DecodeRefreshToken(token string) (uint, error) {
	claims := &jwt.RegisteredClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ts.refreshTokenSecretKey), nil
	})

	if err != nil || !parsedToken.Valid {
		return 0, err
	}

	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, err
	}

	return uint(userId), nil
}

func (ts *TokenService) DecodeAccessToken(token string) (uint, bool, error) {
	claims := &CustomClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ts.accessTokenSecretKey), nil
	})

	if err != nil || !parsedToken.Valid {
		return 0, false, err
	}

	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, false, err
	}

	return uint(userId), claims.Admin, nil
}
