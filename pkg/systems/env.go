package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBhost     string
	DBport     string
	DBuser     string
	DBpassword string
	DBname     string
	DBsslmode  string
}

func LoadDBConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка чтения .env")
	}

	return &DBConfig{
		DBhost:     os.Getenv("HOST_IP"),
		DBport:     os.Getenv("DB_PORT"),
		DBuser:     os.Getenv("DB_USERNAME"),
		DBpassword: os.Getenv("DB_PASSWORD"),
		DBname:     os.Getenv("DB_NAME"),
		DBsslmode:  os.Getenv("DB_SSLMODE"),
	}
}

func GetEnvString(param string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(param)
}
