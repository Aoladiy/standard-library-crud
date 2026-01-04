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
	if id >= len(users) {
		http.Error(w, fmt.Sprintf("No item with id %v", id), http.StatusBadRequest)
		return
	}
	beautifulString, err := json.MarshalIndent(users[id], "", "\t")
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
	beautifulItems := map[uint]User{}
	for i, item := range users {
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
	var item User
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
	users = append(users, item)
	_, err = w.Write([]byte("Successfully created User with ID: " + strconv.Itoa(len(users)-1)))
	if err != nil {
		log.Println(err)
		return
	}
}

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItem User
	rawId := r.PathValue("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		return
	}
	if id >= len(users) {
		http.Error(w, fmt.Sprintf("No item with id %v", id), http.StatusBadRequest)
		return
	}
	oldItem := &users[id]
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
	oldItem.Email = newItem.Email
	if newItem.FullName != nil {
		oldItem.FullName = newItem.FullName
	}
	if newItem.PhoneNumber != nil {
		oldItem.PhoneNumber = newItem.PhoneNumber
	}
	_, err = w.Write([]byte("Successfully updated User with ID: " + rawId))
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
	if id >= len(users) {
		http.Error(w, fmt.Sprintf("No item with id %v", id), http.StatusBadRequest)
		return
	}
	users = append(users[:id], users[id+1:]...)
	_, err = w.Write([]byte("Successfully deleted User with ID: " + rawId))
	if err != nil {
		log.Println(err)
		return
	}
}
