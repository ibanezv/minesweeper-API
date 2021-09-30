package main

import (
	"github/ibanezv/minesweeper-API/cmd/api/handlers"
	"github/ibanezv/minesweeper-API/cmd/api/settings"
	"github/ibanezv/minesweeper-API/internal/accounts"
	"github/ibanezv/minesweeper-API/internal/distributions"
	"github/ibanezv/minesweeper-API/internal/games"
	"github/ibanezv/minesweeper-API/internal/repository"
	"github/ibanezv/minesweeper-API/internal/users"
	"github/ibanezv/minesweeper-API/pkg/database"
	"net/http"
)

func main() {

	configs := settings.LoadConfigurationDB()
	db := database.NewDatabase(configs)
	dbConnection, _ := db.GetConnection()
	minesweeperDao := repository.NewRepository(dbConnection)

	serviceDistributions := distributions.NewService(minesweeperDao)
	serviceGames := games.NewService(minesweeperDao, serviceDistributions)
	serviceUsers := users.NewService(minesweeperDao)
	serviceAccounts := accounts.NewService(minesweeperDao)

	router := handlers.ApiRoutesMapper(serviceGames, serviceDistributions, serviceUsers, serviceAccounts)
	http.Handle("/", router)
	_ = http.ListenAndServe(":8090", nil)
}
