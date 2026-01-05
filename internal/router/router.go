package router

import (
	"net/http"

	"github.com/Aoladiy/standard-library-crud/internal/user"
)

func SetupRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /item/{id}", user.GetUserHandler)
	router.HandleFunc("GET /item", user.GetUsersHandler)
	router.HandleFunc("POST /item", user.CreateUserHandler)
	router.HandleFunc("PATCH /item/{id}", user.UpdateUserHandler)
	router.Handle("DELETE /item/{id}", LoggerMiddleware(http.HandlerFunc(user.DeleteUserHandler)))
	return ChainOfMiddleware(
		router,
		RequestIdMiddleware,
		LoggerMiddleware,
		TimeoutMiddleware,
	)
}
