package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var RedisConn *redis.Client

func RedisConnection() error {
	connectionString := fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	db, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		zap.L().Fatal("Error converting database string to int", zap.Error(err))
	}
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     connectionString,
		Password: os.Getenv("REDIS_PASS"),
		DB:       db,
	})

	pong, err := RedisConn.Ping().Result()
	if err != nil {
		return err
	}

	zap.L().Info("Redis connected", zap.String("ping", pong))
	return nil
}
