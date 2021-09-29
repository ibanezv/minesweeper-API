package users

import (
	"context"
	"github/ibanezv/minesweeper-API/internal/games"
	"github/ibanezv/minesweeper-API/internal/models"
	"github/ibanezv/minesweeper-API/internal/repository"
	"github/ibanezv/minesweeper-API/pkg/database"
)

type Service interface {
	FindUser(context.Context, int) (models.User, error)
	CreateUser(context.Context, models.User) (models.User, error)
	FindUserGames(ctx context.Context, ID int) ([]models.Game, error)
}

type ProcessUser struct {
	repositories database.Repositories
}

func NewService(repositories database.Repositories) *ProcessUser {
	return &ProcessUser{repositories}
}

func (u *ProcessUser) FindUser(ctx context.Context, ID int) (models.User, error) {
	user, err := u.repositories.GetUserById(ctx, ID)
	if err != nil {
		return models.User{}, nil
	}
	return Transform(user), nil
}

func (u *ProcessUser) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	dbUser := transformToDB(user)
	newUser, err := u.repositories.CreateUser(ctx, dbUser)
	if err != nil {
		return models.User{}, err
	}
	return Transform(newUser), nil
}

func (u *ProcessUser) FindUserGames(ctx context.Context, ID int) ([]models.Game, error) {
	games, err := u.repositories.GetUserGames(ctx, ID)
	if err != nil {
		return nil, err
	}
	return transformSlice(games), nil
}

func Transform(userDB repository.Users) models.User {
	user := models.User{}
	user.ID = int(userDB.ID)
	user.NickName = userDB.NickName
	user.AccountID = userDB.AccountID
	return user
}

func transformToDB(user models.User) repository.Users {
	userDb := repository.Users{}
	userDb.ID = uint(user.ID)
	userDb.NickName = user.NickName
	userDb.AccountID = user.AccountID
	return userDb
}

func transformSlice(dbGames []repository.Games) []models.Game {
	coll := []models.Game{}
	for _, item := range dbGames {
		g := games.Transform(item)
		coll = append(coll, g)
	}
	return coll
}
