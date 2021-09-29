package repository

import (
	"context"

	"gorm.io/gorm"
)

type IGames interface {
	GetGameById(ctx context.Context, ID int64) (Games, error)
	CreateGame(ctx context.Context, game Games) (Games, error)
	UpdateGame(ctx context.Context, game Games) (Games, error)
	GetDistributionByGameId(ctx context.Context, ID int64) ([]Distributions, error)
	GetDistributionCell(ctx context.Context, ID int64, row int, col int) (Distributions, error)
	UpdateDistributionCell(ctx context.Context, dbDistribution Distributions) (Distributions, error)
	CreateDistribution(ctx context.Context, dbDistribution Distributions) (Distributions, error)
	GetDistributionCellSelected(ctx context.Context, gameID int64, state string, notValue string) (int, error)
}

type Dao struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Dao {
	return Dao{db}
}

func (dao Dao) GetGameById(ctx context.Context, id int64) (Games, error) {
	games := Games{}
	result := dao.DB.Table("games").Where("id=?", id).First(&games)
	return games, result.Error
}

func (dao Dao) CreateGame(ctx context.Context, game Games) (Games, error) {
	result := dao.DB.Create(&game)
	return game, result.Error
}

func (dao Dao) UpdateGame(ctx context.Context, game Games) (Games, error) {
	newGame := Games{}
	result := dao.DB.Where(Games{ID: game.ID}).
		Assign(game).
		FirstOrCreate(&newGame)
	return newGame, result.Error
}

func (dao Dao) GetDistributionByGameId(ctx context.Context, gameID int64) ([]Distributions, error) {
	distributions := []Distributions{}
	result := dao.DB.Table("distributions").Where("game_id=?", gameID).Find(&distributions)
	return distributions, result.Error
}

func (dao Dao) GetDistributionCell(ctx context.Context, gameID int64, row int, col int) (Distributions, error) {
	distributions := Distributions{}
	result := dao.DB.Table("distributions").Where("game_id=?", gameID, "row_number=?", row, "col_number=?", col).First(&distributions)
	return distributions, result.Error
}

func (dao Dao) GetDistributionCellSelected(ctx context.Context, gameID int64, state string, notValue string) (int, error) {
	distributions := []Distributions{}
	result := dao.DB.Where(Distributions{GameID: gameID, State: state}).Not(Distributions{Value: notValue}).Find(&distributions)
	return len(distributions), result.Error
}

func (dao Dao) UpdateDistributionCell(ctx context.Context, dbDistribution Distributions) (Distributions, error) {
	distributions := Distributions{}
	result := dao.DB.Where(Distributions{GameID: dbDistribution.GameID, RowNumber: dbDistribution.RowNumber, ColNumber: dbDistribution.ColNumber}).
		Assign(dbDistribution).
		FirstOrCreate(&distributions)
	return distributions, result.Error
}

func (dao Dao) CreateDistribution(ctx context.Context, dbDistribution Distributions) (Distributions, error) {
	result := dao.DB.Create(&dbDistribution)
	return dbDistribution, result.Error
}
