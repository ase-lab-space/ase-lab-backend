package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT int
}

func New(filename string) (*Config, error) {
	godotenv.Load(filename)

	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	return &Config{
		PORT: PORT,
	}, nil
}
