package router

import (
	"net/http"

	"github.com/Aoladiy/standard-library-crud/item"
)

func SetupRouter() http.Handler {
	router := http.NewServeMux()
	router.Handle("GET /item/{id}", SecondLoggerMiddleware(LoggerMiddleware(http.HandlerFunc(item.GetItemHandler))))
	router.HandleFunc("GET /item", item.GetItemsHandler)
	router.HandleFunc("POST /item", item.CreateItemHandler)
	router.HandleFunc("PATCH /item/{id}", item.UpdateItemHandler)
	router.HandleFunc("DELETE /item/{id}", item.DeleteItemHandler)
	return RegisterMiddleware(
		router,
		LoggerMiddleware,
		SecondLoggerMiddleware,
	)
}
