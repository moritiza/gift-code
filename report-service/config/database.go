package config

import (
	"os"

	"github.com/moritiza/gift-code/report-service/app/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to PostgreSQL
func ConnectDatabase(cfg Config) *gorm.DB {
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSL_MODE") +
		" TimeZone=" + os.Getenv("DB_TIME_ZONE")

	return connector(cfg, dsn)
}

// Connect to test PostgreSQL
func ConnectTestDatabase(cfg Config) *gorm.DB {
	dsn := "host=" + os.Getenv("TEST_DB_HOST") +
		" user=" + os.Getenv("TEST_DB_USER") +
		" password=" + os.Getenv("TEST_DB_PASSWORD") +
		" dbname=" + os.Getenv("TEST_DB_NAME") +
		" port=" + os.Getenv("TEST_DB_PORT") +
		" sslmode=" + os.Getenv("TEST_DB_SSL_MODE") +
		" TimeZone=" + os.Getenv("TEST_DB_TIME_ZONE")

	return connector(cfg, dsn)
}

// Connector for PostgreSQL
func connector(cfg Config, dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		cfg.Logger.Fatalf("Could not connect to the database!\n%v", err)
	}

	db.AutoMigrate(entity.Report{})

	return db
}

// Disconnect PostgreSQL connection
func DisconnectDatabase(cfg Config) {
	closer, err := cfg.Database.DB()
	if err != nil {
		cfg.Logger.Fatalf("Could not disconnect the database!\n%v", err)
	}

	closer.Close()
}
