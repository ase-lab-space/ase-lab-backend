package controllers

import (
	"net/http"

	"github.com/ase-lab-space/ase-lab-backend/controllers/services"
	"github.com/gin-gonic/gin"
)

type SlackController struct {
	service services.ISlackService
}

func NewSlackController(service services.ISlackService) *SlackController {
	return &SlackController{
		service: service,
	}
}

type SubmitContactRequest struct {
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Status string `json:"status" binding:"required"`
	Body   string `json:"body" binding:"required"`
}

// request to slack api to notify
func (s *SlackController) SubmitContact(c *gin.Context) {
	var req SubmitContactRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	if err := s.service.PostContactMessage(req.Name, req.Email, req.Status, req.Body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to post a message"})
		return
	}

	c.JSON(http.StatusOK, nil)
}
