package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type envVariables struct {
	Addr              string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

type Item struct {
	Message *string `json:"message,omitempty"`
}

var items []Item

func main() {
	envVars := loadEnvVariables()
	router := setupRouter()
	setupServer(router, envVars)
}

func getItemHandler(w http.ResponseWriter, r *http.Request) {
	rawId := r.PathValue("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		return
	}
	if id >= len(items) {
		http.Error(w, fmt.Sprintf("No item with id %v", id), http.StatusBadRequest)
		return
	}
	beautifulString, err := json.MarshalIndent(items[id], "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write(beautifulString)
	if err != nil {
		log.Println(err)
		return
	}
}

func getItemsHandler(w http.ResponseWriter, r *http.Request) {
	beautifulItems := map[uint]Item{}
	for i, item := range items {
		beautifulItems[uint(i)] = item
	}
	beautifulString, err := json.MarshalIndent(beautifulItems, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write(beautifulString)
	if err != nil {
		log.Println(err)
		return
	}
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	var item Item
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&item)
	if err != nil {
		http.Error(w, "Cannot decode json", http.StatusUnprocessableEntity)
		return
	}
	if decoder.More() {
		http.Error(w, "Bad body of request", http.StatusBadRequest)
		return
	}
	items = append(items, item)
	_, err = w.Write([]byte("Successfully created Item with ID: " + strconv.Itoa(len(items)-1)))
	if err != nil {
		log.Println(err)
		return
	}
}

func updateItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	rawId := r.PathValue("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		return
	}
	if id >= len(items) {
		http.Error(w, fmt.Sprintf("No item with id %v", id), http.StatusBadRequest)
		return
	}
	oldItem := &items[id]
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newItem)
	if err != nil {
		http.Error(w, "Cannot decode json", http.StatusUnprocessableEntity)
		return
	}
	if decoder.More() {
		http.Error(w, "Bad body of request", http.StatusBadRequest)
		return
	}
	if newItem.Message != nil {
		oldItem.Message = newItem.Message
	}
	_, err = w.Write([]byte("Successfully updated Item with ID: " + rawId))
	if err != nil {
		log.Println(err)
		return
	}
}

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	rawId := r.PathValue("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		return
	}
	if id >= len(items) {
		http.Error(w, fmt.Sprintf("No item with id %v", id), http.StatusBadRequest)
		return
	}
	items = append(items[:id], items[id+1:]...)
	_, err = w.Write([]byte("Successfully deleted Item with ID: " + rawId))
	if err != nil {
		log.Println(err)
		return
	}
}

func loadEnvVariables() envVariables {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
		return envVariables{}
	}
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Println("applying default value for server port")
		serverPort = "8080"
	}
	serverReadTimeout, err := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	if err != nil {
		log.Println(err)
		serverReadTimeout = 10
	}
	serverReadHeaderTimeout, err := strconv.Atoi(os.Getenv("SERVER_READ_HEADER_TIMEOUT"))
	if err != nil {
		log.Println(err)
		serverReadHeaderTimeout = 5
	}
	serverWriteTimeout, err := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
	if err != nil {
		log.Println(err)
		serverWriteTimeout = 5
	}
	serverIdleTimeout, err := strconv.Atoi(os.Getenv("SERVER_IDLE_TIMEOUT"))
	if err != nil {
		log.Println(err)
		serverIdleTimeout = 60
	}
	log.Println("successfully loaded env variables")
	return envVariables{
		Addr:              os.Getenv("APP_HOST") + ":" + serverPort,
		ReadTimeout:       time.Second * time.Duration(serverReadTimeout),
		ReadHeaderTimeout: time.Second * time.Duration(serverReadHeaderTimeout),
		WriteTimeout:      time.Second * time.Duration(serverWriteTimeout),
		IdleTimeout:       time.Second * time.Duration(serverIdleTimeout),
	}
}

func setupRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /item/{id}", getItemHandler)
	router.HandleFunc("GET /item", getItemsHandler)
	router.HandleFunc("POST /item", createItemHandler)
	router.HandleFunc("PATCH /item/{id}", updateItemHandler)
	router.HandleFunc("DELETE /item/{id}", deleteItemHandler)
	return router
}

func setupServer(router http.Handler, envVars envVariables) {
	server := &http.Server{
		Addr:                         envVars.Addr,
		Handler:                      router,
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
