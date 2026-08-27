package service_test

import (
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"errors"
	"testing"

	"github.com/lib/pq"
)

type stubUserRepo struct {
	service.UserRepository
	dbError error
}

func (s *stubUserRepo) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return false, nil
}

func (s *stubUserRepo) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	return false, nil
}

func (s *stubUserRepo) Create(ctx context.Context, user *model.User) error {
	return s.dbError
}

func TestUserService_Register_RaceConditionUniqueViolation(t *testing.T) {
	// Подготовка: симулируем ошибку СУБД unique_violation (код 23505)
	pgErr := &pq.Error{
		Code: "23505",
	}
	repo := &stubUserRepo{dbError: pgErr}

	// Передаем nil вместо jwtManager, до генерации токена выполнение не дойдет
	svc := service.NewUserService(repo, nil)

	req := &model.UserCreateRequest{
		Username: "concurrent_user",
		Email:    "race@test.com",
		Password: "SecurePassword123!",
	}

	// Действие
	_, err := svc.Register(context.Background(), req)

	// Проверка
	if !errors.Is(err, service.ErrUserAlreadyExists) {
		t.Fatalf("ожидался маппинг гонки СУБД в ошибку %v, получено: %v", service.ErrUserAlreadyExists, err)
	}
}

func TestUserService_Register_InvalidEmailFormat(t *testing.T) {
	// Подготовка
	svc := service.NewUserService(nil, nil)
	req := &model.UserCreateRequest{
		Username: "valid_username",
		Email:    "broken-email-without-at",
		Password: "SecurePassword123!",
	}

	// Действие
	_, err := svc.Register(context.Background(), req)

	// Проверка
	if err == nil {
		t.Fatal("ожидался отказ в регистрации из-за некорректного формата email")
	}
}
