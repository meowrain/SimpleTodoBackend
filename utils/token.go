package utils

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"todoBackend/app/config"
)

func GenerateToken(user_id uint) (string, error) {
	token_lifespan, err := strconv.Atoi(config.Cfg.Jwt.TokenLifeSpan)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Cfg.Jwt.ApiSecret))
}
