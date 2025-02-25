package database

import (
	"fmt"
	"linker/internal/migration"
	"linker/internal/seeder"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := migration.Migrate(db); err != nil {
		return nil, err
	}

	if err := seeder.Seed(db); err != nil {
		return nil, err
	}

    DB = db
	
	return db, nil
}

func GetDB() *gorm.DB {
    return DB
}
