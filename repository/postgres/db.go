package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Config struct {
		Username string
		Password string
		DBName   string
		Host     string
		Port     int
	}

	PostgresDB struct {
		config Config
		db     *gorm.DB
	}
)

func New(config Config) *PostgresDB {
	var db *gorm.DB
	var dialector gorm.Dialector

	// Setup the connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Host,
		config.Username,
		config.Password,
		config.DBName,
		config.Port)
	dialector = postgres.Open(dsn)

	db, err := gorm.Open(dialector)
	if err != nil {
		// if we cannot connect to the database, the application basically has no functionality
		panic("could not connect to the database: " + err.Error())
	}

	// Use automigration to keep the tables up to date
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("could not migrate: " + err.Error())
	}

	return &PostgresDB{
		config: config,
		db:     db,
	}
}
