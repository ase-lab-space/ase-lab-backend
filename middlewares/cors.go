package middlewares

import (
	"github.com/ase-lab-space/ase-lab-backend/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CorsMiddleware struct {
	cfg    *config.Config
	engine *gin.Engine
}

func (m *CorsMiddleware) Setup() error {
	config := cors.DefaultConfig()
	config.AllowOrigins = m.cfg.ALLOWED_ORIGINS
	m.engine.Use(cors.New(config))

	return nil
}

func NewCorsMiddleware(cfg *config.Config, engine *gin.Engine) Middleware {
	return &CorsMiddleware{
		cfg:    cfg,
		engine: engine,
	}
}
