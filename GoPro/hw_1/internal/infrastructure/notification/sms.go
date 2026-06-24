package notification

import "fmt"

// SMSSender - реализует интерфейс service.Notifier
type SMSSender struct{}

func NewSMSSender() *SMSSender {
	return &SMSSender{}
}

func (s *SMSSender) Send(recipient string, message string) error {
	fmt.Printf("SMS-cообщение отправлено на номер %s: %s\n", recipient, message)
	return nil
}
