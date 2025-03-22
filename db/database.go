package db

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/kwangminnam/rn-with-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig, DBMigrate func(db *gorm.DB) error) *gorm.DB {
	uri := fmt.Sprintf(
		`host=%s user=%s password=%s dbname=%s port=5432 sslmode=%s`,
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Info("Connected to database")

	if err := DBMigrate(db); err != nil {
		log.Fatalf("unable to migrate database: %v", err)
	}

	return db
}
