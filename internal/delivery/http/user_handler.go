package http

import (
	"encoding/json"
	"net/http"

	"github.com/santigorbe/clean_arq/internal/domain"
	"github.com/santigorbe/clean_arq/internal/usecase"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	usecase *usecase.UserUseCase
}

func NewUserHandler(r *mux.Router, uc *usecase.UserUseCase) {
	handler := &UserHandler{usecase: uc}
	r.HandleFunc("/users", handler.GetAll).Methods("GET")
	r.HandleFunc("/users/{id}", handler.GetByID).Methods("GET")
	r.HandleFunc("/users", handler.Create).Methods("POST")
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.usecase.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user, err := h.usecase.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.usecase.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
