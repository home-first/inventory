package database

import (
	"github.com/google/uuid"

	"github.com/home-first/inventory/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/home-first/inventory/logger"
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
	var gormLogger = logger.GormLogger{}
	database, err := gorm.Open(sqlite.Open(config.Cfg.DatabaseFile), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		panic(err)
	}
	log.Info().Msg("Successfully connected to database")
	log.Info().Msg("Auto migrating database")
	database.AutoMigrate(
		&Item{},
		&Collection{},
		&User{},
	)

	DB = database
}
