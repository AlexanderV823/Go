package auth

import (
	"crypto/rand"
	"errors"
	"math/big"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmptyPassword         = errors.New("password cannot be empty")
	ErrPasswordTooShort      = errors.New("password is too short")
	ErrNoUppercase           = errors.New("password must contain at least one uppercase letter")
	ErrNoLowercase           = errors.New("password must contain at least one lowercase letter")
	ErrNoDigit               = errors.New("password must contain at least one digit")
)

// HashPassword хеширует пароль используя bcrypt
func HashPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", ErrEmptyPassword
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// CheckPassword проверяет соответствие пароля и его хеша
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ValidatePasswordStrength проверяет надежность пароля
func ValidatePasswordStrength(password string) error {
	// Проверка на пустоту
	if len(password) == 0 {
		return ErrEmptyPassword
	}

	// Требование: Минимум 6 символов
	if len(password) < 6 {
		return ErrPasswordTooShort
	}

	var (
		hasUpper bool
		hasLower bool
		hasDigit bool
	)

	// Перебираем каждый символ пароля (как rune для поддержки UTF-8)
	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		}
	}

	// Проверка наличия заглавных букв
	if !hasUpper {
		return ErrNoUppercase
	}

	// Проверка наличия строчных букв
	if !hasLower {
		return ErrNoLowercase
	}

	// Проверка наличия цифр
	if !hasDigit {
		return ErrNoDigit
	}

	return nil
}

// GenerateRandomPassword генерирует случайный пароль
func GenerateRandomPassword(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than zero")
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"

	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}

	return string(result), nil
}
