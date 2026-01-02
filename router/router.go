package router

import (
	"net/http"
	"standard-library-crud/item"
)

func SetupRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /item/{id}", item.GetItemHandler)
	router.HandleFunc("GET /item", item.GetItemsHandler)
	router.HandleFunc("POST /item", item.CreateItemHandler)
	router.HandleFunc("PATCH /item/{id}", item.UpdateItemHandler)
	router.HandleFunc("DELETE /item/{id}", item.DeleteItemHandler)
	return router
}
