package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Não foi encontrado nenhum arquivo .env")
	}
}

func GetDBConnectionString() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dbTimezone := os.Getenv("DB_TIMEZONE")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode, dbTimezone)
}