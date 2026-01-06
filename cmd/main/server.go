package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Aoladiy/standard-library-crud/internal/envLoading"
)

func setupServer(handler http.Handler) {
	server := &http.Server{
		Addr:              envLoading.EnvVars.Addr,
		Handler:           handler,
		ReadTimeout:       envLoading.EnvVars.ReadTimeout,
		ReadHeaderTimeout: envLoading.EnvVars.ReadHeaderTimeout,
		WriteTimeout:      envLoading.EnvVars.WriteTimeout,
		IdleTimeout:       envLoading.EnvVars.IdleTimeout,
	}

	log.Printf("server listening %v started\n", envLoading.EnvVars.Addr)
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println(err)
			return
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Println("Error while shutdown", err)
		return
	}
	log.Println("Graceful shutdown complete")
}
