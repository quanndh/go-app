package db

import (
	"fmt"
	"github.com/quanndh/go-app/adapter/models"
	"github.com/quanndh/go-app/public/config"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectDBIn struct {
	fx.In
	Config *config.Configuration
}

func ConnectDB(c ConnectDBIn) *gorm.DB {
	cfg := c.Config.App.Datasource
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port)
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
