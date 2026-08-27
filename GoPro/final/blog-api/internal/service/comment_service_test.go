package service_test

import (
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"errors"
	"testing"
)

// Компиляция заглушек гарантирует соответствие контракту интерфейсов
var _ service.PostRepository = (*stubCommentPostRepo)(nil)

type stubCommentPostRepo struct {
	service.PostRepository
	existsResult bool
}

func (s *stubCommentPostRepo) Exists(ctx context.Context, id int) (bool, error) {
	return s.existsResult, nil
}

func TestCommentService_Create_PostNotExists(t *testing.T) {
	// Подготовка: Пост не существует в БД
	postRepo := &stubCommentPostRepo{existsResult: false}
	svc := service.NewCommentService(nil, postRepo, nil)

	req := &model.CommentCreateRequest{Content: "Привет, отличный пост!"}

	// Действие
	_, err := svc.Create(context.Background(), 999, req, 1)

	// Проверка
	if !errors.Is(err, service.ErrPostNotExists) {
		t.Fatalf("ожидалась ошибка %v, получена: %v", service.ErrPostNotExists, err)
	}
}

func TestCommentService_Create_EmptyContent(t *testing.T) {
	// Подготовка: Передаем пустую строку с пробелами
	svc := service.NewCommentService(nil, nil, nil)
	req := &model.CommentCreateRequest{Content: ""}

	// Действие
	_, err := svc.Create(context.Background(), 1, req, 1)

	// Проверка
	if err == nil {
		t.Fatal("ожидалась ошибка валидации пустого комментария, но метод отработал успешно")
	}
}
