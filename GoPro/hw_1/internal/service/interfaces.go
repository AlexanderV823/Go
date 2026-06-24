package service

import "hw_1/internal/domain"

// RepositoryWriter определяет контракт для сохранения данных
type RepositoryWriter interface {
	Save(order *domain.Order) error
}

// DBInitializer определяет контракт для настройки схемы
type DBInitializer interface {
	InitSchema() error
}

// Notifier определяет контракт для отправки уведомлений
type Notifier interface {
	Send(recipient string, message string) error
}
