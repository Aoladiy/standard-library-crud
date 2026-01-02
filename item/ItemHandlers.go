package item

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	rawId := r.PathValue("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		return
	}
	if id >= len(items) {
		http.Error(w, fmt.Sprintf("No item with id %v", id), http.StatusBadRequest)
		return
	}
	beautifulString, err := json.MarshalIndent(items[id], "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write(beautifulString)
	if err != nil {
		log.Println(err)
		return
	}
}

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	beautifulItems := map[uint]Item{}
	for i, item := range items {
		beautifulItems[uint(i)] = item
	}
	beautifulString, err := json.MarshalIndent(beautifulItems, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write(beautifulString)
	if err != nil {
		log.Println(err)
		return
	}
}

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var item Item
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&item)
	if err != nil {
		http.Error(w, "Cannot decode json", http.StatusUnprocessableEntity)
		return
	}
	if decoder.More() {
		http.Error(w, "Bad body of request", http.StatusBadRequest)
		return
	}
	items = append(items, item)
	_, err = w.Write([]byte("Successfully created Item with ID: " + strconv.Itoa(len(items)-1)))
	if err != nil {
		log.Println(err)
		return
	}
}

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	rawId := r.PathValue("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		return
	}
	if id >= len(items) {
		http.Error(w, fmt.Sprintf("No item with id %v", id), http.StatusBadRequest)
		return
	}
	oldItem := &items[id]
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newItem)
	if err != nil {
		http.Error(w, "Cannot decode json", http.StatusUnprocessableEntity)
		return
	}
	if decoder.More() {
		http.Error(w, "Bad body of request", http.StatusBadRequest)
		return
	}
	if newItem.Message != nil {
		oldItem.Message = newItem.Message
	}
	_, err = w.Write([]byte("Successfully updated Item with ID: " + rawId))
	if err != nil {
		log.Println(err)
		return
	}
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	rawId := r.PathValue("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		return
	}
	if id >= len(items) {
		http.Error(w, fmt.Sprintf("No item with id %v", id), http.StatusBadRequest)
		return
	}
	items = append(items[:id], items[id+1:]...)
	_, err = w.Write([]byte("Successfully deleted Item with ID: " + rawId))
	if err != nil {
		log.Println(err)
		return
	}
}
