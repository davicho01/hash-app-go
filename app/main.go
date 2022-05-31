package main

import (
	"app/handlers"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"
)

func createChannel() (chan os.Signal, func()) {
	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)

	return stopChannel, func() {
		close(stopChannel)
	}
}

func start(server *http.Server, logger *log.Logger) {
	logger.Println("application Started")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	} else {
		log.Println("Application Stopped Gracefully")
	}
}

func shutdown(ctx context.Context, server *http.Server, logger *log.Logger) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	} else {
		logger.Println("Application Shutting Down...")
	}
}

func main() {

	logger := log.New(os.Stdout, "Main App: ", log.LstdFlags)

	hashHandler := handlers.NewHash()
	shutdownHandler := handlers.NewShutdown()

	serveMux := http.NewServeMux()
	serveMux.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {

			var stats = regexp.MustCompile(`/stats`)
			var hash = regexp.MustCompile(`/hash`)
			var hashId = regexp.MustCompile(`/hash/[0-9]`)
			var shutdown = regexp.MustCompile(`/shutdown`)

			switch {
			case stats.MatchString(r.URL.Path):
				hashHandler.GetHashStats(w, r)
			case hashId.MatchString(r.URL.Path):
				hashHandler.GetHashByIdentifier(w, r)
			case hash.MatchString(r.URL.Path):
				hashHandler.GenerateHash(w, r)
			case shutdown.MatchString(r.URL.Path):
				shutdownHandler.Shutdown(w, r)
			default:
				w.Write([]byte("Unknown Pattern"))
			}
		},
	)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go start(server, logger)

	stopCh, closeCh := createChannel()
	defer closeCh()
	logger.Println("notified:", <-stopCh)

	shutdown(context.Background(), server, logger)

}
