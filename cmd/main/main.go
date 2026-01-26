package main

// @title Standard Library CRUD API
// @version 1.0
// @description CRUD API for managing users.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.basic BasicAuth

import (
	"github.com/Aoladiy/standard-library-crud/internal/db"
	"github.com/Aoladiy/standard-library-crud/internal/envLoading"
	"github.com/Aoladiy/standard-library-crud/internal/router"
	"github.com/Aoladiy/standard-library-crud/internal/user"
)

func main() {
	envVars := envLoading.LoadEnvVariables()
	database := db.InitDB(envVars)
	repo := user.NewRepo(database)
	service := user.NewService(repo)
	handler := user.NewHandler(service)
	setupRouter := router.SetupRouter(envVars, handler)
	setupServer(envVars, setupRouter, database)
}
