package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	s Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s: s}
}

func (h Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	rawId := r.PathValue("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		return
	}
	user, err := h.s.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	beautifulString, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(beautifulString)
	if err != nil {
		log.Println(err)
		return
	}
}

func (h Handler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.s.GetUsers()
	if err != nil {
		http.Error(w, "cannot read all users", http.StatusInternalServerError)
		return
	}
	beautifulString, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(beautifulString)
	if err != nil {
		log.Println(err)
		return
	}
}

func (h Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
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
	id, err := h.s.CreateUser(user)
	if err != nil {
		http.Error(w, "cannot create user", http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte("Successfully created User with ID: " + strconv.Itoa(id)))
	if err != nil {
		log.Println(err)
		return
	}
}

func (h Handler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
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
	newUser.Id = id
	err = validateUser(newUser)
	if err != nil {
		http.Error(w, "Validation failed"+err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err, ok := h.s.UpdateUser(newUser)
	if err != nil && !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		_, err = w.Write([]byte("Nothing updated, most likely there is just no rows with such id: " + rawId))
	} else {
		_, err = w.Write([]byte("Successfully updated User with ID: " + rawId))
	}
	if err != nil {
		log.Println(err)
		return
	}
}

func (h Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	rawId := r.PathValue("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		return
	}
	err, ok := h.s.DeleteUserById(id)
	if err != nil && !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		_, err = w.Write([]byte("Nothing deleted, most likely there is just no rows with such id: " + rawId))
	} else {
		_, err = w.Write([]byte("Successfully deleted User with ID: " + rawId))
	}
	if err != nil {
		log.Println(err)
		return
	}
}
