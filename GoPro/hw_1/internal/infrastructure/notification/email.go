package notification

import "fmt"

// EmailSender — отправляет e-mail сообщения
type EmailSender struct{}

func NewEmailService() *EmailSender {
	return &EmailSender{}
}

func (e *EmailSender) Send(recipient string, message string) error {
	fmt.Printf("Сообщение e-mail отправлено для %s: %s\n", recipient, message)
	return nil
}
