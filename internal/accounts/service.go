package accounts

import (
	"context"
	"github/ibanezv/minesweeper-API/internal/models"
	"github/ibanezv/minesweeper-API/internal/repository"
	"github/ibanezv/minesweeper-API/internal/users"
	"github/ibanezv/minesweeper-API/pkg/database"
)

type Service interface {
	FindAccount(context.Context, int) (models.Accounts, error)
	CreateAccount(context.Context, models.Accounts) (models.Accounts, error)
	FindAccountUsers(ctx context.Context, ID int) ([]models.User, error)
}

type ProcessAccounts struct {
	repositories database.Repositories
}

func NewService(repositories database.Repositories) *ProcessAccounts {
	return &ProcessAccounts{repositories}
}

func (a ProcessAccounts) CreateAccount(ctx context.Context, account models.Accounts) (models.Accounts, error) {
	dbAccount := transformToDB(account)
	dbAccount, err := a.repositories.CreateAccount(ctx, dbAccount)
	if err != nil {
		return models.Accounts{}, err
	}
	return transform(dbAccount), nil
}

func (a ProcessAccounts) FindAccount(ctx context.Context, ID int) (models.Accounts, error) {
	dbAccount, err := a.repositories.GetAccountById(ctx, ID)
	if err != nil {
		return models.Accounts{}, err
	}
	return transform(dbAccount), nil
}

func (a ProcessAccounts) FindAccountUsers(ctx context.Context, ID int) ([]models.User, error) {
	users, err := a.repositories.GetAccountUsers(ctx, ID)
	if err != nil {
		return nil, nil
	}
	return transformToUserSlice(users), nil
}

func transform(dbAccount repository.Accounts) models.Accounts {
	account := models.Accounts{}
	account.ID = int(dbAccount.ID)
	account.Email = dbAccount.Email
	return account
}

func transformToDB(account models.Accounts) repository.Accounts {
	accountDb := repository.Accounts{}
	accountDb.ID = uint(account.ID)
	accountDb.Email = account.Email
	return accountDb
}

func transformToUserSlice(usersDb []repository.Users) []models.User {
	coll := []models.User{}
	for _, item := range usersDb {
		u := users.Transform(item)
		coll = append(coll, u)
	}
	return coll
}
