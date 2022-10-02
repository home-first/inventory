package database

import (
	"github.com/google/uuid"

	"github.com/home-first/inventory/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid primary_key unique" json:"id"`
	CreatedAt int       `gorm:"not null autoCreateTime" json:"createdAt"`
	UpdatedAt int       `gorm:"not null autoUpdateTime" json:"updatedAt"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID, err = uuid.NewRandom()
	}
	return err
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
