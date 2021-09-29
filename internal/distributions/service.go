package distributions

import (
	"context"
	"errors"
	"github/ibanezv/minesweeper-API/internal/models"
	"github/ibanezv/minesweeper-API/internal/repository"
	"github/ibanezv/minesweeper-API/pkg/database"
	"math/rand"
	"reflect"
	"time"
)

const (
	CellValueMine         = "mine"
	CellValueSelected     = "selected"
	CellValueQuestionMark = "question"
	CellValueFlagMark     = "flag"
	CellStateShowed       = "showed"
	CellStateHidden       = "hidden"
	StateGameOver         = "over"
	StateGameSuccess      = "game was finished"
)

var ErrMine = errors.New("mine selected")
var ErrInvalidRequest = errors.New("invalid request data")
var ErrGameOver = errors.New("game is over")
var ErrGameFinished = errors.New("game is finished")

type Service interface {
	FindDistribution(context.Context, int64) ([]models.Distribution, error)
	UpdateCellDistribution(context.Context, models.Distribution) (bool, error)
	CreateDistribution(context.Context, int, int, int) [][]models.Distribution
	AddDistribution(context.Context, models.Distribution) (models.Distribution, error)
	validateCompleteDistribution(context.Context, repository.Games) (bool, error)
}

type ProcessDistribution struct {
	repositories database.Repositories
}

func NewService(repositories database.Repositories) *ProcessDistribution {
	return &ProcessDistribution{repositories}
}

func (d *ProcessDistribution) FindDistribution(ctx context.Context, gameID int64) ([]models.Distribution, error) {
	dbDistribution, err := d.repositories.GetDistributionByGameId(ctx, gameID)
	if err != nil {
		return []models.Distribution{}, err
	}
	distribution := transformList(dbDistribution)
	return distribution, nil
}

func (d *ProcessDistribution) AddDistribution(ctx context.Context, distribution models.Distribution) (models.Distribution, error) {
	dbDistribution := tranformToDB(distribution)
	newDbDistribution, err := d.repositories.CreateDistribution(ctx, dbDistribution)
	if err != nil {
		return models.Distribution{}, err
	}
	return transform(newDbDistribution), nil
}

func (d *ProcessDistribution) CreateDistribution(ctx context.Context, rowCount int, colCount int, minesCount int) [][]models.Distribution {
	lstMinesPositions := makeMinesPosition(rowCount*colCount, minesCount)
	lstDistributions := make([][]models.Distribution, 0, rowCount)
	count := 0
	for i := 0; i < rowCount; i++ {
		row := make([]models.Distribution, 0, colCount)
		for j := 0; j < colCount; j++ {
			distribution := models.Distribution{}
			if contains(lstMinesPositions, count) {
				distribution.Value = CellValueMine
			}
			distribution.ColNumber = j
			distribution.RowNumber = i
			distribution.State = CellStateHidden
			row = append(row, distribution)
			count++
		}
		lstDistributions = append(lstDistributions, row)
	}
	return lstDistributions
}

func makeMinesPosition(limit, count int) []int {
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	minesPos := make([]int, count)
	var num int
	i := 0
	for i < count {
		num = rand.Intn(limit)
		if !contains(minesPos, num) {
			minesPos[i] = num
			i++
		}
	}
	return minesPos
}

func contains(list []int, element int) bool {
	for _, value := range list {
		if element == value {
			return true
		}
	}
	return false
}

func (d *ProcessDistribution) UpdateCellDistribution(ctx context.Context, distribution models.Distribution) (bool, error) {
	game, err := d.repositories.GetGameById(ctx, distribution.GameID)
	if err != nil {
		return false, err
	}

	dbDistribution, err := d.repositories.GetDistributionByGameId(ctx, distribution.GameID)
	if err != nil {
		return false, err
	}

	if err := validationDataRequest(game, dbDistribution, distribution); err != nil {
		return false, err
	}

	arr := convertToSlice(ctx, game.CountRows, game.CountCols, dbDistribution)

	if checkIfMine(arr, distribution) {
		game.State = StateGameOver
		d.repositories.UpdateGame(ctx, game)
		return false, ErrMine
	}

	if distribution.Value == CellValueFlagMark || distribution.Value == CellValueQuestionMark {
		arr = markCell(ctx, arr, distribution)
	} else {
		arr = selectCell(ctx, arr, distribution.RowNumber, distribution.ColNumber)
	}

	for i, _ := range arr {
		for j, _ := range arr[i] {
			_, err := d.repositories.UpdateDistributionCell(ctx, arr[i][j])
			if err != nil {
				return false, err
			}
		}
	}

	isComplete, err := d.validateCompleteDistribution(ctx, game)
	if err != nil {
		return false, err
	}
	if isComplete {
		return true, nil
	}

	return false, nil
}

