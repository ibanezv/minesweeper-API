package accounts

import (
	"context"
	"github/ibanezv/minesweeper-API/internal/models"
	"github/ibanezv/minesweeper-API/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateAccount(t *testing.T) {
	tests := []struct {
		name          string
		account       models.Accounts
		expectedError bool
	}{
		{
			name:          "create account success",
			account:       models.Accounts{ID: 0, Email: "test@test"},
			expectedError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountService := NewService(RepositoryMock{})
			_, err := accountService.CreateAccount(context.Background(), tt.account)
			assert.Equal(t, tt.expectedError, err != nil)
		})
	}
}

func Test_FindAccount(t *testing.T) {
	tests := []struct {
		name          string
		id            int
		expectedError bool
	}{
		{
			name:          "create account success",
			id:            1,
			expectedError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountService := NewService(RepositoryMock{})
			account, err := accountService.FindAccount(context.Background(), tt.id)
			assert.Equal(t, tt.expectedError, err != nil)
			assert.Equal(t, tt.id, account.ID)
		})
	}
}

func Test_FindAccountUsers(t *testing.T) {
	tests := []struct {
		name               string
		id                 int
		expectedError      bool
		countUsersExpected int
	}{
		{
			name:               "create account success",
			id:                 1,
			expectedError:      false,
			countUsersExpected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountService := NewService(RepositoryMock{})
			users, err := accountService.FindAccountUsers(context.Background(), tt.id)
			assert.Equal(t, tt.expectedError, err != nil)
			assert.Equal(t, tt.countUsersExpected, len(users))
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
	return repository.Users{ID: uint(ID), NickName: "test"}, nil
}

func (r RepositoryMock) CreateUser(ctx context.Context, user repository.Users) (repository.Users, error) {
	return repository.Users{ID: 1, NickName: "test", AccountID: 1}, nil
}
func (r RepositoryMock) GetUserGames(ctx context.Context, ID int) ([]repository.Games, error) {
	return []repository.Games{{ID: 1, CountRows: 4, CountCols: 4, CountMines: 2}, {ID: 1, CountRows: 4, CountCols: 4, CountMines: 2}}, nil
}
func (r RepositoryMock) GetAccountById(ctx context.Context, ID int) (repository.Accounts, error) {
	return repository.Accounts{ID: 1, Email: "test@test"}, nil
}
func (r RepositoryMock) CreateAccount(ctx context.Context, accountDb repository.Accounts) (repository.Accounts, error) {
	return repository.Accounts{}, nil
}
func (r RepositoryMock) GetAccountUsers(ctx context.Context, ID int) ([]repository.Users, error) {
	return []repository.Users{{ID: 1, NickName: "test1", AccountID: 1}, {ID: 2, NickName: "test2", AccountID: 1}}, nil
}
