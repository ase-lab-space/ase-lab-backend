package repositories

//go:generate moq -out gen/slack_repository_moq.go -pkg gen . ISlackRepository

type Channel string

const (
	FormNotification Channel = "#0_form_notifications"
)

type ISlackRepository interface {
	SendMessage(channel Channel, body string) error
}
