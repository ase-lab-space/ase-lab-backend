package middlewares

import (
	"github.com/ase-lab-space/ase-lab-backend/config"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	Setup() error
}

type Middlewares []Middleware

func (middlewares *Middlewares) Setup() error {
	for _, middleware := range *middlewares {
		err := middleware.Setup()
		if err != nil {
			return err
		}
	}

	return nil
}

func NewMiddlewares(cfg *config.Config, engine *gin.Engine) Middlewares {
	return Middlewares{
		NewCorsMiddleware(cfg, engine),
	}
}
