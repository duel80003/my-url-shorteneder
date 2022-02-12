package usecases

import (
	"fmt"
	"github.com/jxskiss/base62"
	"os"

	"github.com/duel80003/my-url-shorteneder/entities"
	"github.com/duel80003/my-url-shorteneder/repositories"
	"github.com/duel80003/my-url-shorteneder/tools"
	"github.com/teris-io/shortid"
)

var (
	logger = tools.Logger
	domain = os.Getenv("DOMAIN")
	sid, _ = shortid.New(1, shortid.DefaultABC, 2342)
)

func NewShorterURLCase() *ShorterURLCase {
	return &ShorterURLCase{
		shorterURLRepository: repositories.NewShorterURL(),
	}
}

type ShorterURLCase struct {
	shorterURLRepository *repositories.ShorterURL
}

// GenerateShortURL create short url and insert into redis
func (s *ShorterURLCase) GenerateShortURL(url string) (shortedURL string, err error) {
	logger.Info("[ShorterURLCase] GenerateShortURL")
	id, _ := sid.Generate()
	encode := base62.Encode([]byte(url))

	// check the input url is duplicate
	shortedURL, err = s.shorterURLRepository.Get(encode)
	if err == nil {
		return
	}
	encodeID := base62.Encode([]byte(id))
	shortedURL = fmt.Sprintf("%s/%s", domain, encodeID)

	body1 := entities.NewSetBody(string(encodeID), url)
	body2 := entities.NewSetBody(string(encode), shortedURL)
	data := []*entities.SetBody{body1, body2}
	err = s.shorterURLRepository.BatchSet(data)
	return
}

// GetOriginalURL return original url by id
func (s *ShorterURLCase) GetOriginalURL(id string) (url string, err error) {
	logger.Info("[ShorterURLCase] GetOriginalURL")
	url, err = s.shorterURLRepository.Get(id)
	return
}
