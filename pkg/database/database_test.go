package database

import (
	"github/ibanezv/minesweeper-API/cmd/api/settings"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestConfigureDatabaseConnection(t *testing.T) {
	// given
	tests := []struct {
		testName      string
		db            func(connectionConfig string, opts gorm.Option) (db *gorm.DB, err error)
		config        settings.Database
		expectedError bool
	}{
		{
			testName: "Configure database with not empty values",
			db: func(connectionConfig string, opts gorm.Option) (db *gorm.DB, err error) {
				dbSqlite, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
					SkipDefaultTransaction: true,
				})
				return dbSqlite, err
			},
			config: settings.Database{
				Username:        "root",
				Password:        "passroot",
				Host:            "localhost",
				DatabaseName:    "",
				MaxOpenConns:    20,
				MaxIdleConns:    10,
				ConnMaxLifetime: 30000,
				ConnMaxIdleTime: 30000,
				Timeout:         1500,
				ReadTimeout:     1000,
				WriteTimeout:    1000,
			},
			expectedError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			db := NewDatabase(tt.config)
			db.openConnection = tt.db
			// when
			_, err := db.GetConnection()

			// then
			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
