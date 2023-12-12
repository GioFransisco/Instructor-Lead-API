package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type JwtConfig struct {
	Issuer        string
	SecretKey     []byte
	TokenLifeTime time.Duration
}

type ApiConfig struct {
	ApiPort string
}

type LogConfig struct {
	LogerFile string
}

type DbConfig struct {
	DbHost   string
	DbUser   string
	DbPass   string
	DbName   string
	DbPort   string
	DbDriver string
}

type Config struct {
	JwtConfig
	ApiConfig
	DbConfig
	LogConfig
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		DbHost:   os.Getenv("DB_HOST"),
		DbUser:   os.Getenv("DB_USER"),
		DbPass:   os.Getenv("DB_PASS"),
		DbName:   os.Getenv("DB_NAME"),
		DbPort:   os.Getenv("DB_PORT"),
		DbDriver: os.Getenv("DB_DRIVER"),
	}

	c.LogConfig = LogConfig{
		LogerFile: os.Getenv("LOG_FILE"),
	}

	tokenLifeTime, _ := strconv.Atoi(os.Getenv("TOKEN_LIFE_TIME"))

	c.JwtConfig = JwtConfig{
		Issuer:        os.Getenv("ISSUER_NAME"),
		SecretKey:     []byte(os.Getenv("SECRET_KEY")),
		TokenLifeTime: time.Duration(tokenLifeTime),
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	if c.ApiConfig.ApiPort == "" || c.DbConfig.DbHost == "" || c.DbConfig.DbName == "" || c.DbConfig.DbPort == "" || c.DbConfig.DbDriver == "" || c.JwtConfig.Issuer == "" || c.JwtConfig.SecretKey == nil || c.JwtConfig.TokenLifeTime == 0 {
		return errors.New("all environtment required")
	}

	return nil
}

func NewConfig() *Config {
	cfg := &Config{}

	if err := cfg.readConfig(); err != nil {
		panic(err)
	}

	return cfg
}
