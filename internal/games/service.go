package games

import (
	"context"
	"errors"
	"github/ibanezv/minesweeper-API/internal/distributions"
	"github/ibanezv/minesweeper-API/internal/models"
	"github/ibanezv/minesweeper-API/internal/repository"
	"github/ibanezv/minesweeper-API/pkg/database"
)

var ErrGameNotFound = errors.New("game not found")
var ErrGameBadRequest = errors.New("game data error")

type Service interface {
	FindGame(context.Context, int64) (models.Game, error)
	CreateGame(context.Context, models.Game) (models.Game, error)
}

type ProcessGame struct {
	repositories         database.Repositories
	distributionsService distributions.Service
}

func NewService(repositories database.Repositories, distributionsService distributions.Service) *ProcessGame {
	return &ProcessGame{repositories, distributionsService}
}

func (g *ProcessGame) FindGame(ctx context.Context, ID int64) (models.Game, error) {
	dbGame, err := g.repositories.GetGameById(ctx, ID)

	if err != nil {
		return models.Game{}, err
	}

	game := Transform(dbGame)
	return game, nil
}

func (g *ProcessGame) CreateGame(ctx context.Context, game models.Game) (models.Game, error) {
	if !validationData(game) {
		return models.Game{}, ErrGameBadRequest
	}

	distribution := g.distributionsService.CreateDistribution(ctx, game.CountRows, game.CountCols, game.CountMines)
	dbGame := transformToDB(game)
	newGame, err := g.repositories.CreateGame(ctx, dbGame)
	if err != nil {
		return models.Game{}, err
	}
	for i := 0; i < len(distribution); i++ {
		row := distribution[i]
		for j := 0; j < len(row); j++ {
			d := row[j]
			d.GameID = int64(newGame.ID)
			_, err = g.distributionsService.AddDistribution(ctx, d)
			if err != nil {
				return models.Game{}, err
			}
		}
	}
	return Transform(newGame), nil
}

func validationData(game models.Game) bool {
	return game.CountRows >= 4 && game.CountCols >= 4 && game.CountMines > 0
}

func Transform(dbGame repository.Games) models.Game {
	var game models.Game
	game.ID = int64(dbGame.ID)
	game.CountRows = dbGame.CountRows
	game.CountCols = dbGame.CountCols
	game.CountMines = dbGame.CountMines
	game.UserID = dbGame.UserID
	return game
}

func transformToDB(game models.Game) repository.Games {
	var dbGame repository.Games
	dbGame.CountRows = game.CountRows
	dbGame.CountCols = game.CountCols
	dbGame.CountMines = game.CountMines
	dbGame.ID = game.ID
	dbGame.UserID = game.UserID
	dbGame.State = game.State
	return dbGame
}
