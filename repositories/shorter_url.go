package repositories

import (
	"github.com/duel80003/my-url-shorteneder/drivers"
	"github.com/duel80003/my-url-shorteneder/entities"
	"github.com/duel80003/my-url-shorteneder/tools"
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

func (s *ShorterURL) Insert(data *entities.SetBody) (err error) {
	err = s.client.Set(data.Key, data.Value, getExpirationTime()).Err()
	return
}

func (s *ShorterURL) BatchSet(data []*entities.SetBody) (err error) {
	_, err = s.client.TxPipelined(func(pipeliner redis.Pipeliner) error {
		for i := range data {
			pipeliner.Set(data[i].Key, data[i].Value, getExpirationTime())
		}
		return nil
	})
	return
}

func (s *ShorterURL) Get(key interface{}) (url string, err error) {
	var k string
	switch key.(type) {
	case []byte:
		k = string(key.([]byte))
	default:
		k = key.(string)
	}
	url, err = s.client.Get(k).Result()
	if err != nil {
		logger.Errorf("[ShorterURL] Get key error: %s", err)
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
