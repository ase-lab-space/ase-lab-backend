package services

import (
	"testing"

	"github.com/ase-lab-space/ase-lab-backend/services/repositories"
	"github.com/ase-lab-space/ase-lab-backend/services/repositories/gen"
	"github.com/stretchr/testify/assert"
)

func TestPostContactMessage(t *testing.T) {
	repositoryMock := &gen.ISlackRepositoryMock{
		SendMessageFunc: func(channel repositories.Channel, body string) error {
			return nil
		},
	}

	s := NewSlackService(repositoryMock)

	err := s.PostContactMessage("John Doe", "test@example.com", "大学生", "これはテストです")

	if err != nil {
		t.Errorf("error occurred: %v", err)
	}

	assert.Equal(t, len(repositoryMock.SendMessageCalls()), 1)
	assert.Equal(t, repositoryMock.SendMessageCalls()[0].Channel, repositories.FormNotification)
	assert.Equal(t, repositoryMock.SendMessageCalls()[0].Body,
		`<!channel>
:raising_hand: ASE-Lab. Contactフォームからご連絡をいただきました :raising_hand:

*お名前*
John Doe

*メールアドレス*
test@example.com

*ご職業*
大学生

*本文*
これはテストです
`)
}
