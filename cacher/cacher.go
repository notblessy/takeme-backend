// Package cacher implements caching with redis
package cacher

import (
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
)

const defaultTTL = 24 * time.Hour

// ErrKeyNotExist :nodoc:
var ErrKeyNotExist = errors.New("key not exist")

// Cacher :nodoc:
type Cacher interface {
	Get(key string) (interface{}, error)
	Store(item Item) error

	SetRedisConnectionPool(*redis.Pool)
}

type cacher struct {
	redis      *redis.Pool
	defaultTTL time.Duration
}

// NewCacher :nodoc:
func NewCacher() Cacher {
	return &cacher{
		defaultTTL: defaultTTL,
	}
}

// Get :nodoc:
func (c *cacher) Get(key string) (cache interface{}, err error) {
	client := c.redis.Get()

	value, err := client.Do("GET", key)
	if err != nil {
		return nil, err
	}

	return value, nil
}

// Store :nodoc:
func (c *cacher) Store(item Item) error {
	client := c.redis.Get()

	_, err := client.Do("SETEX", item.GetKey(), c.cacheTTL(item), item.GetValue())
	if err != nil {
		return err
	}

	return nil
}

func (c *cacher) cacheTTL(item Item) (ttl int64) {
	if ttl = item.GetTTLInt64(); ttl > 0 {
		return
	}

	return int64(c.defaultTTL.Seconds())
}

func (c *cacher) SetRedisConnectionPool(pool *redis.Pool) {
	c.redis = pool
}
