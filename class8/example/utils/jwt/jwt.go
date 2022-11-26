package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWT struct {
	Config *Config
}

type Config struct {
	SecretKey   string // 密钥
	ExpiresTime int64  // 过期时间,单位:秒
	BufferTime  int64  // 缓冲时间,缓冲时间内会获得新的token刷新令牌,此时一个用户会存在两个有效令牌,但是前端只留一个,另一个会丢失
	Issuer      string // 签发者
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT(config *Config) *JWT {
	return &JWT{Config: config}
}

func (j *JWT) CreateClaims(baseClaims *BaseClaims) CustomClaims {
	claims := CustomClaims{
		BufferTime: j.Config.BufferTime,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Truncate(time.Second)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.Config.ExpiresTime) * time.Second)),
			Issuer:    j.Config.Issuer,
		},
		BaseClaims: *baseClaims,
	}
	return claims
}

func (j *JWT) GenerateToken(claims *CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	signingKey := []byte(j.Config.SecretKey)
	return token.SignedString(signingKey)
}

// ParseToken 解析JWT
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	signingKey := []byte(j.Config.SecretKey)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return signingKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
