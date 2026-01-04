package main

import (
	"github.com/Aoladiy/standard-library-crud/internal/envLoading"
	"github.com/Aoladiy/standard-library-crud/internal/router"
)

func main() {
	envVars := envLoading.LoadEnvVariables()
	handler := router.SetupRouter()
	setupServer(handler, envVars)
}
