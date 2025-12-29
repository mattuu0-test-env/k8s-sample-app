package db

import (
	"log/slog"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := os.Getenv("DATABASE_URI")
	if dsn == "" {
		slog.Error("DATABASE_URI environment variable is not set")
		panic("DATABASE_URI environment variable is not set")
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		panic("failed to connect database")
	}
	slog.Info("connected to database")
}