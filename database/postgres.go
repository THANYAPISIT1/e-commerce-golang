package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"e-commerce/entities"
)

var PostgresDB *gorm.DB

func ConnectPostgres() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	PostgresDB = db
	return nil
}

func AutoMigrate() error {
	if PostgresDB == nil {
		return fmt.Errorf("PostgresDB is nil")
	}

	fmt.Println("Starting migration for User table...")

	err := PostgresDB.AutoMigrate(&entities.User{})
	if err != nil {
		fmt.Printf("Migration error: %v\n", err)
		return fmt.Errorf("failed to migrate User table: %w", err)
	}

	fmt.Println("Migration completed successfully!")
	return nil
}
