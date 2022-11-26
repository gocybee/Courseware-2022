package boot

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	g "main/app/global"
	"time"
)

func MysqlDBSetup() {
	config := g.Config.DataBase.Mysql

	db, err := gorm.Open(mysql.Open(config.GetDsn()))
	if err != nil {
		g.Logger.Fatal("initialize mysql failed.", zap.Error(err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(g.Config.DataBase.Mysql.GetConnMaxIdleTime())
	sqlDB.SetConnMaxLifetime(g.Config.DataBase.Mysql.GetConnMaxLifeTime())
	sqlDB.SetMaxIdleConns(g.Config.DataBase.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(g.Config.DataBase.Mysql.MaxOpenConns)
	err = sqlDB.Ping()
	if err != nil {
		g.Logger.Fatal("connect to mysql db failed.", zap.Error(err))
	}

	g.MysqlDB = db

	g.Logger.Info("initialize mysql successfully!")
}

func RedisSetup() {
	config := g.Config.DataBase.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       config.Db,
		PoolSize: config.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		g.Logger.Fatal("connect to redis instance failed.", zap.Error(err))
	}

	g.Rdb = rdb

	g.Logger.Info("initialize redis client successfully!")
}
