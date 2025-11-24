package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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
//	inputPath - путь к исходному файлу, string
//	outputPath - путь к выходному файлу, string
//  process - функция-обертка, преобразующая строку к верхнему регистру
func ReadProcessWrite(inputPath string, outputPath string, process func(string) (string, error)) error {

	fileIn, error := os.Open(inputPath)

	if error != nil {

		return fmt.Errorf("Ошибка открытия файла: %w", ErrFileNotFound)

	}

	defer fileIn.Close()

	scanner := bufio.NewScanner(fileIn)

	fileOut, error := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	defer fileOut.Close()

	if error != nil {

		return fmt.Errorf("Ошибка записи файла: %w", ErrNotCreateWriteFile)

	}

	writer := bufio.NewWriter(fileOut) //Создаёт буфер (по умолчанию 4 КБ)

	for scanner.Scan() {

		line := scanner.Text() // Получить строку

		upperString, err := process(line)

		if err != nil {
			return fmt.Errorf("Ошибка преобразования строк: %w", ErrStringToUpper)
		}

		fmt.Println(upperString) //Вывод чисто для информации. Для наглядности отладки

		_, error := writer.WriteString(upperString) //Метод WriteString пишет данные в буфер, а не напрямую в файл. Первый параметр количество записанных байт

		if error != nil {

			return fmt.Errorf("Ошибка записи файла: %w", ErrNotWriteBuffer)

		}

	}

	if error := scanner.Err(); error != nil {

		if error == io.EOF {

			writer.Flush() // Сбросить буфер в файл!

			return fmt.Errorf("Чтение файла: %w", ErrEOFFile)

		} else {

			return fmt.Errorf("Ошибка чтения файла: %w", ErrNotReadFile)

		}
	}

	return nil //Требования о возврате чего-то в случае успеха нет. Возвращать нужно только ошибки

}

func main() {

	var (
		inputPath  string
		outputPath string
	)

	fmt.Print("Введите имя файла для чтения: ")

	if _, error := fmt.Scan(&inputPath); error != nil {

		fmt.Errorf("Ошибка ввода. Не правильный тип значения: %w", ErrFilePath)

	}

	fmt.Print("Введите имя файла для записи: ")

	if _, error := fmt.Scan(&outputPath); error != nil {

		fmt.Errorf("Ошибка ввода. Не правильный тип значения: %w", ErrFilePath)

	}

proc := func(s string) (upperString string, err error) {
	if err != nil {
		return "", fmt.Errorf("Ошибка вызова toUpperWithError: %w", err)
	}

	upperString = strings.ToUpper(s)

	return upperString, nil
}

	err := ReadProcessWrite(inputPath, outputPath, proc)
	//Что передавать аргументом toUpperWithError, если функция принимает строку, которая возникнет при чтении файла из inputPath?

	if err != nil {

		fmt.Println("Ошибка при выполнении функции ReadProcessWrite\n %w", err)
		log.Println("Ошибка при выполнении функции ReadProcessWrite.")

	}

}
