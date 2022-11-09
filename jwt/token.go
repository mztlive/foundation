package jwt

import (
	"errors"
	"time"

	libjwt "github.com/golang-jwt/jwt/v4"
)

// AuthClaim 用户认证信息
type AuthClaim struct {
	UserID string `json:"userId"`
	libjwt.StandardClaims
}

var secret = []byte("adsystem-crm-secret")

// TokenExpireDuration token过期时间
const TokenExpireDuration = 120 * time.Hour

// GetToken 获得token
func GetToken(userID string) (string, error) {
	c := AuthClaim{
		UserID: userID,
		StandardClaims: libjwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "TEST001",
		},
	}

	//使用指定的签名方法创建签名对象
	token := libjwt.NewWithClaims(libjwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(secret)
}

// DeToken 解析token
func DeToken(tokenStr string) (*AuthClaim, error) {
	token, err := libjwt.ParseWithClaims(tokenStr, &AuthClaim{}, func(tk *libjwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claim, ok := token.Claims.(*AuthClaim); ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("invalid token ")
}
