package router

import (
	"github.com/duel80003/my-url-shorter/deliveries"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var (
	d = deliveries.NewShorterURLDelivery()
)

var env = os.Getenv("ENV")

// CORSHeadersMiddleware add common response header
func CORSHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if env == "Local" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		}
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// OptionsMiddleware handle options request
func OptionsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func HttpRouters() *mux.Router {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(OptionsMiddleware)
	r.Use(CORSHeadersMiddleware)

	r.HandleFunc("/encode", d.Encode).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/{id}", d.Redirect).Methods(http.MethodGet)
	//info.Use(middlewes.AuthServerInfoMiddleware)
	return r
}
