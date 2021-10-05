package test

import (
	"github.com/joho/godotenv"
	"github.com/moritiza/gift-code/discount-service/config"
	"github.com/sirupsen/logrus"
)

// Prepare for testing application
func Prepare() config.Config {
	var cfg config.Config

	cfg.Logger = logrus.New()
	cfg.Validator = config.Validator()

	err := godotenv.Load("../../.env")
	if err != nil {
		cfg.Logger.Fatalf("Could not load environments!\n%v", err)
	}

	cfg.Database = config.ConnectTestDatabase(cfg)
	return cfg
}
