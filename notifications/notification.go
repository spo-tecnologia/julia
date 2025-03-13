package notifications

import "github.com/OdairPianta/julia/models"

type Notification interface {
	GetNotifiable() models.User
	GetTitle() string
	GetContent() (string, error)
	GetNotifiers() []Notifier
	Send() error
}
