package notifications

import (
	"fmt"
)

type WhatsappNotifier struct{}

func (w *WhatsappNotifier) Notify(notification Notification) error {
	content, err := notification.GetContent()
	if err != nil {
		return err
	} // lógica para enviar mensagem WhatsApp com o conteúdo
	fmt.Println("Enviando mensagem WhatsApp com conteúdo:", content)
	return nil
}
