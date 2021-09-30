package distributions

import (
	"context"
	"errors"
	"github/ibanezv/minesweeper-API/internal/models"
	"github/ibanezv/minesweeper-API/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateDistribution(t *testing.T) {
	//given
	var tests = []struct {
		name       string
		countRows  int
		countCols  int
		countMines int
	}{
		{
			name:       "create success",
			countRows:  10,
			countCols:  15,
			countMines: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewService(RepositoryMock{})
			//when
			distributions := service.CreateDistribution(context.Background(), tt.countRows, tt.countCols, tt.countMines)

			//then
			assert.Equal(t, tt.countRows, len(distributions))
			assert.Equal(t, tt.countCols, len(distributions[0]))
			assert.Equal(t, tt.countMines, getCountMines(distributions))
		})
	}
}

func Test_FindDistribution(t *testing.T) {
	//given
	var tests = []struct {
		name          string
		ID            int64
		expectedError bool
	}{
		{
			name:          "get distribution success",
			ID:            1,
			expectedError: false,
		},
		{
			name:          "get distribution fail",
			ID:            2,
			expectedError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewService(RepositoryMock{})

			//when
			_, err := service.FindDistribution(context.Background(), tt.ID)

			//then
			assert.Equal(t, tt.expectedError, err != nil)

		})
	}
}

func Test_AddDistribution(t *testing.T) {
	//given
	var tests = []struct {
		name          string
		distribution  models.Distribution
		expectedError bool
	}{
		{
			name:          "add distribution success",
			distribution:  models.Distribution{GameID: 1, RowNumber: 1, ColNumber: 3, Value: "selected", State: "hidden"},
			expectedError: false,
		},
		{
			name:          "add distribution fail",
			distribution:  models.Distribution{GameID: 2, RowNumber: 1, ColNumber: 3, Value: "selected", State: "hidden"},
			expectedError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewService(RepositoryMock{})

			//when
			_, err := service.AddDistribution(context.Background(), tt.distribution)

			//then
			assert.Equal(t, tt.expectedError, err != nil)

		})
	}
}

func Test_UpdateCellDistribution(t *testing.T) {
	//given
	var tests = []struct {
		name            string
		distribution    models.Distribution
		expectedError   bool
		gameIsCompleted bool
	}{
		{
			name:            "update distribution success",
			distribution:    models.Distribution{GameID: 1, RowNumber: 1, ColNumber: 1, Value: "selected", State: "hidden"},
			expectedError:   false,
			gameIsCompleted: false,
		},
		{
			name:            "update distribution mine selected",
			distribution:    models.Distribution{GameID: 1, RowNumber: 1, ColNumber: 2, Value: "selected", State: "hidden"},
			expectedError:   true,
			gameIsCompleted: false,
		},
		{
			name:            "update distribution fail",
			distribution:    models.Distribution{GameID: 2, RowNumber: 1, ColNumber: 2, Value: "selected", State: "hidden"},
			expectedError:   true,
			gameIsCompleted: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewService(RepositoryMock{})

			//when
			isComplete, err := service.UpdateCellDistribution(context.Background(), tt.distribution)

			//then
			assert.Equal(t, tt.expectedError, err != nil)
			assert.Equal(t, tt.gameIsCompleted, isComplete)

		})
	}
}

func getCountMines(distribution [][]models.Distribution) int {
	count := 0
	for _, row := range distribution {
		for _, dist := range row {
			if dist.Value == CellValueMine {
				count++
			}
		}
	}
	return count
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
	return repository.Games{}, nil
}
func (r RepositoryMock) UpdateGame(ctx context.Context, game repository.Games) (repository.Games, error) {
	return repository.Games{}, nil
}

func (r RepositoryMock) GetDistributionByGameId(ctx context.Context, ID int64) ([]repository.Distributions, error) {
	if ID == 2 {
		return nil, errors.New("test error")
	}
	if ID == 1 {
		distributions := []repository.Distributions{}
		dist1 := repository.Distributions{GameID: ID, RowNumber: 0, ColNumber: 0, Value: "", State: "hidden"}
		dist2 := repository.Distributions{GameID: ID, RowNumber: 0, ColNumber: 1, Value: "", State: "hidden"}
		dist3 := repository.Distributions{GameID: ID, RowNumber: 0, ColNumber: 2, Value: "", State: "hidden"}
		distributions = append(distributions, dist1)
		distributions = append(distributions, dist2)
		distributions = append(distributions, dist3)

		dist4 := repository.Distributions{GameID: ID, RowNumber: 1, ColNumber: 0, Value: "", State: "hidden"}
		dist5 := repository.Distributions{GameID: ID, RowNumber: 1, ColNumber: 1, Value: "", State: "hidden"}
		dist6 := repository.Distributions{GameID: ID, RowNumber: 1, ColNumber: 2, Value: "mine", State: "hidden"}
		distributions = append(distributions, dist4)
		distributions = append(distributions, dist5)
		distributions = append(distributions, dist6)

		dist7 := repository.Distributions{GameID: ID, RowNumber: 2, ColNumber: 0, Value: "", State: "hidden"}
		dist8 := repository.Distributions{GameID: ID, RowNumber: 2, ColNumber: 1, Value: "", State: "hidden"}
		dist9 := repository.Distributions{GameID: ID, RowNumber: 2, ColNumber: 2, Value: "", State: "hidden"}
		distributions = append(distributions, dist7)
		distributions = append(distributions, dist8)
		distributions = append(distributions, dist9)

		return distributions, nil
	}
	return nil, nil
}
func (r RepositoryMock) GetDistributionCell(ctx context.Context, ID int64, row int, col int) (repository.Distributions, error) {
	return repository.Distributions{}, nil
}
func (r RepositoryMock) UpdateDistributionCell(ctx context.Context, dbDistribution repository.Distributions) (repository.Distributions, error) {
	return repository.Distributions{}, nil
}
func (r RepositoryMock) CreateDistribution(ctx context.Context, dbDistribution repository.Distributions) (repository.Distributions, error) {
	if dbDistribution.GameID == 1 {
		return repository.Distributions{}, nil
	}
	if dbDistribution.GameID == 2 {
		return repository.Distributions{}, errors.New("test error")
	}
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
