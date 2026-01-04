package main

import (
	"github.com/Aoladiy/standard-library-crud/envLoading"
	"github.com/Aoladiy/standard-library-crud/router"
)

func main() {
	envVars := envLoading.LoadEnvVariables()
	handler := router.SetupRouter()
	setupServer(handler, envVars)
}
