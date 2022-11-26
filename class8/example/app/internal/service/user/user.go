package user

import (
	"context"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"
	"gorm.io/gorm"
	g "main/app/global"
	"main/app/internal/model"
	"main/utils/jwt"
	"time"
)

type SUser struct{}

var insUser = SUser{}

func (s *SUser) CheckUserIsExist(ctx context.Context, username string) error {
	userSubject := &model.UserSubject{}
	err := g.MysqlDB.WithContext(ctx).
		Table("user_subject").
		Select("username").
		Where("username = ?", username).
		First(userSubject).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			g.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "user_subject"),
			)
			return fmt.Errorf("internal err")
		}
	} else {
		return fmt.Errorf("username already exist")
	}

	return nil
}

func (s *SUser) EncryptPassword(password string) string {
	d := sha3.Sum224([]byte(password))
	return hex.EncodeToString(d[:])
}

func (s *SUser) CreateUser(ctx context.Context, userSubject *model.UserSubject) {
	g.MysqlDB.WithContext(ctx).
		Table("user_subject").
		Create(userSubject)
}

func (s *SUser) CheckPassword(ctx context.Context, userSubject *model.UserSubject) error {
	err := g.MysqlDB.WithContext(ctx).
		Table("user_subject").
		Where(&model.UserSubject{
			Username: userSubject.Username,
			Password: userSubject.Password,
		}).
		First(userSubject).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			g.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "user_subject"),
			)
			return fmt.Errorf("internal err")
		} else {
			return fmt.Errorf("invalid username or password")
		}
	}

	return nil
}

func (s *SUser) GenerateToken(ctx context.Context, userSubject *model.UserSubject) (string, error) {
	jwtConfig := g.Config.Middleware.Jwt

	j := jwt.NewJWT(&jwt.Config{
		SecretKey:   jwtConfig.SecretKey,
		ExpiresTime: jwtConfig.ExpiresTime,
		BufferTime:  jwtConfig.BufferTime,
		Issuer:      jwtConfig.Issuer})
	claims := j.CreateClaims(&jwt.BaseClaims{
		Id:         userSubject.Id,
		Username:   userSubject.Username,
		CreateTime: userSubject.CreateTime,
		UpdateTime: userSubject.UpdateTime,
	})

	tokenString, err := j.GenerateToken(&claims)
	if err != nil {
		g.Logger.Error("generate token failed.", zap.Error(err))
		return "", fmt.Errorf("internal err")
	}

	err = g.Rdb.Set(ctx,
		fmt.Sprintf("jwt:%d", userSubject.Id),
		tokenString,
		time.Duration(jwtConfig.ExpiresTime)*time.Second).Err()
	if err != nil {
		g.Logger.Error("set redis cache failed.",
			zap.Error(err),
			zap.String("key", "jwt:[id]"),
			zap.Int64("id", userSubject.Id),
		)
		return "", fmt.Errorf("internal err")
	}

	return tokenString, nil
}
