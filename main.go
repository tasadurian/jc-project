package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/thetommytwitch/jc-project/server"
)

const port = ":8080"

func main() {
	sigs := make(chan os.Signal, 1)

	hs, logger := setup(sigs)

	go func() {
		logger.Printf("Listening on port %s\n", hs.Addr)

		if err := hs.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()

	shutdown(hs, logger, 5*time.Second, sigs)
}

func setup(sigs chan os.Signal) (*http.Server, *log.Logger) {
	logger := log.New(os.Stdout, "", 0)

	return &http.Server{
		Addr:    port,
		Handler: server.New(sigs),
	}, logger
}

func shutdown(hs *http.Server, logger *log.Logger, timeout time.Duration, sigs chan os.Signal) {
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	<-sigs

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	logger.Printf("\nShutdown with timeout: %s\n", timeout)

	if err := hs.Shutdown(ctx); err != nil {
		logger.Printf("Error: %v\n", err)
	} else {
		logger.Println("Server stopped")
	}
}
