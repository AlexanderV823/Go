package writeData

import (
	"bufio"
	"fmt"
	"os"
)

// Записывает данные в файл через буфер. Если файл не существует, то создает новый.
// file -  путь к файлу для записи,
// data - данные для записи.
// Возврашает error, или nil при успешной записи.
func writeData(file string, data []byte) error {

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}
	defer f.Close() // Обязательно сразу объявляем закрытие файла

	writer := bufio.NewWriter(f) // Создаем буфер для записи
	writer.Write(data)

	if err != nil {
		return fmt.Errorf("ошибка записи данных в файл: %w", err)
	}

	writer.Flush() // Сбрасываем буфер в файл

	return nil
}
