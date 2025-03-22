package db

import (
	"github.com/kwangminnam/rn-with-go/models"
	"gorm.io/gorm"
)

func DBMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{})

}
