package handlers

import (
	"encoding/json"
	"errors"
	"github/ibanezv/minesweeper-API/internal/accounts"
	"github/ibanezv/minesweeper-API/internal/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AccountsHandler struct {
	serviceAccounts accounts.Service
}

func NewHandlerAccounts(serviceAccounts accounts.Service) AccountsHandler {
	return AccountsHandler{serviceAccounts}
}

func (h *AccountsHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var strAccountID = params["id"]
	accountID, err := strconv.Atoi(strAccountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	accounts, err := h.serviceAccounts.FindAccount(ctx, accountID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(HeaderKey, HeaderValue)
	jsonResp, _ := json.Marshal(accounts)
	w.Write(jsonResp)
}

func (h *AccountsHandler) PostAccount(w http.ResponseWriter, r *http.Request) {
	account, err := getAccountInput(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ctx := r.Context()
	newAccount, err := h.serviceAccounts.CreateAccount(ctx, account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set(HeaderKey, HeaderValue)
	jsonResp, _ := json.Marshal(newAccount)
	w.Write(jsonResp)
}

func (h *AccountsHandler) GetAccountsUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var strAccountID = params["id"]
	accountID, err := strconv.Atoi(strAccountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	users, err := h.serviceAccounts.FindAccountUsers(ctx, accountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(HeaderKey, HeaderValue)
	jsonResp, _ := json.Marshal(users)
	w.Write(jsonResp)
}

func getAccountInput(request *http.Request) (models.Accounts, error) {
	var input models.Accounts
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		return models.Accounts{}, err
	}
	return input, nil
}
