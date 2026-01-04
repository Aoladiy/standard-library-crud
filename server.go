package main

import (
	"log"
	"net/http"

	"github.com/Aoladiy/standard-library-crud/envLoading"
)

func setupServer(handler http.Handler, envVars envLoading.EnvVariables) {
	server := &http.Server{
		Addr:                         envVars.Addr,
		Handler:                      handler,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  envVars.ReadTimeout,
		ReadHeaderTimeout:            envVars.ReadHeaderTimeout,
		WriteTimeout:                 envVars.WriteTimeout,
		IdleTimeout:                  envVars.IdleTimeout,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
		HTTP2:                        nil,
		Protocols:                    nil,
	}

	log.Printf("server listening %v started\n", envVars.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
		return
	}
}
