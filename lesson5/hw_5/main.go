package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Переменные сообщений об ошибке
var (
	ErrFileNotFound       = errors.New("файл не найден")
	ErrNotReadFile        = errors.New("не удалось прочитать содержимое файла")
	ErrNotWriteFile       = errors.New("не удалось записать содержимое в выходной файл")
	ErrNotCreateWriteFile = errors.New("не удалось создать выходной файл")
	ErrNotWriteBuffer     = errors.New("не удалось записать содержимое в буфер")
	ErrEOFFile            = errors.New("достигнут конец чтения файла")
	ErrFilePath           = errors.New("укажите полный путь к файлу и расширение")
	ErrStringToUpper      = errors.New("не получилось преобразовать строку к верхнему регистру")
)

// Функция читает содержиоме файла, переводит все строки в верхний регистр и записывает в новый файл
//
// inputPath - путь к исходному файлу, string
// outputPath - путь к выходному файлу, string
// process - функция-обертка, преобразующая строку к верхнему регистру
func ReadProcessWrite(inputPath string, outputPath string, process func(string) (string, error)) (err error) {

	fileIn, err := os.Open(inputPath)

	if err != nil {

		return fmt.Errorf("Ошибка открытия файла: %s (%s)", ErrFileNotFound)

	}

	defer fileIn.Close()

	scanner := bufio.NewScanner(fileIn)

	fileOut, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {

		return fmt.Errorf("Ошибка записи файла: %s (%s)", ErrNotCreateWriteFile, err)

	}

	defer fileOut.Close()

	writer := bufio.NewWriter(fileOut) //Создаёт буфер (по умолчанию 4 КБ)

	for scanner.Scan() {

		line := scanner.Text() // Получить строку

		upperString, err := process(line)

		if err != nil {
			return fmt.Errorf("ошибка преобразования строк: %s (%s)", ErrStringToUpper, err)
		}

		fmt.Print(upperString) //Вывод чисто для информации. Для наглядности отладки

		_, err = writer.WriteString(upperString) //Метод WriteString пишет данные в буфер, а не напрямую в файл. Первый параметр количество записанных байт

		if err != nil {
			if err.Error() != "EOF" {
				return fmt.Errorf("ошибка записи в буфер: %s (%s)", ErrNotWriteBuffer, err)
			}
		}
	}

	err = writer.Flush() // Сбросить буфер в файл!

	if err != nil {
		return fmt.Errorf("ошибка записи в буфер: %s (%s)", ErrNotWriteBuffer, err)
	}

	fmt.Printf("Запись в файл из буфера...")

	return nil //Требования о возврате чего-то в случае успеха нет. Возвращать нужно только ошибки
}

func main() {

	var (
		inputPath  string
		outputPath string
	)

	fmt.Print("Введите имя файла для чтения: ")

	if _, err := fmt.Scan(&inputPath); err != nil {

		fmt.Errorf("ошибка ввода. Не правильный тип значения: %s (%s)", ErrFilePath, err)

	}

	fmt.Print("Введите имя файла для записи: ")

	if _, err := fmt.Scan(&outputPath); err != nil {

		fmt.Errorf("ошибка ввода. Не правильный тип значения: %s (%s)", ErrFilePath, err)

	}

	proc := func(s string) (upperString string, err error) {

		if err != nil {
			return "", fmt.Errorf("Ошибка вызова toUpperWithError: %s", err)
		}

		upperString = strings.ToUpper(s)
		n := "\n"

		return upperString + n, nil
	}

	err := ReadProcessWrite(inputPath, outputPath, proc)

	if err != nil {

		fmt.Println("Ошибка при выполнении функции ReadProcessWrite: %s", err)
		log.Println("Ошибка при выполнении функции ReadProcessWrite.")

	}

}
