package main

import (
	"fmt"
	"log"

	"github.com/ase-lab-space/ase-lab-backend/config"
	"github.com/ase-lab-space/ase-lab-backend/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	envFile := ".env"
	log.Printf("Loading %s", envFile)
	cfg, err := config.New(envFile)
	if err != nil {
		log.Fatal(err)
	}
	gin.SetMode(cfg.GIN_MODE)

	r := routers.New(cfg)
	log.Printf("Running on %d", cfg.PORT)
	r.Run(fmt.Sprintf(":%d", cfg.PORT))
}
