package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func Load() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf(".env file not found: %v", err)
	}

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("API_PORT"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("DATABASE_HOST"),
		Port:     viper.GetString("DATABASE_PORT"),
		User:     viper.GetString("DATABASE_USER"),
		Pass:     viper.GetString("DATABASE_PASS"),
		Database: viper.GetString("DATABASE_NAME"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
