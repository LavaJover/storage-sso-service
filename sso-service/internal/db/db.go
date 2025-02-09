package db

import (
	"log"

	"github.com/LavaJover/storage-sso-service/sso-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initialize database and make neccessary migrations
func InitDB(dsn string) *gorm.DB{
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatalf("failed to init db: %v\n", err)
	}

	err = db.AutoMigrate(&models.User{})

	if err != nil{
		log.Fatal("failed to migrate User model", "msg", err.Error())
	}

	return db
}