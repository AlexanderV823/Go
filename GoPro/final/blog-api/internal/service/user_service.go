package service

import (
	"blog-api/internal/model"
	"blog-api/internal/repository"
	"blog-api/pkg/auth"
	"context"
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
)

type UserService struct {
	userRepo   *repository.UserRepo
	jwtManager *auth.JWTManager
}

func NewUserService(userRepo *repository.UserRepo, jwtManager *auth.JWTManager) *UserService {
	return &UserService{
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

func (s *UserService) Register(ctx context.Context, req *model.UserCreateRequest) (*model.TokenResponse, error) {
	if err := validateUserCreateRequest(req); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	emailExists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if emailExists {
		return nil, ErrUserAlreadyExists
	}

	userExists, err := s.userRepo.ExistsByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if userExists {
		return nil, ErrUserAlreadyExists
	}

	passwordHash, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		PasswordHash: passwordHash,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	token, expiresAt, err := s.jwtManager.GenerateToken(user.ID, user.Email, user.Username)
	if err != nil {
		return nil, err
	}

	// ИСПРАВЛЕНО: Вызвали метод .ToResponse() для приведения типов
	return &model.TokenResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User:      user.ToResponse(),
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *model.UserLoginRequest) (*model.TokenResponse, error) {
	if err := validateUserLoginRequest(req); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// ИСПРАВЛЕНО: Изменили user.PasswordHash на user.Password
	if !auth.CheckPassword(req.Password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	token, expiresAt, err := s.jwtManager.GenerateToken(user.ID, user.Email, user.Username)
	if err != nil {
		return nil, err
	}

	// ИСПРАВЛЕНО: Вызвали метод .ToResponse() для приведения типов
	return &model.TokenResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User:      user.ToResponse(),
	}, nil
}


func (s *UserService) GetByID(ctx context.Context, id int) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.userRepo.GetByEmail(ctx, email)
}

// validateUserCreateRequest проверяет корректность данных для регистрации
func validateUserCreateRequest(req *model.UserCreateRequest) error {
	if len(req.Username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(req.Email) {
		return errors.New("invalid email format")
	}
	return auth.ValidatePasswordStrength(req.Password)
}

// validateUserLoginRequest проверяет корректность данных для входа
func validateUserLoginRequest(req *model.UserLoginRequest) error {
	if req.Email == "" || req.Password == "" {
		return errors.New("email and password are required")
	}
	return nil
}
