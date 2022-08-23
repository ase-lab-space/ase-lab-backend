package repositories_test

import (
	"testing"

	"github.com/ase-lab-space/ase-lab-backend/config"
	"github.com/ase-lab-space/ase-lab-backend/repositories"
	"github.com/ase-lab-space/ase-lab-backend/repositories/gen"
	irepositories "github.com/ase-lab-space/ase-lab-backend/services/repositories"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
)

func TestSendMessage(t *testing.T) {
	cfg := config.NewConfigMock()
	r := repositories.NewSlackRepository(cfg)
	slackApiMock := &gen.ISlackClientMock{
		PostMessageFunc: func(s string, msgOptions ...slack.MsgOption) (string, string, error) {
			return "", "", nil
		},
	}
	repositories.NewSlackApi = func(_ string) repositories.ISlackClient {
		return slackApiMock
	}

	if err := r.SendMessage(irepositories.FormNotification, "some body"); err != nil {
		t.Errorf("error occured: %v", err)
	}
	assert.Equal(t, len(slackApiMock.PostMessageCalls()), 1)
	call := slackApiMock.PostMessageCalls()[0]
	assert.EqualValues(t, call.S, irepositories.FormNotification)
}
