package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type ServerConfig struct {
	Address string
	Debug   bool
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func (pgConfig *PostgresConfig) GetURL() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pgConfig.Host, pgConfig.Port, pgConfig.User, pgConfig.Password, pgConfig.Name)
}

func Server() ServerConfig {
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatal(err)
	}

	return ServerConfig{
		Address: os.Getenv("API_URL"),
		Debug:   debug,
	}
}

func Database() PostgresConfig {
	return PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
}
