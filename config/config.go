package config

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Config struct {
	PORT                                  int
	GIN_MODE                              string
	CONTACT_FORM_NOTIFICATOR_ACCESS_TOKEN string
	ALLOWED_ORIGINS                       []string
}

func (c *Config) Setup() error {
	gin.SetMode(c.GIN_MODE)
	return nil
}

func New(filename string) (*Config, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}

	godotenv.Load(filename)

	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	return &Config{
		PORT:                                  PORT,
		GIN_MODE:                              os.Getenv("GIN_MODE"),
		CONTACT_FORM_NOTIFICATOR_ACCESS_TOKEN: os.Getenv("CONTACT_FORM_NOTIFICATOR_ACCESS_TOKEN"),
		ALLOWED_ORIGINS: []string{
			"http://localhost:8080",
			"https://ase-lab.space",
			"https://ase.cmd-moon.space",
		},
	}, nil
}
