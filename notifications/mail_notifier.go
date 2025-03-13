package notifications

import (
	"fmt"
	"net/smtp"
	"os"
)

type MailNotifier struct{}

func (m *MailNotifier) Notify(notification Notification) error {
	var mailHost = os.Getenv("MAIL_HOST")
	var mailPort = os.Getenv("MAIL_PORT")
	var mailUsername = os.Getenv("MAIL_USERNAME")
	var mailPassword = os.Getenv("MAIL_PASSWORD")
	// var mailEncryption = os.Getenv("MAIL_ENCRYPTION")
	var mailFromAddress = os.Getenv("MAIL_FROM_ADDRESS")
	// var mailFromName = os.Getenv("MAIL_FROM_NAME")
	// var mailCharset = os.Getenv("MAIL_CHARSET")
	// var mailAutoTls = os.Getenv("MAIL_AUTO_TLS")
	var appName = os.Getenv("APP_NAME")

	var mailContent, err = notification.GetContent()
	if err != nil {
		return err
	}

	var completeMessage = []byte("From: " + appName + " <" + mailFromAddress + ">\r\n" +
		"To: " + notification.GetNotifiable().Email + "\r\n" +
		"Subject: " + notification.GetTitle() + "\r\n" +
		"MIME-version: 1.0;" + "\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";" + "\r\n" +
		mailContent + "\r\n")

	var bytesMessage = []byte(completeMessage)

	auth := smtp.PlainAuth("", mailUsername, mailPassword, mailHost)

	// Send actual message
	err = smtp.SendMail(mailHost+":"+mailPort, auth, mailFromAddress, []string{notification.GetNotifiable().Email}, bytesMessage)
	if err != nil {
		return err
	}
	fmt.Println("Enviando e-mail com conteúdo:", notification.GetNotifiable().Email)
	return nil
}
