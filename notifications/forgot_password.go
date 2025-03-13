package notifications

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/OdairPianta/julia/models"
)

type ForgotPasswordNotification struct {
	User models.User
}

func (notification *ForgotPasswordNotification) GetNotifiable() models.User {
	return notification.User
}

func (notification *ForgotPasswordNotification) GetTitle() string {
	return "Recuperação de senha"
}

func (notification *ForgotPasswordNotification) GetContent() (string, error) {
	var appName = os.Getenv("APP_NAME")
	var passwordRecoveyUrl = os.Getenv("APP_URL") + "/password_recovery/" + fmt.Sprint(notification.User.ResetPasswordToken)

	var body string
	t, err := template.ParseFiles("resources/templates/forgot_password.html")
	if err != nil {
		return "", err
	}

	templateData := struct {
		Name    string
		URL     string
		AppName string
	}{
		Name:    notification.User.Name,
		URL:     passwordRecoveyUrl,
		AppName: appName,
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, templateData); err != nil {
		return "", err
	}
	body = buf.String()
	return body, nil
}

func (notification *ForgotPasswordNotification) GetNotifiers() []Notifier {
	return []Notifier{
		&MailNotifier{},
		&FcmNotifier{},
		&WhatsappNotifier{},
	}
}

func (notification *ForgotPasswordNotification) Send() error {
	var notifiers = notification.GetNotifiers()
	for _, notifier := range notifiers {
		if err := notifier.Notify(notification); err != nil {
			return err
		}
	}
	return nil
}
