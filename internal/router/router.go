package router

import (
	"net/http"

	"github.com/Aoladiy/standard-library-crud/internal/envLoading"
	"github.com/Aoladiy/standard-library-crud/internal/user"
)

var envVars envLoading.EnvVariables

func SetupRouter(v envLoading.EnvVariables, h user.Handler) http.Handler {
	envVars = v
	router := http.NewServeMux()
	router.HandleFunc("GET /user/{id}", h.GetUserHandler)
	router.HandleFunc("GET /user", h.GetUsersHandler)
	router.HandleFunc("POST /user", h.CreateUserHandler)
	router.HandleFunc("PUT /user/{id}", h.UpdateUserHandler)
	router.HandleFunc("DELETE /user/{id}", h.DeleteUserHandler)
	return ChainOfMiddleware(
		router,
		RequestIdMiddleware,
		LoggerMiddleware,
		BasicAuthMiddleware,
		TimeoutMiddleware,
	)
}
