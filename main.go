package main

import (
	"fmt"
	"log"

	"github.com/ase-lab-space/ase-lab-backend/config"
	"github.com/ase-lab-space/ase-lab-backend/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.New(".env")
	if err != nil {
		log.Fatal(err)
	}
	gin.SetMode(cfg.GIN_MODE)

	r := routers.New(cfg)
	r.Run(fmt.Sprintf(":%d", cfg.PORT))
}
