package configs

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	ServerPort string
}

func Connect() (*gorm.DB, error) {
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		POSTGRES_HOST, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}


func LoadConfig() Config {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}
	if port[0] != ':' {
		port = ":" + port
	}
	
	return Config{
		ServerPort: port,
	}
}