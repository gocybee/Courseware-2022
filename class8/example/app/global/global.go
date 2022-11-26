package global

import (
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"main/app/internal/model/config"
)

var (
	Config *config.Config

	Logger *zap.Logger

	MysqlDB *gorm.DB
	Rdb     *redis.Client
)
