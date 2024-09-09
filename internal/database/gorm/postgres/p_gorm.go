package postgres

import (
	"fmt"

	"github.com/madmuzz05/golang-assigment-2/internal/config"
	"github.com/madmuzz05/golang-assigment-2/service/entity"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDB struct {
	DB  *gorm.DB
	Trx *gorm.DB
}

func LoadGorm(cfg *config.Config) (*GormDB, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseUsername, cfg.DatabasePassword, cfg.DatabaseName, cfg.DatabaseSSL)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}
	db.Debug().AutoMigrate(entity.Order{}, entity.Item{})
	log.Info().Msg("connected successfully to the database with gorm!")

	return &GormDB{
		DB: db,
	}, nil
}
func (g *GormDB) BeginTransaction() {
	g.Trx = g.DB.Begin()
}

func (g *GormDB) CommitTransaction() {
	g.Trx.Commit()
	g.Trx = nil
}

func (g *GormDB) RollbackTransaction() {
	g.Trx.Rollback()
	g.Trx = nil
}

func (g *GormDB) GetDB() *gorm.DB {
	return g.DB
}
