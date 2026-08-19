package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

// Claims представляет данные, хранимые в JWT токене
type Claims struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// JWTManager управляет созданием и валидацией JWT токенов
type JWTManager struct {
	secretKey []byte
	ttl       time.Duration
}

// NewJWTManager создает новый экземпляр JWT менеджера
func NewJWTManager(secretKey string, ttlHours int) *JWTManager {
	// Инициализировать JWTManager
	return &JWTManager{
		secretKey: []byte(secretKey),
		ttl:       time.Duration(ttlHours) * time.Hour,
	}
}

// GenerateToken создает новый JWT токен для пользователя
func (m *JWTManager) GenerateToken(userID int, email, username string) (string, time.Time, error) {
	// 2. Установить время истечения токена (текущее время + ttl)
	expiresAt := time.Now().Add(m.ttl)

	// 1. Создать Claims с данными пользователя и стандартными полями
	claims := Claims{
		UserID:   userID,
		Email:    email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// 3. Создать токен используя алгоритм подписи HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 4. Подписать токен секретным ключом
	tokenString, err := token.SignedString(m.secretKey)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to sign token: %w", err)
	}

	// 5. Вернуть подписанную строку токена и время истечения
	return tokenString, expiresAt, nil
}

// ValidateToken проверяет и парсит JWT токен
func (m *JWTManager) ValidateToken(tokenString string) (*Claims, error) {
	// 1. Распарсить токен с проверкой подписи
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что метод подписи соответствует ожидаемому HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.secretKey, nil
	})

	// Обработать ошибки валидации
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	// 2. Извлечь claims из токена
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	// 4. Вернуть claims если токен валидный
	return claims, nil
}

// RefreshToken обновляет существующий токен (включая только что истекшие)
func (m *JWTManager) RefreshToken(tokenString string) (string, time.Time, error) {
	// 1. Пробуем распарсить токен без строгой проверки времени истечения (чтобы обновить протухший)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.secretKey, nil
	}, jwt.WithoutClaimsValidation()) // Игнорируем валидацию exp для чтения данных

	if err != nil {
		return "", time.Time{}, ErrInvalidToken
	}

	// 2. Извлечь данные пользователя из старого токена
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", time.Time{}, ErrInvalidToken
	}

	// 3. Сгенерировать новый токен с теми же данными
	return m.GenerateToken(claims.UserID, claims.Email, claims.Username)
}

// GetUserIDFromToken быстро извлекает ID пользователя из токена без полной валидации подписи
func (m *JWTManager) GetUserIDFromToken(tokenString string) (int, error) {
	parser := jwt.NewParser()
	claims := &Claims{}

	// Парсим только структуру данных, без проверки криптографической подписи
	_, _, err := parser.ParseUnverified(tokenString, claims)
	if err != nil {
		return 0, ErrInvalidToken
	}

	return claims.UserID, nil
}
