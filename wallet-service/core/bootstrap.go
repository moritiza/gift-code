package core

import (
	"github.com/joho/godotenv"
	"github.com/moritiza/gift-code/wallet-service/config"
	"github.com/sirupsen/logrus"
)

func Bootstrap() config.Config {
	var cfg config.Config

	// Create new logrus
	cfg.Logger = logrus.New()

	// Create new validator
	cfg.Validator = config.Validator()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		cfg.Logger.Fatalf("Could not load environments!\n%v", err)
	}

	// Connect to database
	cfg.Database = config.ConnectDatabase(cfg)
	return cfg
}
