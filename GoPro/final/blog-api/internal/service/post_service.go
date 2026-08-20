package service

import (
	"blog-api/internal/model"
	"blog-api/internal/repository"
	"context"
	"errors"
	"fmt"
)

var (
	ErrPostNotFound   = errors.New("post not found")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden       = errors.New("forbidden")
	ErrInvalidTitle   = errors.New("invalid title: must be between 1 and 200 characters")
	ErrInvalidContent = errors.New("invalid content: cannot be empty")
)

type PostService struct {
	postRepo repository.PostRepository
	userRepo repository.UserRepository
}

func NewPostService(postRepo repository.PostRepository, userRepo repository.UserRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
		userRepo: userRepo,
	}
}

func (s *PostService) Create(ctx context.Context, userID int, req *model.PostCreateRequest) (*model.Post, error) {
	// 1. Валидация данных
	if err := validatePostCreateRequest(req); err != nil {
		return nil, err
	}

	// 2. Создать модель поста с данными из запроса и userID
	post := &model.Post{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: userID,
	}

	// 3. Сохранить через репозиторий
	if err := s.postRepo.Create(ctx, post); err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	// Обогащаем данными автора
	author, err := s.userRepo.GetByID(ctx, userID)
	if err == nil && author != nil {
		post.Author = author
	}

	// 4. Вернуть созданный пост
	return post, nil
}

func (s *PostService) GetByID(ctx context.Context, id int) (*model.Post, error) {
	// 1. Получить пост через репозиторий
	post, err := s.postRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) || err.Error() == "post not found" {
			return nil, ErrPostNotFound
		}
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	// 2. Загрузить информацию об авторе
	author, err := s.userRepo.GetByID(ctx, post.AuthorID)
	if err == nil && author != nil {
		post.Author = author
	}

	// 3. Вернуть пост
	return post, nil
}

func (s *PostService) GetAll(ctx context.Context, limit, offset int) ([]*model.Post, int, error) {
	// 1. Валидировать и нормализовать параметры пагинации
	if limit <= 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	// 2. Получить посты через репозиторий
	posts, err := s.postRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all posts: %w", err)
	}

	// 3. Получить общее количество для пагинации
	totalCount, err := s.postRepo.GetTotalCount(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get posts count: %w", err)
	}

	// 4. Обогатить данные информацией об авторах
	for _, post := range posts {
		author, err := s.userRepo.GetByID(ctx, post.AuthorID)
		if err == nil && author != nil {
			post.Author = author
		}
	}

	// 5. Вернуть посты и общее количество
	return posts, totalCount, nil
}

func (s *PostService) Update(ctx context.Context, id int, userID int, req *model.PostUpdateRequest) (*model.Post, error) {
	// 1. Получить существующий пост
	post, err := s.postRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrPostNotFound
	}

	// 2. Проверить что userID является автором
	if post.AuthorID != userID {
		return nil, ErrForbidden
	}

	// 3. Валидировать новые данные
	if err := validatePostUpdateRequest(req); err != nil {
		return nil, err
	}

	// 4. Обновить только измененные поля
	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Content != "" {
		post.Content = req.Content
	}

	// 5. Сохранить через репозиторий
	if err := s.postRepo.Update(ctx, post); err != nil {
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	// Обогащаем данными автора
	author, err := s.userRepo.GetByID(ctx, userID)
	if err == nil && author != nil {
		post.Author = author
	}

	// 6. Вернуть обновленный пост
	return post, nil
}

func (s *PostService) Delete(ctx context.Context, id int, userID int) error {
	// 1. Найти пост и проверить существование
	post, err := s.postRepo.GetByID(ctx, id)
	if err != nil {
		return ErrPostNotFound
	}

	// 2. Проверить что userID является автором
	if post.AuthorID != userID {
		return ErrForbidden
	}

	// 3. Удалить через репозиторий
	if err := s.postRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	return nil
}

func (s *PostService) GetByAuthor(ctx context.Context, authorID int, limit, offset int) ([]*model.Post, int, error) {
	// 1. Валидировать параметры пагинации
	if limit <= 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	// 2. Получить посты автора через репозиторий
	posts, err := s.postRepo.GetByAuthorID(ctx, authorID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get posts by author: %w", err)
	}

	// 3. Получить общее количество постов автора
	totalCount, err := s.postRepo.GetCountByAuthorID(ctx, authorID)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get author posts count: %w", err)
	}

	// 4. Обогатить информацию об авторе к постам
	author, err := s.userRepo.GetByID(ctx, authorID)
	if err == nil && author != nil {
		for _, post := range posts {
			post.Author = author
		}
	}

	// 5. Вернуть результат с общим количеством
	return posts, totalCount, nil
}

// validatePostCreateRequest проверяет корректность данных для создания поста
func validatePostCreateRequest(req *model.PostCreateRequest) error {
	if req == nil {
		return errors.New("request body cannot be nil")
	}
	if len(req.Title) == 0 || len(req.Title) > 200 {
		return ErrInvalidTitle
	}
	if len(req.Content) == 0 {
		return ErrInvalidContent
	}
	return nil
}

// validatePostUpdateRequest проверяет корректность данных для обновления поста
func validatePostUpdateRequest(req *model.PostUpdateRequest) error {
	if req == nil {
		return errors.New("request body cannot be nil")
	}
	if req.Title != "" && len(req.Title) > 200 {
		return ErrInvalidTitle
	}
	return nil
}
