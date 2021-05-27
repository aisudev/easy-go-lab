package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLiteDB : Centralize database instance connection
var SQLiteDB *gorm.DB

// InitialDatabase : Initial database function to connect with DBMS
func InitialDatabase() error {
	var err error
	SQLiteDB, err = gorm.Open(sqlite.Open("database/lab.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	autoMigrateDatabase()

	fmt.Println("Database connection successfully...")
	return nil
}

func autoMigrateDatabase() {
	SQLiteDB.AutoMigrate(&StudentScore{})
}
