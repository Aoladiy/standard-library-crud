package main

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
	service := user.NewService(*repo)
	handler := user.NewHandler(*service)
	setupRouter := router.SetupRouter(envVars, *handler)
	setupServer(envVars, setupRouter)
}
