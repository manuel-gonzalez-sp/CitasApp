package persistence

import (
	"citasapp/internal/infra/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database = defaultDatabase()

func defaultDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("citasapp.db"), &gorm.Config{})
	if err != nil {
		utils.Logger.Fatalf("failed to connect to database: %v\n", err)
	}
	return db
}
