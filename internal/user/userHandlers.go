package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	rawId := r.PathValue("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		return
	}
	if id >= len(users) {
		http.Error(w, fmt.Sprintf("No user with id %v", id), http.StatusBadRequest)
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

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	beautifulUsers := map[uint]User{}
	for i, user := range users {
		beautifulUsers[uint(i)] = user
	}
	beautifulString, err := json.MarshalIndent(beautifulUsers, "", "\t")
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

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Cannot decode json", http.StatusUnprocessableEntity)
		return
	}
	if decoder.More() {
		http.Error(w, "Bad body of request", http.StatusBadRequest)
		return
	}
	err = validateUser(user)
	if err != nil {
		http.Error(w, "Validation failed"+err.Error(), http.StatusUnprocessableEntity)
		return
	}
	users = append(users, user)
	_, err = w.Write([]byte("Successfully created User with ID: " + strconv.Itoa(len(users)-1)))
	if err != nil {
		log.Println(err)
		return
	}
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	rawId := r.PathValue("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		return
	}
	if id >= len(users) {
		http.Error(w, fmt.Sprintf("No user with id %v", id), http.StatusBadRequest)
		return
	}
	err = validateUser(newUser)
	if err != nil {
		http.Error(w, "Validation failed"+err.Error(), http.StatusUnprocessableEntity)
		return
	}
	oldUser := &users[id]
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Cannot decode json", http.StatusUnprocessableEntity)
		return
	}
	if decoder.More() {
		http.Error(w, "Bad body of request", http.StatusBadRequest)
		return
	}
	oldUser.Email = newUser.Email
	if newUser.FullName != nil {
		oldUser.FullName = newUser.FullName
	}
	if newUser.PhoneNumber != nil {
		oldUser.PhoneNumber = newUser.PhoneNumber
	}
	_, err = w.Write([]byte("Successfully updated User with ID: " + rawId))
	if err != nil {
		log.Println(err)
		return
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	rawId := r.PathValue("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		return
	}
	if id >= len(users) {
		http.Error(w, fmt.Sprintf("No user with id %v", id), http.StatusBadRequest)
		return
	}
	users = append(users[:id], users[id+1:]...)
	_, err = w.Write([]byte("Successfully deleted User with ID: " + rawId))
	if err != nil {
		log.Println(err)
		return
	}
}
