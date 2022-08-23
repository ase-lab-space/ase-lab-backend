package repositories

import (
	"github.com/ase-lab-space/ase-lab-backend/config"
	"github.com/ase-lab-space/ase-lab-backend/services/repositories"
	"github.com/slack-go/slack"
)

//go:generate moq -out gen/slack_repository_moq.go -pkg gen . ISlackClient
type ISlackClient interface {
	PostMessage(string, ...slack.MsgOption) (string, string, error)
}

var NewSlackApi = func(accessToken string) ISlackClient {
	return slack.New(accessToken)
}

type SlackRepository struct {
	cfg *config.Config
}

func (s *SlackRepository) SendMessage(channel repositories.Channel, body string) error {
	api := NewSlackApi(s.cfg.CONTACT_FORM_NOTIFICATOR_ACCESS_TOKEN)
	_, _, err := api.PostMessage(string(channel), slack.MsgOptionText(body, true))
	if err != nil {
		return err
	}

	return nil
}

func NewSlackRepository(cfg *config.Config) *SlackRepository {
	return &SlackRepository{
		cfg: cfg,
	}
}
