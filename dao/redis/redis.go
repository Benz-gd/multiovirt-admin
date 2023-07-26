package redis

import (
	"context"
	redisV1 "github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"multiovirt-admin/settings"
)

var Rdb *redisV1.Client

func Init(cfg *settings.RedisConfig) *redisV1.Client {
	var ctx = context.Background()
	redisaddr := viper.GetString("redis.host")
	redisuser := viper.GetString("redis.user")
	redispwd := viper.GetString("redis.password")
	redisdb := viper.GetInt("redis.db")
	redis := redisV1.NewClient(&redisV1.Options{
		Addr:     redisaddr,
		Username: redisuser,
		Password: redispwd,
		DB:       redisdb,
	})
	_, err := redis.Ping(ctx).Result()
	if err != nil {
		//fmt.Println(err)
		zap.L().Error("func initredis:", zap.Error(err))
	}
	return redis
}

func Close() {
	_ = Rdb.Close()
}
