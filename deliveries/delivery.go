package deliveries

import (
	"encoding/json"
	"github.com/duel80003/my-url-shorter/tools"
	"github.com/duel80003/my-url-shorter/usecases"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var logger = tools.Logger

func NewShorterURLDelivery() *ShorterURLDelivery {
	return &ShorterURLDelivery{
		useCase: usecases.NewShorterURLCase(),
	}
}

type ShorterURLDelivery struct {
	useCase *usecases.ShorterURLCase
}

func (s *ShorterURLDelivery) Encode(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]interface{})
	defer func() {
		bytes, _ := json.Marshal(m)
		_ = r.Body.Close()
		_, _ = w.Write(bytes)
	}()
	body, err := ioutil.ReadAll(r.Body)
	reqMap := make(map[string]interface{})
	err = json.Unmarshal(body, &reqMap)
	url, ok1 := reqMap["url"]
	str, ok2 := url.(string)
	if !ok1 || !ok2 {
		w.WriteHeader(http.StatusBadRequest)
		m["message"] = "Bad request"
		return
	}

	result, err := s.useCase.GenerateShortURL(str)
	if err != nil {
		logger.Errorf("[ShorterURLDelivery] GenerateShortURL error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		m["message"] = "Internal server error"
		return
	}

	m["message"] = "success"
	m["shorterUrl"] = result
}

func (s *ShorterURLDelivery) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	url, err := s.useCase.GetOriginalURL(id)
	if err != nil {
		logger.Errorf("[ShorterURLDelivery] GetOriginalURL error %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}
