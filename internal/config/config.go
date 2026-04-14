package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	port int
	db   *sql.DB
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
		port: portInt,
		db:   db,
	}, nil
}

func getEnvKey(key string, fallbackValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallbackValue
}

func loadDb(dbHost string, dbPort string, dbUser string, dbPassword string, dbName string) (*sql.DB, error) {
	pgDbPath := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", pgDbPath)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db, nil
}
