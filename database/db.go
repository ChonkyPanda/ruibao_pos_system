package database

import (
	"fmt"
	"log"
	"os"

	"ruibao_pos_system/models"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	if err := godotenv.Load("secret.env"); err != nil {
		log.Println("Warning: No .env file found, reading from system environment")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// 1. Connect to default 'postgres' database to ensure target DB exists
	systemDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable",
		host, user, password, port,
	)

	sysDB, err := gorm.Open(postgres.Open(systemDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to system database: %v", err)
	}

	// 2. Create the target database if it doesn't exist
	checkQuery := fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", dbName)
	var exists int
	sysDB.Raw(checkQuery).Scan(&exists)

	if exists == 0 {
		createQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
		if err := sysDB.Exec(createQuery).Error; err != nil {
			log.Fatalf("Failed to create database %s: %v", dbName, err)
		}
		log.Printf("Database '%s' created successfully!", dbName)
	}

	// Close system connection
	sqlSysDB, _ := sysDB.DB()
	sqlSysDB.Close()

	// 3. Connect to your actual target database
	targetDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, port,
	)

	DB, err = gorm.Open(postgres.Open(targetDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL target DB: %v", err)
	}

	// 4. Auto-migrate models
	err = DB.AutoMigrate(
		&models.User{},
		&models.Product{},
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate schema: %v", err)
	}

	fmt.Println("Database connection established successfully!")
}