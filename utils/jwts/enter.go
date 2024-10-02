package jwts

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"todoBackend/app/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JwtPayload struct {
	UserId   int    `json:"user_id"`
	NickName string `json:"nick_name"`
}
type CustomClaims struct {
	JwtPayload
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT令牌
func GenerateToken(payload JwtPayload) (string, error) {
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.Cfg.Jwt.TokenLifeSpan))),
		},
		JwtPayload: payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg.Jwt.ApiSecret))
}
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.Jwt.ApiSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid jwt token")
}

// GetTokenFromHttpHeader 从请求中提取令牌
func GetTokenFromHttpHeader(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(bearerToken) == 0 {
		return ""
	}
	parts := strings.Split(bearerToken, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1] // 返回实际的JWT
	}
	return bearerToken
}

// ExtractTokenID 从请求中提取用户ID
func ExtractTokenID(c *gin.Context) (uint, error) {
	jwtToken := GetTokenFromHttpHeader(c)
	if jwtToken == "" {
		return 0, errors.New("jwt token缺失")
	}
	customClaims, err := ParseToken(jwtToken)
	if err != nil {
		return 0, errors.New("非法的jwt token")
	}
	uid := uint(customClaims.UserId)
	return uid, nil
}

func ExtractUserNickname(c *gin.Context) (string, error) {
	jwtToken := GetTokenFromHttpHeader(c)
	customClaims, err := ParseToken(jwtToken)
	if err != nil {
		return "", err
	}
	return customClaims.NickName, nil
}

// Valid 验证令牌有效性
func Valid(c *gin.Context) error {
	tokenString := GetTokenFromHttpHeader(c)
	if tokenString == "" {
		return errors.New("jwt token缺失")
	}
	_, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Cfg.Jwt.ApiSecret), nil
	})
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}
	return nil
}
