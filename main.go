package main

import (
	"fmt"
	"log"

	"github.com/ase-lab-space/ase-lab-backend/config"
	"github.com/ase-lab-space/ase-lab-backend/middlewares"
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
	if err := cfg.Setup(); err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()

	middlewares := middlewares.NewMiddlewares(cfg, engine)
	if err := middlewares.Setup(); err != nil {
		log.Fatal(err)
	}

	routes := routers.NewRoutes(engine, cfg)
	routes.Setup()

	log.Printf("Running on %d", cfg.PORT)
	engine.Run(fmt.Sprintf(":%d", cfg.PORT))
}
