package routers

import (
	"github.com/ase-lab-space/ase-lab-backend/config"
	"github.com/ase-lab-space/ase-lab-backend/controllers"
	"github.com/ase-lab-space/ase-lab-backend/repositories"
	"github.com/ase-lab-space/ase-lab-backend/services"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Setup()
}

type Routes []Route

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}

func NewRoutes(engine *gin.Engine, cfg *config.Config) Routes {
	// SlackRoute
	slackRepository := repositories.NewSlackRepository(cfg)
	slackService := services.NewSlackService(slackRepository)
	slackController := controllers.NewSlackController(slackService)
	slackRoute := NewSlackRoute(engine, slackController)

	return Routes{
		slackRoute,
	}
}

func New(cfg *config.Config) *gin.Engine {
	engine := gin.Default()

	routes := NewRoutes(engine, cfg)
	routes.Setup()

	return engine
}
