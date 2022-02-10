package repositories

import (
	"github.com/duel80003/my-url-shorter/drivers"
	"github.com/duel80003/my-url-shorter/entities"
	"github.com/duel80003/my-url-shorter/tools"
	"github.com/go-redis/redis"
	"os"
	"strconv"
	"sync"
	"time"
)

const defaultExpiration = 3

var (
	expiration = os.Getenv("EXPIRATION")
	t          time.Duration
	once       sync.Once
	logger     = tools.Logger
)

type ShorterURL struct {
	client *redis.Client
}

func NewShorterURL() *ShorterURL {
	return &ShorterURL{
		client: drivers.RedisClient,
	}
}

func (s *ShorterURL) Insert(data *entities.ShorterURL) (err error) {
	err = s.client.Set(data.ID, data.URL, getExpirationTime()).Err()
	return
}

func (s *ShorterURL) Get(ID string) (url string, err error) {
	url, err = s.client.Get(ID).Result()
	if err != nil {
		logger.Error("cccccccc")
	}
	return
}

func getExpirationTime() time.Duration {
	once.Do(func() {
		value, err := strconv.Atoi(expiration)
		if err != nil {
			value = defaultExpiration
		}
		// expiration time unit: day
		t = time.Duration(value) * time.Hour * 24
	})
	return t
}
