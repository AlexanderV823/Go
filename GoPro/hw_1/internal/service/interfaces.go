package service

import "hw_1/internal/domain"

// OrderWriter определяет контракт для сохранения данных
type OrderWriter interface {
	Save(order *domain.Order) error
}

// DBInitializer определяет контракт для настройки схемы
type DBInitializer interface {
	InitSchema() error
}

// MessageSender определяет контракт для отправки уведомлений
type MessageSender interface {
	Send(recipient string, message string) error
}
