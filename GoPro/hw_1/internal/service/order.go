package service

import (
	"fmt"
	"log"
	"hw_1/internal/domain"
)

// OrderService координирует бизнес-процесс создания заказа
type OrderService struct {
	repo     OrderWriter
	notifier MessageSender
}

func NewOrderService(repo OrderWriter, notifier MessageSender) *OrderService {
	return &OrderService{
		repo:     repo,
		notifier: notifier,
	}
}

func (s *OrderService) CreateOrder(customer string, products []string, total float64) error {
	order := &domain.Order{
		Customer: customer,
		Products: products,
		Total:    total,
		Status:   "pending",
	}

	// Сохранение
	if err := s.repo.Save(order); err != nil {
		return fmt.Errorf("ошибка сохранения заказа: %w", err)
	}

	// Отправка
	msg := fmt.Sprintf("Заказ на сумму %.2f успешно оформлен!", total)
	if err := s.notifier.Send(customer, msg); err != nil {
		log.Printf("Внимание: не удалось отправить уведомление: %v", err)
	}

	return nil
}
