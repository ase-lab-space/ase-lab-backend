package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT     int
	GIN_MODE string
}

func New(filename string) (*Config, error) {
	godotenv.Load(filename)

	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	return &Config{
		PORT:     PORT,
		GIN_MODE: os.Getenv("GIN_MODE"),
	}, nil
}
