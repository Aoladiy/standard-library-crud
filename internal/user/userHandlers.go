package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	s *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{s: s}
}

// GetUserHandler godoc
// @Summary Get a user by ID
// @Description Retrieves a single user by their ID.
// @Tags users
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {object} User
// @Failure 400 {string} string "Bad request"
// @Router /user/{id} [get]
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

// GetUsersHandler godoc
// @Summary List users
// @Description Retrieves all users.
// @Tags users
// @Produce json
// @Success 200 {array} User
// @Failure 500 {string} string "Internal server error"
// @Router /user [get]
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

// CreateUserHandler godoc
// @Summary Create a user
// @Description Creates a new user.
// @Tags users
// @Accept json
// @Produce plain
// @Param user body User true "User payload"
// @Success 200 {string} string "Successfully created User with ID: {id}"
// @Failure 422 {string} string "Unprocessable entity"
// @Failure 500 {string} string "Internal server error"
// @Router /user [post]
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

// UpdateUserHandler godoc
// @Summary Update a user
// @Description Updates an existing user by ID.
// @Tags users
// @Accept json
// @Produce plain
// @Param id path int true "User ID"
// @Param user body User true "User payload"
// @Success 200 {string} string "Successfully updated User with ID: {id}"
// @Failure 400 {string} string "Bad request"
// @Failure 422 {string} string "Unprocessable entity"
// @Router /user/{id} [put]
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
	err = h.s.UpdateUser(newUser)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Nothing updated, most likely there is just no rows with such id: "+rawId, http.StatusBadRequest)
		return
	}
	_, err = w.Write([]byte("Successfully updated User with ID: " + rawId))
	if err != nil {
		log.Println(err)
		return
	}
}

// DeleteUserHandler godoc
// @Summary Delete a user
// @Description Deletes a user by ID.
// @Tags users
// @Param id path int true "User ID"
// @Produce plain
// @Success 200 {string} string "Successfully deleted User with ID: {id}"
// @Failure 400 {string} string "Bad request"
// @Router /user/{id} [delete]
func (h Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	rawId := r.PathValue("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println(err)
		return
	}
	err = h.s.DeleteUserById(id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Nothing deleted, most likely there is just no rows with such id: "+rawId, http.StatusBadRequest)
		return
	}
	_, err = w.Write([]byte("Successfully deleted User with ID: " + rawId))
	if err != nil {
		log.Println(err)
		return
	}
}
