package user

import (
	"encoding/json"
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
	user, err := Users.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	beautifulString, err := json.MarshalIndent(user, "", "\t")
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
	for i, user := range Users.GetUsers() {
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
	decoder.DisallowUnknownFields()
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
	Users.AddUser(user)
	_, err = w.Write([]byte("Successfully created User with ID: " + strconv.Itoa(len(Users.GetUsers())-1)))
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
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Cannot decode json", http.StatusUnprocessableEntity)
		return
	}
	if decoder.More() {
		http.Error(w, "Bad body of request", http.StatusBadRequest)
		return
	}
	err = validateUser(newUser)
	if err != nil {
		http.Error(w, "Validation failed"+err.Error(), http.StatusUnprocessableEntity)
		return
	}
	_, err = Users.UpdateUser(id, newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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
	err = Users.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = w.Write([]byte("Successfully deleted User with ID: " + rawId))
	if err != nil {
		log.Println(err)
		return
	}
}
