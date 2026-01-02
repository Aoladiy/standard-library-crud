package main

import (
	"standard-library-crud/envLoading"
	"standard-library-crud/router"
)

func main() {
	envVars := envLoading.LoadEnvVariables()
	handler := router.SetupRouter()
	setupServer(handler, envVars)
}
