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

	"github.com/Aoladiy/standard-library-crud/internal/db"
	"github.com/Aoladiy/standard-library-crud/internal/envLoading"
)

func setupServer(v envLoading.EnvVariables, handler http.Handler) {
	server := &http.Server{
		Addr:              v.Addr,
		Handler:           handler,
		ReadTimeout:       v.ReadTimeout,
		ReadHeaderTimeout: v.ReadHeaderTimeout,
		WriteTimeout:      v.WriteTimeout,
		IdleTimeout:       v.IdleTimeout,
	}

	log.Printf("server listening %v started\n", v.Addr)
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println(err)
			return
		}
	}()
	err := gracefulShutdown(server)
	if err != nil {
		return
	}
}

func gracefulShutdown(server *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Println("Error while shutdown", err)
		return err
	}
	err = db.DB.Close()
	if err != nil {
		log.Println("Error while closing db", err)
		return err
	}
	log.Println("Graceful shutdown complete")
	return nil
}
