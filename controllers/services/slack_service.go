package services

type ISlackService interface {
	PostContactMessage(name, email, status, body string) error
}
