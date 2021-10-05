package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Config store general configs
type Config struct {
	Logger       *logrus.Logger
	Database     *gorm.DB
	TestDatabase *gorm.DB
	Validator    *validator.Validate
}
