package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

// MakePath - кроссплатформенно формирует путь для сохранения файла относительно домашнего каталога пользователя.
// fileName - имя файла
func MakePath(fileName string) (filePath string, err error) {

	// Находим домашний каталог
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("ошибка создания файла: %w", err)
	}

	// Объединяем путь к домашнему каталогу и имя папки
	folderPath := filepath.Join(homeDir, "MyTasks")

	// Создаем папку, если она не существует.
	err = os.MkdirAll(folderPath, 0755)
	if err != nil {
		return "", fmt.Errorf("ошибка создания файла: %w", err)
	}

	// Возвращаем созданный файл
	return filepath.Join(folderPath, fileName), nil
}
