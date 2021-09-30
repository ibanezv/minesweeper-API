package users

import (
	"context"
	"errors"
	"github/ibanezv/minesweeper-API/internal/models"
	"github/ibanezv/minesweeper-API/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindUser(t *testing.T) {
	//given
	tests := []struct {
		name          string
		ID            int
		expectedError bool
	}{
		{
			name:          "find succes",
			ID:            1,
			expectedError: false,
		},
		{
			name:          "find user not found",
			ID:            2,
			expectedError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userService := NewService(RepositoryMock{})
			user, err := userService.FindUser(context.Background(), tt.ID)
			assert.Equal(t, tt.expectedError, err != nil)
			if err == nil {
				assert.Equal(t, tt.ID, user.ID)
			}
		})
	}
}

func Test_CreateUser(t *testing.T) {
	//given
	tests := []struct {
		name          string
		user          models.User
		expectedError bool
		expectedID    int
	}{
		{
			name:          "create user succes",
			user:          models.User{ID: 0, NickName: "test", AccountID: 1},
			expectedError: false,
			expectedID:    1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userService := NewService(RepositoryMock{})
			user, err := userService.CreateUser(context.Background(), tt.user)
			assert.Equal(t, tt.expectedError, err != nil)
			assert.Equal(t, tt.expectedID, user.ID)
		})
	}
}

func Test_GetUserGames(t *testing.T) {
	//given
	tests := []struct {
		name          string
		id            int
		expectedError bool
		countExpected int
	}{
		{
			name:          "get user games",
			id:            1,
			expectedError: false,
			countExpected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userService := NewService(RepositoryMock{})
			games, err := userService.FindUserGames(context.Background(), tt.id)
			assert.Equal(t, tt.expectedError, err != nil)
			assert.Equal(t, tt.countExpected, len(games))
		})
	}
}

type RepositoryMock struct {
}

func (r RepositoryMock) GetGameById(ctx context.Context, ID int64) (repository.Games, error) {
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
	if ID == 1 {
		return repository.Users{ID: uint(ID), NickName: "test"}, nil
	}
	if ID == 2 {
		return repository.Users{ID: uint(ID), NickName: "test"}, errors.New("test")
	}
	return repository.Users{ID: uint(ID), NickName: "test"}, nil
}

func (r RepositoryMock) CreateUser(ctx context.Context, user repository.Users) (repository.Users, error) {
	return repository.Users{ID: 1, NickName: "test", AccountID: 1}, nil
}
func (r RepositoryMock) GetUserGames(ctx context.Context, ID int) ([]repository.Games, error) {
	return []repository.Games{{ID: 1, CountRows: 4, CountCols: 4, CountMines: 2}, {ID: 1, CountRows: 4, CountCols: 4, CountMines: 2}}, nil
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
