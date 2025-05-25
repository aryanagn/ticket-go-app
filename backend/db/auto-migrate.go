package db

import (
	"github.com/aryanagn/ticket-go-app/backend/models"
	
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{}, &models.Ticket{}, &models.User{})
}
