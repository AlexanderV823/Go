package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
	"time"
	"unicode"
)

var jwtSecret []byte

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// InitAuth инициализирует секретный ключ для JWT
func InitAuth() {
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	if len(jwtSecret) < 32 {
		panic("JWT_SECRET must be at least 32 characters long")
	}
}

// HashPassword хеширует пароль с использованием bcrypt
func HashPassword(password string) (string, error) {

	// Генерируем хеш из слайса байт пароля с дефолтной стоимостью
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Преобразуем слайс байт в строку и возвращаем результат
	return string(hashedBytes), nil
}

// CheckPassword проверяет пароль против хеша
func CheckPassword(password, hash string) bool {

	// Сравниваем хеш и чистый пароль, предварительно приведя их к []byte
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	// Если ошибки нет (err == nil), значит пароли совпадают
	return err == nil
}

// GenerateToken создает JWT токен для пользователя
func GenerateToken(user User) (string, error) {

	// Создаем Claims структуру с данными пользователя
	claims := Claims{
		UserID:   user.ID,
		Email:    user.Email,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// Устанавливаем время истечения токена через 24 часа
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			// Устанавливаем время выпуска токена
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	// Создаем новый токен, указывая метод шифрования и claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен с помощью секретного ключа
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken проверяет и парсит JWT токен
func ValidateToken(tokenString string) (*Claims, error) {

	// 1. Создаем пустую структуру claims для заполнения данными
	claims := &Claims{}

	// 2. Используем jwt.ParseWithClaims() для парсинга токена
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 3. В keyFunc проверяем, что алгоритм подписи именно HMAC (*jwt.SigningMethodHMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// 4. Возвращаем jwtSecret как ключ для проверки подписи
		return jwtSecret, nil
	})

	// Обрабатываем возможную ошибку парсинга или валидации срока действия
	if err != nil {
		return nil, err
	}

	// 5. Проверяем, что токен действительно валиден (token.Valid)
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// 6. Возвращаем claims и ошибку nil
	return claims, nil
}

// ValidatePassword проверяет требования к паролю
func ValidatePassword(password string) error {

	// 1. Проверка минимальной длины
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasDigit   bool
		hasSpecial bool
	)

	// Перебираем каждый символ (rune) в строке для поддержки UTF-8
	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}

	// 2. Проверка наличия заглавных букв
	if !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}

	// 3. Проверка наличия строчных букв (полезно для общей надежности)
	if !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}

	// 4. Проверка наличия цифр
	if !hasDigit {
		return fmt.Errorf("password must contain at least one digit")
	}

	// 5. Проверка наличия специальных символов
	if !hasSpecial {
		return fmt.Errorf("password must contain at least one special character")
	}

	return nil
}

// ValidateEmail проверяет формат email (базовая проверка)
func ValidateEmail(email string) error {

	// 1. Проверка на пустую строку
	if email == "" {
		return fmt.Errorf("email is required")
	}

	// 2. Ограничение по длине (согласно RFC 5321 максимальная длина email — 254 символа)
	if len(email) > 254 {
		return fmt.Errorf("email is too long")
	}

	// 3. Проверка формата с помощью регулярного выражения
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}

	return nil
}
