package drivers

import (
	"github.com/go-redis/redis"
	"os"
)

var (
	RedisClient *redis.Client
	addr        = os.Getenv("REDIS_ADDR")
)

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})
}
