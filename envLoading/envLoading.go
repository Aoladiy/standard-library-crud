package envLoading

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	Addr              string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

func LoadEnvVariables() EnvVariables {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
		return EnvVariables{}
	}
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Println("applying default value for server port")
		serverPort = "8080"
	}
	serverReadTimeout, err := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	if err != nil {
		log.Println("error loading SERVER_READ_TIMEOUT env variable: " + err.Error())
		serverReadTimeout = 10
	}
	serverReadHeaderTimeout, err := strconv.Atoi(os.Getenv("SERVER_READ_HEADER_TIMEOUT"))
	if err != nil {
		log.Println("error loading SERVER_READ_HEADER_TIMEOUT env variable: " + err.Error())
		serverReadHeaderTimeout = 5
	}
	serverWriteTimeout, err := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
	if err != nil {
		log.Println("error loading SERVER_WRITE_TIMEOUT env variable: " + err.Error())
		serverWriteTimeout = 5
	}
	serverIdleTimeout, err := strconv.Atoi(os.Getenv("SERVER_IDLE_TIMEOUT"))
	if err != nil {
		log.Println("error loading SERVER_IDLE_TIMEOUT env variable: " + err.Error())
		serverIdleTimeout = 60
	}
	log.Println("successfully loaded env variables")
	return EnvVariables{
		Addr:              os.Getenv("APP_HOST") + ":" + serverPort,
		ReadTimeout:       time.Second * time.Duration(serverReadTimeout),
		ReadHeaderTimeout: time.Second * time.Duration(serverReadHeaderTimeout),
		WriteTimeout:      time.Second * time.Duration(serverWriteTimeout),
		IdleTimeout:       time.Second * time.Duration(serverIdleTimeout),
	}
}
