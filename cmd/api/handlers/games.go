package handlers

import (
	"encoding/json"
	"errors"
	"github/ibanezv/minesweeper-API/internal/games"
	"github/ibanezv/minesweeper-API/internal/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type GamesHandler struct {
	serviceGames games.Service
}

func NewHandlerGames(serviceGames games.Service) GamesHandler {
	return GamesHandler{serviceGames: serviceGames}
}

func (g *GamesHandler) GetGame(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var strGameID = params["id"]
	gameID, err := strconv.ParseInt(strGameID, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	game, err := g.serviceGames.FindGame(ctx, gameID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(HeaderKey, HeaderValue)
	jsonResp, _ := json.Marshal(game)
	w.Write(jsonResp)
}

func (g *GamesHandler) PostGame(w http.ResponseWriter, r *http.Request) {
	game, err := getGameInput(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ctx := r.Context()
	newGame, err := g.serviceGames.CreateGame(ctx, game)
	if err != nil {
		if errors.Is(err, games.ErrGameBadRequest) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set(HeaderKey, HeaderValue)
	jsonResp, _ := json.Marshal(newGame)
	w.Write(jsonResp)
}

func getGameInput(request *http.Request) (models.Game, error) {
	var input models.Game
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		return models.Game{}, err
	}
	return input, nil
}
