package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"rongqin.cn/todo_task/conf"
)

type JWTClaims struct {
	UID      uint   // 用户ID
	Username string // 用户名
	jwt.StandardClaims
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	ExpiresAt    int64  `json:"expiresAt"`
	RefreshToken string `json:"refreshToken"`
}

// token 接口
type ITokenService interface {
	// token生成
	Encode(uid uint, username string) (string, error)
	// 解析token
	Decode(token string) *JWTClaims
}

type TokenService struct{}

// 生成token
func (ts *TokenService) Encode(uid uint, username string) (*Token, error) {

	expireTime := time.Now().Add(time.Hour * time.Duration(conf.Config.SecurityConfig.ExpireTime)).Unix() // 过期时间
	claims := JWTClaims{uid, username, jwt.StandardClaims{ExpiresAt: expireTime}}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString([]byte(conf.Config.SecurityConfig.PrivateKey))
	if err != nil {
		return nil, err
	}
	return &Token{AccessToken: tokenStr, ExpiresAt: expireTime}, nil
}

// 解析token
func (ts *TokenService) Decode(tokenStr string) (*JWTClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.Config.SecurityConfig.PrivateKey), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {

		return claims, nil
	}
	return nil, errors.New("token 解析失败")
}
