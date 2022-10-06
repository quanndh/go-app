package db

import (
	"github.com/quanndh/go-app/adapter/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=123456 dbname=go port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		return nil
	}

	return db
}
