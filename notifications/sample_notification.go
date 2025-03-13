package notifications

import (
	"strconv"

	"github.com/OdairPianta/julia/models"
)

type SampleNotification struct {
	User        models.User
	SampleModel models.SampleModel
}

func (notification *SampleNotification) GetTitle() string {
	return "Notificação de exemplo " + strconv.FormatUint(uint64(notification.SampleModel.ID), 10)
}

func (notification *SampleNotification) GetContent() (string, error) {
	return "Está é uma notificação de exemplo referente ao model cadastrado: " + notification.SampleModel.SampleString, nil
}

func (notification *SampleNotification) GetNotifiable() models.User {
	return notification.User
}

func (notification *SampleNotification) GetNotifiers() []Notifier {
	return []Notifier{
		&MailNotifier{},
		&FcmNotifier{},
		&WhatsappNotifier{},
	}
}

func (notification *SampleNotification) Send() error {
	var notifiers = notification.GetNotifiers()
	for _, notifier := range notifiers {
		if err := notifier.Notify(notification); err != nil {
			return err
		}
	}
	return nil
}
