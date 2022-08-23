package services

import (
	"fmt"

	"github.com/ase-lab-space/ase-lab-backend/controllers/services"
	"github.com/ase-lab-space/ase-lab-backend/services/repositories"
)

type SlackService struct {
	repository repositories.ISlackRepository
}

func (s *SlackService) PostContactMessage(name, email, status, body string) error {
	message := fmt.Sprintf(
		`<!channel>
:raising_hand: ASE-Lab. Contactフォームからご連絡をいただきました :raising_hand:

*お名前*
%s

*メールアドレス*
%s

*ご職業*
%s

*本文*
%s
`, name, email, status, body,
	)
	if err := s.repository.SendMessage(repositories.FormNotification, message); err != nil {
		return err
	}

	return nil
}

func NewSlackService(repository repositories.ISlackRepository) services.ISlackService {
	return &SlackService{
		repository: repository,
	}
}
