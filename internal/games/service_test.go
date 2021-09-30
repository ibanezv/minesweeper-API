package games

import (
	"context"
	"errors"
	"github/ibanezv/minesweeper-API/internal/distributions"
	"github/ibanezv/minesweeper-API/internal/models"
	"github/ibanezv/minesweeper-API/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateGame(t *testing.T) {
	//given
	tests := []struct {
		name          string
		game          models.Game
		expectedError bool
	}{
		{
			name:          "create game success",
			game:          models.Game{ID: 0, CountRows: 4, CountCols: 4, CountMines: 1, State: "in_progres"},
			expectedError: false,
		},
		{
			name:          "create game fail countRows invalid",
			game:          models.Game{ID: 0, CountRows: 3, CountCols: 4, CountMines: 1, State: "in_progres"},
			expectedError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distributionService := distributions.NewService(RepositoryMock{})
			service := NewService(RepositoryMock{}, distributionService)
			//when
			_, err := service.CreateGame(context.Background(), tt.game)

			//then
			assert.Equal(t, tt.expectedError, err != nil)
		})
	}
}

func Test_FindGame(t *testing.T) {
	//given
	tests := []struct {
		name          string
		ID            int
		expectedError bool
	}{
		{
			name:          "find game success",
			ID:            1,
			expectedError: false,
		},
		{
			name:          "find game not found",
			ID:            2,
			expectedError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distributionService := distributions.NewService(RepositoryMock{})
			service := NewService(RepositoryMock{}, distributionService)
			//when
			_, err := service.FindGame(context.Background(), int64(tt.ID))

			//then
			assert.Equal(t, tt.expectedError, err != nil)
		})
	}
}

type RepositoryMock struct {
}

func (r RepositoryMock) GetGameById(ctx context.Context, ID int64) (repository.Games, error) {
	if ID == 1 {
		return repository.Games{ID: ID, CountRows: 3, CountCols: 3, CountMines: 1, UserID: 1}, nil
	}
	if ID == 2 {
		return repository.Games{}, errors.New("test error")
	}
	return repository.Games{}, nil
}

func (r RepositoryMock) CreateGame(ctx context.Context, game repository.Games) (repository.Games, error) {
	return repository.Games{ID: 1, UserID: game.ID, CountRows: game.CountRows, CountCols: game.CountCols, State: game.State}, nil
}
func (r RepositoryMock) UpdateGame(ctx context.Context, game repository.Games) (repository.Games, error) {
	return repository.Games{}, nil
}

func (r RepositoryMock) GetDistributionByGameId(ctx context.Context, ID int64) ([]repository.Distributions, error) {
	return nil, nil
}
func (r RepositoryMock) GetDistributionCell(ctx context.Context, ID int64, row int, col int) (repository.Distributions, error) {
	return repository.Distributions{}, nil
}
func (r RepositoryMock) UpdateDistributionCell(ctx context.Context, dbDistribution repository.Distributions) (repository.Distributions, error) {
	return repository.Distributions{}, nil
}
func (r RepositoryMock) CreateDistribution(ctx context.Context, dbDistribution repository.Distributions) (repository.Distributions, error) {
	return repository.Distributions{}, nil
}
func (r RepositoryMock) GetDistributionCellSelected(ctx context.Context, gameID int64, state string, notValue string) (int, error) {
	return 0, nil
}
func (r RepositoryMock) GetUserById(ctx context.Context, ID int) (repository.Users, error) {
	return repository.Users{}, nil
}

func (r RepositoryMock) CreateUser(ctx context.Context, user repository.Users) (repository.Users, error) {
	return repository.Users{}, nil
}
func (r RepositoryMock) GetUserGames(ctx context.Context, ID int) ([]repository.Games, error) {
	return nil, nil
}
func (r RepositoryMock) GetAccountById(ctx context.Context, ID int) (repository.Accounts, error) {
	return repository.Accounts{}, nil
}
func (r RepositoryMock) CreateAccount(ctx context.Context, accountDb repository.Accounts) (repository.Accounts, error) {
	return repository.Accounts{}, nil
}
func (r RepositoryMock) GetAccountUsers(ctx context.Context, ID int) ([]repository.Users, error) {
	return nil, nil
}
