package database

import (
	"github.com/home-first/inventory/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type BaseModel struct {
	ID        uint `gorm:"primary_key autoIncrement unique" json:"id"`
	CreatedAt int  `gorm:"not null autoCreateTime" json:"createdAt"`
	UpdatedAt int  `gorm:"not null autoUpdateTime" json:"updatedAt"`
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open(config.Cfg.DatabaseFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(
		&Item{},
		&Collection{},
		&User{},
	)

	DB = database
}
