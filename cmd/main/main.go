package main

import (
	"github.com/Aoladiy/standard-library-crud/internal/db"
	"github.com/Aoladiy/standard-library-crud/internal/envLoading"
	"github.com/Aoladiy/standard-library-crud/internal/router"
)

func main() {
	envLoading.LoadEnvVariables()
	db.InitDB()
	handler := router.SetupRouter()
	setupServer(handler)
}
