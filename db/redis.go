package db

import (
	"github.com/gomodule/redigo/redis"
	"github.com/notblessy/takeme-backend/config"
	"github.com/sirupsen/logrus"
)

// RedisConnectionPool :nodoc:
func RedisConnectionPool() *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(
				"tcp",
				config.RedisHost(),
				redis.DialDatabase(config.RedisDB()),
			)
			if err != nil {
				logrus.Fatal("Failed to connect redis")
			}
			return conn, err
		},
	}

	return pool
}
