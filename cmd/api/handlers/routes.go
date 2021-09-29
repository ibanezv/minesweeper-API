package handlers

import (
	"github/ibanezv/minesweeper-API/internal/accounts"
	"github/ibanezv/minesweeper-API/internal/distributions"
	"github/ibanezv/minesweeper-API/internal/games"
	"github/ibanezv/minesweeper-API/internal/users"
	"net/http"

	"github.com/gorilla/mux"
)

func ApiRoutesMapper(gamesService games.Service, distributionsService distributions.Service,
	usersService users.Service, accountsService accounts.Service) {
	r := mux.NewRouter()
	gamesHandler := NewHandlerGames(gamesService)
	distributionsHandler := NewHandlerDistributions(distributionsService)
	userHandler := NewHandlerUsers(usersService)
	accountHandler := NewHandlerAccounts(accountsService)
	r.HandleFunc("/games/{id}", gamesHandler.GetGame).Methods(http.MethodGet)
	r.HandleFunc("/games", gamesHandler.PostGame).Methods(http.MethodPost)
	r.HandleFunc("/games/{id}/distributions", distributionsHandler.GetDistribution).Methods(http.MethodGet)
	r.HandleFunc("/games/{id}/distributions", distributionsHandler.PatchDistribution).Methods(http.MethodPatch)
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/users", userHandler.PostUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}/games", userHandler.GetUserGames).Methods(http.MethodGet)
	r.HandleFunc("/accounts/{id}", accountHandler.GetAccount).Methods(http.MethodGet)
	r.HandleFunc("/accounts", accountHandler.PostAccount).Methods(http.MethodPost)
	r.HandleFunc("/accounts/{id}/users", accountHandler.GetAccountsUser).Methods(http.MethodGet)
	http.Handle("/", r)
	_ = http.ListenAndServe(":8090", nil)
}
