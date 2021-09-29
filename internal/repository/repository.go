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
	GetUserById(ctx context.Context, ID int) (Users, error)
	CreateUser(ctx context.Context, user Users) (Users, error)
	GetUserGames(ctx context.Context, ID int) ([]Games, error)
	GetAccountById(ctx context.Context, ID int) (Accounts, error)
	CreateAccount(ctx context.Context, accountDb Accounts) (Accounts, error)
	GetAccountUsers(ctx context.Context, ID int) ([]Users, error)
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

func (dao Dao) GetUserById(ctx context.Context, ID int) (Users, error) {
	userDb := Users{}
	result := dao.DB.Table("users").Where("id=?", ID).First(&userDb)
	return userDb, result.Error
}

func (dao Dao) CreateUser(ctx context.Context, user Users) (Users, error) {
	result := dao.DB.Create(&user)
	return user, result.Error
}

func (dao Dao) GetUserGames(ctx context.Context, ID int) ([]Games, error) {
	userGames := []Games{}
	result := dao.DB.Table("games").Where("user_id=?", ID).Find(&userGames)
	return userGames, result.Error
}

func (dao Dao) GetAccountById(ctx context.Context, ID int) (Accounts, error) {
	accountDb := Accounts{}
	result := dao.DB.Table("accounts").Where("id=?", ID).First(&accountDb)
	return accountDb, result.Error
}

func (dao Dao) CreateAccount(ctx context.Context, accountDb Accounts) (Accounts, error) {
	result := dao.DB.Create(&accountDb)
	return accountDb, result.Error
}

func (dao Dao) GetAccountUsers(ctx context.Context, ID int) ([]Users, error) {
	accountUsers := []Users{}
	result := dao.DB.Table("users").Where("account_id=?", ID).Find(&accountUsers)
	return accountUsers, result.Error
}
