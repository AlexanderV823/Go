package service

import (
	"blog-api/internal/model"
	"blog-api/internal/repository"
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrCommentNotFound = errors.New("comment not found")
	ErrPostNotExists   = errors.New("post does not exist")
)

type CommentService struct {
	commentRepo repository.CommentRepository
	postRepo    repository.PostRepository
	userRepo    repository.UserRepository
}

func NewCommentService(
	commentRepo repository.CommentRepository,
	postRepo repository.PostRepository,
	userRepo repository.UserRepository,
) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
		userRepo:    userRepo,
	}
}

func (s *CommentService) Create(ctx context.Context, userID int, req *model.CommentCreateRequest) (*model.Comment, error) {
	if err := validateCommentCreateRequest(req); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	exists, err := s.postRepo.Exists(ctx, req.PostID)
	if err != nil {
		return nil, fmt.Errorf("check post existence: %w", err)
	}
	if !exists {
		return nil, ErrPostNotExists
	}

	comment := &model.Comment{
		PostID:    req.PostID,
		AuthorID:  userID,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.commentRepo.Create(ctx, comment); err != nil {
		return nil, fmt.Errorf("create comment in repo: %w", err)
	}

	author, err := s.userRepo.GetByID(ctx, userID)
	if err == nil && author != nil {
		comment.Author = author
	}

	return comment, nil
}

func (s *CommentService) GetByID(ctx context.Context, id int) (*model.Comment, error) {
	comment, err := s.commentRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrCommentNotFound) {
			return nil, ErrCommentNotFound
		}
		return nil, fmt.Errorf("get comment by id: %w", err)
	}

	author, err := s.userRepo.GetByID(ctx, comment.AuthorID)
	if err == nil && author != nil {
		comment.Author = author
	}

	return comment, nil
}

func (s *CommentService) GetByPost(ctx context.Context, postID int, limit, offset int) ([]*model.Comment, int, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	exists, err := s.postRepo.Exists(ctx, postID)
	if err != nil {
		return nil, 0, fmt.Errorf("check post existence: %w", err)
	}
	if !exists {
		return nil, 0, ErrPostNotExists
	}

	comments, err := s.commentRepo.GetByPostID(ctx, postID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get comments by post id: %w", err)
	}

	total, err := s.commentRepo.GetCountByPostID(ctx, postID)
	if err != nil {
		return nil, 0, fmt.Errorf("get comments count by post id: %w", err)
	}

	for _, comment := range comments {
		author, err := s.userRepo.GetByID(ctx, comment.AuthorID)
		if err == nil && author != nil {
			comment.Author = author
		}
	}

	return comments, total, nil
}

func (s *CommentService) Update(ctx context.Context, id int, userID int, req *model.CommentUpdateRequest) (*model.Comment, error) {
	if err := validateCommentUpdateRequest(req); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	comment, err := s.commentRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrCommentNotFound) {
			return nil, ErrCommentNotFound
		}
		return nil, fmt.Errorf("get comment for update: %w", err)
	}

	if comment.AuthorID != userID {
		return nil, ErrForbidden
	}

	comment.Content = req.Content
	comment.UpdatedAt = time.Now()

	if err := s.commentRepo.Update(ctx, comment); err != nil {
		return nil, fmt.Errorf("update comment in repo: %w", err)
	}

	author, err := s.userRepo.GetByID(ctx, userID)
	if err == nil && author != nil {
		comment.Author = author
	}

	return comment, nil
}

func (s *CommentService) Delete(ctx context.Context, id int, userID int) error {
	comment, err := s.commentRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrCommentNotFound) {
			return ErrCommentNotFound
		}
		return fmt.Errorf("get comment for delete: %w", err)
	}

	if comment.AuthorID != userID {
		return ErrForbidden
	}

	if err := s.commentRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete comment from repo: %w", err)
	}

	return nil
}

func (s *CommentService) GetByAuthor(ctx context.Context, authorID int, limit, offset int) ([]*model.Comment, int, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	comments, err := s.commentRepo.GetByAuthorID(ctx, authorID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get comments by author id: %w", err)
	}

	total, err := s.commentRepo.GetCountByAuthorID(ctx, authorID)
	if err != nil {
		return nil, 0, fmt.Errorf("get comments count by author id: %w", err)
	}

	author, err := s.userRepo.GetByID(ctx, authorID)
	if err == nil && author != nil {
		for _, comment := range comments {
			comment.Author = author
		}
	}

	return comments, total, nil
}

// validateCommentCreateRequest проверяет корректность данных для создания комментария
func validateCommentCreateRequest(req *model.CommentCreateRequest) error {
	if req == nil {
		return errors.New("request body is empty")
	}
	if req.PostID <= 0 {
		return errors.New("post_id must be greater than 0")
	}
	if req.Content == "" {
		return errors.New("content cannot be empty")
	}
	if len([]rune(req.Content)) > 1000 {
		return errors.New("content cannot exceed 1000 characters")
	}
	return nil
}

// validateCommentUpdateRequest проверяет корректность данных для обновления комментария
func validateCommentUpdateRequest(req *model.CommentUpdateRequest) error {
	if req == nil {
		return errors.New("request body is empty")
	}
	if req.Content == "" {
		return errors.New("content cannot be empty")
	}
	if len([]rune(req.Content)) > 1000 {
		return errors.New("content cannot exceed 1000 characters")
	}
	return nil
}
