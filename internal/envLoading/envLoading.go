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
	LoadedUsername    string
	LoadedPassword    string
	DBHost            string
	DBUser            string
	DBPass            string
}

type DbParams struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

func (d *DbParams) GetDsn() string {
	return "postgres://" + d.User + ":" + d.Password + "@" + d.Host + ":" + d.Port + "/" + d.DbName
}

var EnvVars EnvVariables
var DbConnParams DbParams

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
	loadedUsername := os.Getenv("ADMIN_USERNAME")
	if loadedUsername == "" {
		log.Println("NO ADMIN_USERNAME SET")
	}
	loadedPassword := os.Getenv("ADMIN_PASSWORD")
	if loadedPassword == "" {
		log.Println("NO ADMIN_PASSWORD SET")
	}

	loadedDbUser := os.Getenv("DB_USER")
	if loadedDbUser == "" {
		log.Println("NO DB_USER SET")
	}
	loadedDbPassword := os.Getenv("DB_PASSWORD")
	if loadedDbPassword == "" {
		log.Println("NO DB_PASSWORD SET")
	}
	loadedDbHost := os.Getenv("DB_HOST")
	if loadedDbHost == "" {
		log.Println("NO DB_HOST SET")
	}
	loadedDbPort := os.Getenv("DB_PORT")
	if loadedDbPort == "" {
		log.Println("NO DB_PORT SET")
	}
	loadedDbName := os.Getenv("DB_NAME")
	if loadedDbName == "" {
		log.Println("NO DB_NAME SET")
	}

	log.Println("successfully loaded env variables")

	DbConnParams = DbParams{
		User:     loadedDbUser,
		Password: loadedDbPassword,
		Host:     loadedDbHost,
		Port:     loadedDbPort,
		DbName:   loadedDbName,
	}

	EnvVars = EnvVariables{
		Addr:              os.Getenv("APP_HOST") + ":" + serverPort,
		ReadTimeout:       time.Second * time.Duration(serverReadTimeout),
		ReadHeaderTimeout: time.Second * time.Duration(serverReadHeaderTimeout),
		WriteTimeout:      time.Second * time.Duration(serverWriteTimeout),
		IdleTimeout:       time.Second * time.Duration(serverIdleTimeout),
		LoadedUsername:    loadedUsername,
		LoadedPassword:    loadedPassword,
	}
	return EnvVars
}
