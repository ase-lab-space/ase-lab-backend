package routers

import (
	"github.com/ase-lab-space/ase-lab-backend/controllers"
	"github.com/gin-gonic/gin"
)

type SlackRoute struct {
	gin        *gin.Engine
	controller *controllers.SlackController
}

func (r *SlackRoute) Setup() {
	slack := r.gin.Group("/slack")
	{
		slack.POST("/contact", r.controller.SubmitContact)
	}
}

func NewSlackRoute(engine *gin.Engine, controller *controllers.SlackController) *SlackRoute {
	return &SlackRoute{
		gin:        engine,
		controller: controller,
	}
}
