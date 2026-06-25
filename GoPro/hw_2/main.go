package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type User struct {
	Name  string `validate:"min=3"`
	Age   int    `validate:"min=18;max=65"`
	Email string `validate:"regexp=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"`
}

// Validate проверяет поля структуры на соответствие правилам из тегов
func Validate(v interface{}) error {
	// Получаем reflect.Type и reflect.Value от v
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	// Если передан указатель на структуру, извлекаем саму структуру
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("validate: expected a struct, got %s", val.Kind())
	}

	// Проходим по всем полям через NumField()
	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := typ.Field(i)

		// Извлекаем тег через tag.Get("validate")
		tag := fieldType.Tag.Get("validate")
		if tag == "" {
			continue
		}

		// Разбираем правила через strings.Split с разделителем ";"
		rules := strings.Split(tag, ";")
		for _, rule := range rules {
			// Разбираем конкретное правило (например, min=3)
			parts := strings.Split(rule, "=")
			if len(parts) != 2 {
				continue
			}

			key := parts[0]
			param := parts[1]

			// Проверяем значение в зависимости от правила
			switch key {
			case "min":
				limit, err := strconv.Atoi(param)
				if err != nil {
					return fmt.Errorf("invalid min param for field %s", fieldType.Name)
				}

				switch fieldVal.Kind() {
				case reflect.String:
					// Учитываем кириллицу и иероглифы через конвертацию в руны []rune
					strRunes := []rune(fieldVal.String())
					if len(strRunes) < limit {
						return fmt.Errorf("field %s length is less than %d", fieldType.Name, limit)
					}
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if fieldVal.Int() < int64(limit) {
						return fmt.Errorf("field %s value is less than %d", fieldType.Name, limit)
					}
				}

			case "max":
				limit, err := strconv.Atoi(param)
				if err != nil {
					return fmt.Errorf("invalid max param for field %s", fieldType.Name)
				}

				switch fieldVal.Kind() {
				case reflect.String:
					// Учитываем кириллицу и иероглифы через конвертацию в руны []rune
					strRunes := []rune(fieldVal.String())
					if len(strRunes) > limit {
						return fmt.Errorf("field %s length is greater than %d", fieldType.Name, limit)
					}
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if fieldVal.Int() > int64(limit) {
						return fmt.Errorf("field %s value is greater than %d", fieldType.Name, limit)
					}
				}

			case "regexp":
				if fieldVal.Kind() == reflect.String {
					re, err := regexp.Compile(param)
					if err != nil {
						return fmt.Errorf("invalid regexp pattern for field %s", fieldType.Name)
					}
					if !re.MatchString(fieldVal.String()) {
						return fmt.Errorf("field %s does not match expression %s", fieldType.Name, param)
					}
				}
			}
		}
	}

	// Выход: nil, если валидация пройдена успешно
	return nil
}

func main() {
	// Тест 1: Имя "Ив" (2 руны) меньше min=3. Ожидаем ошибку.
	if err := Validate(User{Name: "Ив", Age: 18, Email: "test@example.com"}); err != nil {
		fmt.Println("Validation error:", err)
	}

	// Тест 2: Возраст 70 больше max=65. Ожидаем ошибку.
	if err := Validate(User{Name: "Иван", Age: 70, Email: "test@example.com"}); err != nil {
		fmt.Println("Validation error:", err)
	}

	// Тест 3: Email "invalid email" не проходит regexp. Ожидаем ошибку.
	if err := Validate(User{Name: "Иван", Age: 35, Email: "invalid email"}); err != nil {
		fmt.Println("Validation error:", err)
	}

	// Тест 4: Все данные валидны. Ошибки быть не должно (ничего не выведется).
	if err := Validate(User{Name: "Иван", Age: 35, Email: "test@example.com"}); err != nil {
		fmt.Println("Validation error:", err)
	}
}
