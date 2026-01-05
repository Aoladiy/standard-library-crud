package main

import (
	"log"
	"net/http"

	"github.com/Aoladiy/standard-library-crud/internal/envLoading"
)

func setupServer(handler http.Handler) {
	server := &http.Server{
		Addr:                         envLoading.EnvVars.Addr,
		Handler:                      handler,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  envLoading.EnvVars.ReadTimeout,
		ReadHeaderTimeout:            envLoading.EnvVars.ReadHeaderTimeout,
		WriteTimeout:                 envLoading.EnvVars.WriteTimeout,
		IdleTimeout:                  envLoading.EnvVars.IdleTimeout,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
		HTTP2:                        nil,
		Protocols:                    nil,
	}

	log.Printf("server listening %v started\n", envLoading.EnvVars.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
		return
	}
}
