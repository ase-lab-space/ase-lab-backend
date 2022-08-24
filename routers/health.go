package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthRoute struct {
	gin *gin.Engine
}

func (r *HealthRoute) Setup() {
	r.gin.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
}

func NewHealthRoute(engine *gin.Engine) *HealthRoute {
	return &HealthRoute{
		gin: engine,
	}
}
