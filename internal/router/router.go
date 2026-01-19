package router

import (
	"net/http"

	"github.com/Aoladiy/standard-library-crud/internal/user"
)

func SetupRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /user/{id}", user.GetUserHandler)
	router.HandleFunc("GET /user", user.GetUsersHandler)
	router.HandleFunc("POST /user", user.CreateUserHandler)
	router.HandleFunc("PUT /user/{id}", user.UpdateUserHandler)
	router.Handle("DELETE /user/{id}", LoggerMiddleware(http.HandlerFunc(user.DeleteUserHandler)))
	return ChainOfMiddleware(
		router,
		RequestIdMiddleware,
		LoggerMiddleware,
		BasicAuthMiddleware,
		TimeoutMiddleware,
	)
}
