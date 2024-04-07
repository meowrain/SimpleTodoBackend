package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
	"todoBackend/app/config"
)

// GenerateToken 生成JWT令牌
func GenerateToken(user_id uint) (string, error) {
	// 从配置中获取JWT的有效期（token lifespan）
	tokenLifespan, err := strconv.Atoi(config.Cfg.Jwt.TokenLifeSpan)
	if err != nil {
		return "", err
	}

	// 创建JWT的声明（claims）
	claims := jwt.MapClaims{}
	claims["authorized"] = true // 自定义声明：标识该令牌已经授权
	claims["user_id"] = user_id // 自定义声明：存储用户标识
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	// 设置令牌过期时间，当前时间加上有效期的小时数，转换为Unix时间戳

	// 使用HS256算法创建JWT令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用API密钥对JWT令牌进行签名并返回签名后的令牌字符串
	return token.SignedString([]byte(config.Cfg.Jwt.ApiSecret))
}

// ExtractToken 从请求中提取令牌
func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	//头部拿到数据
	//fmt.Println(bearerToken)
	return strings.Split(bearerToken, " ")[0]
}

// ExtractTokenID 从请求中提取用户ID
func ExtractTokenID(c *gin.Context) (uint, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Cfg.Jwt.ApiSecret), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		//fmt.Println(uid)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}

	return 0, nil
}

// Valid 验证令牌有效性
func Valid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	fmt.Println(tokenString)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Cfg.Jwt.ApiSecret), nil
	})
	if err != nil {
		return err
	}
	return nil
}
