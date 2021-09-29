package handlers

import (
	"github/ibanezv/minesweeper-API/internal/distributions"
	"github/ibanezv/minesweeper-API/internal/games"
	"net/http"

	"github.com/gorilla/mux"
)

func ApiRoutesMapper(gamesService games.Service, distributionsService distributions.Service) {
	r := mux.NewRouter()
	gamesHandler := NewHandlerGames(gamesService)
	distributionsHandler := NewHandlerDistributions(distributionsService)
	r.HandleFunc("/games/{id}", gamesHandler.GetGame).Methods(http.MethodGet)
	r.HandleFunc("/games", gamesHandler.PostGame).Methods(http.MethodPost)
	r.HandleFunc("/games/{id}/distributions", distributionsHandler.GetDistribution).Methods(http.MethodGet)
	r.HandleFunc("/games/{id}/distributions", distributionsHandler.PatchDistribution).Methods(http.MethodPatch)
	http.Handle("/", r)
	_ = http.ListenAndServe(":8090", nil)
}
