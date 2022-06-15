package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"investidea.tech.test/internal/entities"
	"investidea.tech.test/internal/services/database"
	"investidea.tech.test/internal/services/log"
	"investidea.tech.test/pkg/config"
	"logur.dev/logur"
)

var AppName = "DB Migrator"

type MigrateService struct {
	logger logur.LoggerFacade
	db     *database.DB
}

func main() {
	migrateService := MigrateService{}
	migrateService.Init()

	tables := []interface{}{
		entities.Buyer{},
		entities.Seller{},
		entities.Product{},
		entities.Order{},
	}

	err := migrateService.db.GormDB().AutoMigrate(tables...)
	if err != nil {
		migrateService.logger.Error(fmt.Sprintf("Seed failed, details: %v", err))
		return
	}

	migrateService.logger.Info("Seed completed")
}

func (s *MigrateService) Init() {
	v, f := viper.New(), pflag.NewFlagSet(AppName, pflag.ExitOnError)
	cfg := config.New(v, f)
	logger := log.NewLogger(cfg.Log)

	// Override the global standard library logger to make sure everything uses our logger
	log.SetStandardLogger(logger)
	// Start database
	db := database.New(logger, cfg.Database)

	s.db = db
	s.logger = logger
}
