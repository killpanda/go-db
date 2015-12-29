package store

import (
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

// RedisPool will read env and init a redis pool
func RedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   100,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			addr := os.Getenv("REDIS_ADDR")
			port := os.Getenv("REDIS_PORT")
			passwd := os.Getenv("REDIS_PASSWD")
			option := redis.DialPassword(passwd)

			c, err := redis.Dial("tcp", addr+":"+port, option)
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
