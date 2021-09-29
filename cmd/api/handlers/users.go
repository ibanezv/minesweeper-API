package handlers

import (
	"encoding/json"
	"errors"
	"github/ibanezv/minesweeper-API/internal/models"
	"github/ibanezv/minesweeper-API/internal/users"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type UsersHandler struct {
	ServiceUsers users.Service
}

func NewHandlerUsers(serviceUsers users.Service) UsersHandler {
	return UsersHandler{serviceUsers}
}

func (h *UsersHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var strUserID = params["id"]
	userID, err := strconv.Atoi(strUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	user, err := h.ServiceUsers.FindUser(ctx, userID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(HeaderKey, HeaderValue)
	jsonResp, _ := json.Marshal(user)
	w.Write(jsonResp)
}

func (h *UsersHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserInput(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ctx := r.Context()
	newUser, err := h.ServiceUsers.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set(HeaderKey, HeaderValue)
	jsonResp, _ := json.Marshal(newUser)
	w.Write(jsonResp)
}

func (h *UsersHandler) GetUserGames(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var strUserID = params["id"]
	userID, err := strconv.Atoi(strUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	games, err := h.ServiceUsers.FindUserGames(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(HeaderKey, HeaderValue)
	jsonResp, _ := json.Marshal(games)
	w.Write(jsonResp)
}

func getUserInput(request *http.Request) (models.User, error) {
	var input models.User
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		return models.User{}, err
	}
	return input, nil
}
