package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ase-lab-space/ase-lab-backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	cfg, err := config.New(".env")
	if err != nil {
		log.Fatal(err)
	}

	r.Run(fmt.Sprintf(":%d", cfg.PORT))
}
