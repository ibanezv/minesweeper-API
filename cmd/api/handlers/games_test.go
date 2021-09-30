package handlers

import (
	"context"
	"encoding/json"
	"github/ibanezv/minesweeper-API/internal/models"
	"github/ibanezv/minesweeper-API/internal/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_PostGame(t *testing.T) {
	//given
	tests := []struct {
		name         string
		url          string
		game         models.Game
		codeExpected int
	}{
		{
			name:         "post game success",
			url:          "/game",
			game:         models.Game{ID: 0, CountRows: 4, CountCols: 4, CountMines: 1},
			codeExpected: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, _ := json.Marshal(tt.game)
			req := httptest.NewRequest(http.MethodPost, tt.url, strings.NewReader(string(e)))
			w := httptest.NewRecorder()
			h := NewHandlerGames(gameServiceMock{})

			//when
			h.PostGame(w, req)
			resp := w.Result()
			defer resp.Body.Close()

			//then
			assert.Equal(t, tt.codeExpected, resp.StatusCode)
		})
	}
}

func Test_GetGame(t *testing.T) {
	//given
	tests := []struct {
		name         string
		url          string
		idExpected   int
		codeExpected int
	}{
		{
			name:         "get game success",
			url:          "/game/1",
			idExpected:   1,
			codeExpected: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()
			h := NewHandlerGames(gameServiceMock{})
			vars := map[string]string{
				"id": "1",
			}
			req = mux.SetURLVars(req, vars)

			//when
			h.GetGame(w, req)
			resp := w.Result()
			defer resp.Body.Close()

			//then
			assert.Equal(t, tt.codeExpected, resp.StatusCode)
		})
	}
}

type gameServiceMock struct {
}

func (g gameServiceMock) FindGame(context.Context, int64) (models.Game, error) {
	return models.Game{ID: 1, CountRows: 4, CountCols: 4, CountMines: 1}, nil
}
func (g gameServiceMock) CreateGame(context.Context, models.Game) (models.Game, error) {
	return models.Game{}, nil
}

type distributionServiceMock struct {
}

func (d distributionServiceMock) FindDistribution(context.Context, int64) ([]models.Distribution, error) {
	return nil, nil
}
func (d distributionServiceMock) UpdateCellDistribution(context.Context, models.Distribution) (bool, error) {
	return false, nil
}
func (d distributionServiceMock) CreateDistribution(context.Context, int, int, int) [][]models.Distribution {
	return nil
}
func (d distributionServiceMock) AddDistribution(context.Context, models.Distribution) (models.Distribution, error) {
	return models.Distribution{}, nil
}
func (d distributionServiceMock) ValidateCompleteDistribution(context.Context, repository.Games) (bool, error) {
	return false, nil
}

type userServiceMock struct {
}

func (u userServiceMock) FindUser(context.Context, int) (models.User, error) {
	return models.User{ID: 1, NickName: "test"}, nil
}
func (u userServiceMock) CreateUser(context.Context, models.User) (models.User, error) {
	return models.User{}, nil
}
func (u userServiceMock) FindUserGames(ctx context.Context, ID int) ([]models.Game, error) {
	return nil, nil
}

type serviceAccountsMock struct {
}

func (a serviceAccountsMock) FindAccount(context.Context, int) (models.Accounts, error) {
	return models.Accounts{}, nil
}
func (a serviceAccountsMock) CreateAccount(context.Context, models.Accounts) (models.Accounts, error) {
	return models.Accounts{}, nil
}
func (a serviceAccountsMock) FindAccountUsers(ctx context.Context, ID int) ([]models.User, error) {
	return nil, nil
}
