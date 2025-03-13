package notifications

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type FcmNotifier struct{}

func (f *FcmNotifier) Notify(notification Notification) error {
	fcmContent, err := notification.GetContent()
	if err != nil {
		return err
	}
	opt := option.WithCredentialsFile("service_account_key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}
	fcmClient, err := app.Messaging(context.Background())

	if err != nil {
		return err
	}

	_, err = fcmClient.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: notification.GetTitle(),
			Body:  fcmContent,
		},
		Token: notification.GetNotifiable().FCMToken,
	})

	if err != nil {
		return err
	}

	fmt.Println("Enviando notificação FCM para o token:", notification.GetNotifiable().FCMToken)
	return nil
}
