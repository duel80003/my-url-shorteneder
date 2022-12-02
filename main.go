package main

import (
	"context"
	"fmt"
	"github.com/duel80003/my-url-shorteneder/tools"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/duel80003/my-url-shorteneder/drivers"
	"github.com/duel80003/my-url-shorteneder/router"
)

var logger = tools.Logger

func main() {
	r := router.HttpRouters()
	// Add your routes as needed
	host := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	srv := &http.Server{
		Addr: host,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Println(err)
		}
	}()

	logger.Infof("server run on %s", srv.Addr)
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