func (d *ProcessDistribution) validateCompleteDistribution(ctx context.Context, game repository.Games) (bool, error) {
	countSelected, err := d.repositories.GetDistributionCellSelected(ctx, game.ID, CellStateShowed, CellValueMine)
	if err != nil {
		return false, err
	}
	return (game.CountCols*game.CountRows)-game.CountMines == countSelected, nil
}

func validationDataRequest(game repository.Games, dbDistribution []repository.Distributions, distribution models.Distribution) error {
	if reflect.DeepEqual(game, repository.Games{}) {
		return ErrInvalidRequest
	}

	if reflect.DeepEqual(dbDistribution, repository.Distributions{}) {
		return errors.New("game distribution not found")
	}

	size := len(dbDistribution)
	if size != game.CountCols*game.CountRows {
		return errors.New("invalid distribution for game")
	}

	if !validateCellSelected(game, distribution.RowNumber, distribution.ColNumber) {
		return ErrInvalidRequest
	}

	if !validateCellValue(distribution.Value) {
		return ErrInvalidRequest
	}

	if game.State == StateGameOver {
		return ErrGameOver
	}
	return nil
}

func markCell(ctx context.Context, matrix [][]repository.Distributions, distribution models.Distribution) [][]repository.Distributions {
	matrix[distribution.RowNumber][distribution.ColNumber].Value = distribution.Value
	return matrix
}

func selectCell(ctx context.Context, matrix [][]repository.Distributions, selRow, selCol int) [][]repository.Distributions {
	rowLen := len(matrix)
	ColLen := len(matrix[0])
	if matrix[selRow][selCol].Value == "" {
		matrix[selRow][selCol].Value = CellValueSelected
		initRow := selRow - 1
		if initRow < 0 {
			initRow = 0
		}

		endRow := selRow + 1
		if endRow > rowLen {
			endRow = rowLen
		}

		initCol := selCol - 1
		if initCol < 0 {
			initCol = 0
		}

		endCol := selCol + 1
		if endCol > ColLen {
			endCol = ColLen
		}

		for i := initRow; i <= endRow; i++ {
			for j := initCol; j <= endCol; j++ {
				if matrix[i][j].Value != CellValueMine {
					matrix[i][j].Value = CellValueSelected
					matrix[i][j].State = CellStateShowed
				}
			}
		}
	}
	return matrix
}

func checkIfMine(arrDistribution [][]repository.Distributions, distribution models.Distribution) bool {
	return (distribution.Value == CellValueSelected &&
		arrDistribution[distribution.RowNumber][distribution.ColNumber].Value == CellValueMine)
}
func validateCellSelected(game repository.Games, row, col int) bool {
	return (row <= game.CountRows && col <= game.CountCols)
}

func validateCellValue(value string) bool {
	return (value == CellValueSelected || value == CellValueFlagMark || value == CellValueQuestionMark)
}

func convertToSlice(ctx context.Context, countRows int, countCols int,
	dbDistribution []repository.Distributions) [][]repository.Distributions {
	arr := make([][]repository.Distributions, countRows)
	for i := range arr {
		arr[i] = make([]repository.Distributions, countCols)
	}

	for _, item := range dbDistribution {
		arr[item.RowNumber][item.ColNumber].GameID = item.GameID
		arr[item.RowNumber][item.ColNumber].Value = item.Value
		arr[item.RowNumber][item.ColNumber].State = item.State
		arr[item.RowNumber][item.ColNumber].RowNumber = item.RowNumber
		arr[item.RowNumber][item.ColNumber].ColNumber = item.ColNumber
	}
	return arr
}

func transformList(dbDistributions []repository.Distributions) []models.Distribution {
	var distributions []models.Distribution
	for _, item := range dbDistributions {
		dist := models.Distribution{}
		dist.GameID = int64(item.GameID)
		dist.RowNumber = item.RowNumber
		dist.ColNumber = item.ColNumber
		dist.State = item.State
		dist.Value = item.Value
		distributions = append(distributions, dist)
	}
	return distributions
}

func transform(dbDistributions repository.Distributions) models.Distribution {
	dist := models.Distribution{}
	dist.GameID = int64(dbDistributions.GameID)
	dist.RowNumber = dbDistributions.RowNumber
	dist.ColNumber = dbDistributions.ColNumber
	dist.State = dbDistributions.State
	dist.Value = dbDistributions.Value
	return dist
}

func tranformToDB(distribution models.Distribution) repository.Distributions {
	dbDistribution := repository.Distributions{}
	dbDistribution.ColNumber = distribution.ColNumber
	dbDistribution.RowNumber = distribution.RowNumber
	dbDistribution.GameID = distribution.GameID
	dbDistribution.State = distribution.State
	dbDistribution.Value = distribution.Value
	return dbDistribution

}
