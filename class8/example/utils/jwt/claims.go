package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"main/utils/cookie"
	"time"
)

type CustomClaims struct {
	BufferTime           int64 // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
	jwt.RegisteredClaims       // token注册信息
	BaseClaims                 // 用户信息
}

type BaseClaims struct {
	Id         int64
	Username   string
	CreateTime time.Time
	UpdateTime time.Time
}

func GetClaims(secret string, cookie *cookie.Cookie) (*CustomClaims, error) {
	var token string
	ok := cookie.Get("x-token", &token)

	//token, err := c.Cookie("x-token")
	if !ok {
		err := errors.New("get token by cookie failed")
		return nil, err
	}
	j := NewJWT(&Config{SecretKey: secret})
	claims, err := j.ParseToken(token)
	if err != nil {
		err := errors.New("parse token failed")
		return nil, err
	}
	return claims, nil
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(secret string, cookie *cookie.Cookie) (*BaseClaims, error) {
	if cl, err := GetClaims(secret, cookie); err != nil {
		return nil, err
	} else {
		return &cl.BaseClaims, nil
	}
}

// GetUserID 获取从jwt解析出来的用户ID
func GetUserID(secret string, cookie *cookie.Cookie) (int64, error) {
	if cl, err := GetClaims(secret, cookie); err != nil {
		return -1, err
	} else {
		return cl.BaseClaims.Id, nil
	}
}
