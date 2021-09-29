package handlers

import (
	"encoding/json"
	"errors"
	"github/ibanezv/minesweeper-API/internal/distributions"
	"github/ibanezv/minesweeper-API/internal/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	MessageGameOver            = "game over"
	MessageGameFinished        = "game finished successfully"
	MessageDistributionUpdated = "distribution updated"
	HeaderKey                  = "Content-Type"
	HeaderValue                = "application/json"
)

type DistributionsHandler struct {
	serviceDistributions distributions.Service
}

func NewHandlerDistributions(serviceDistributions distributions.Service) DistributionsHandler {
	return DistributionsHandler{serviceDistributions: serviceDistributions}
}

func (d *DistributionsHandler) GetDistribution(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var strGameID = params["id"]
	gameID, err := strconv.ParseInt(strGameID, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ctx := r.Context()
	collDistributions, err := d.serviceDistributions.FindDistribution(ctx, gameID)
	if err != nil {
		if errors.Is(err, distributions.ErrInvalidRequest) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(HeaderKey, HeaderValue)
	jsonResp, _ := json.Marshal(collDistributions)
	w.Write(jsonResp)
}

func (d *DistributionsHandler) PatchDistribution(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var strGameID = params["id"]
	gameID, err := strconv.ParseInt(strGameID, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	distribution, err := getDistributionInput(r)
	distribution.GameID = gameID

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ctx := r.Context()
	isCompleted, err := d.serviceDistributions.UpdateCellDistribution(ctx, distribution)
	if err != nil {
		if errors.Is(err, distributions.ErrInvalidRequest) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if errors.Is(err, distributions.ErrMine) {
			http.Error(w, MessageGameOver, http.StatusBadRequest)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		resp := make(map[string]string)
		resp["message"] = MessageDistributionUpdated
		if isCompleted {
			resp["message"] = MessageGameFinished
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set(HeaderKey, HeaderValue)

		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
	}
}

func getDistributionInput(request *http.Request) (models.Distribution, error) {
	var input models.Distribution
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		return models.Distribution{}, err
	}
	return input, nil
}
