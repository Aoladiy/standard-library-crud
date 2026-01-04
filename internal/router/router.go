package router

import (
	"net/http"

	"github.com/Aoladiy/standard-library-crud/internal/item"
)

func SetupRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /item/{id}", item.GetItemHandler)
	router.HandleFunc("GET /item", item.GetItemsHandler)
	router.HandleFunc("POST /item", item.CreateItemHandler)
	router.HandleFunc("PATCH /item/{id}", item.UpdateItemHandler)
	router.Handle("DELETE /item/{id}", SecondLoggerMiddleware(LoggerMiddleware(http.HandlerFunc(item.DeleteItemHandler))))
	return RegisterMiddleware(
		router,
		LoggerMiddleware,
	)
}
