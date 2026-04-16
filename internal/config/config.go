package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Port int
	DB   *gorm.DB
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	port := getEnvKey("PORT", "")
	dbHost := getEnvKey("DB_HOST", "")
	dbPort := getEnvKey("DB_PORT", "")
	dbUser := getEnvKey("DB_USERNAME", "")
	dbPassword := getEnvKey("DB_PASSWORD", "")
	dbName := getEnvKey("DB_NAME", "")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	db, err := loadDb(
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)

	return &Config{
		Port: portInt,
		DB:   db,
	}, nil
}

func getEnvKey(key string, fallbackValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallbackValue
}

func loadDb(dbHost string, dbPort string, dbUser string, dbPassword string, dbName string) (*gorm.DB, error) {
	pgDbPath := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(pgDbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil
}
