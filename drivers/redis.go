package drivers

import (
	"github.com/duel80003/my-url-shorteneder/tools"
	"github.com/go-redis/redis"
	"os"
)

var (
	RedisClient *redis.Client
	addr        = os.Getenv("REDIS_ADDR")
	logger      = tools.Logger
)

func init() {
	logger.Debugf("redis addr %s", addr)
	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})
}
