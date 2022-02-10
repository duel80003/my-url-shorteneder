package usecases

import (
	"fmt"
	"github.com/duel80003/my-url-shorter/entities"
	"github.com/duel80003/my-url-shorter/repositories"
	"github.com/duel80003/my-url-shorter/tools"
	"github.com/jxskiss/base62"
	"github.com/teris-io/shortid"
	"os"
)

var (
	logger = tools.Logger
	domain = os.Getenv("DOMAIN")
)

func NewShorterURLCase() *ShorterURLCase {
	return &ShorterURLCase{
		shorterURLRepository: repositories.NewShorterURL(),
	}
}

type ShorterURLCase struct {
	shorterURLRepository *repositories.ShorterURL
}

func (s *ShorterURLCase) GenerateShortURL(url string) (shortedURL string, err error) {
	logger.Info("[ShorterURLCase] GenerateShortURL")
	id, _ := shortid.Generate()
	ID := base62.Encode([]byte(id))
	shortedURL = fmt.Sprintf("%s/%s", domain, string(ID))
	data := &entities.ShorterURL{
		ID:  string(ID),
		URL: url,
	}
	err = s.shorterURLRepository.Insert(data)
	return
}

func (s *ShorterURLCase) GetOriginalURL(id string) (url string, err error) {
	logger.Info("[ShorterURLCase] GetOriginalURL")
	url, err = s.shorterURLRepository.Get(id)
	return
}
