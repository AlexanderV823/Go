package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Validate проверяет поля структуры на соответствие тегам validate
func Validate(v interface{}) error {
	val := reflect.ValueOf(v)

	// Если передали указатель, получаем значение, на которое он указывает
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Функция работает только со структурами
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("validate: expected a struct, got %s", val.Kind())
	}

	typ := val.Type()

	// Обходим все поля структуры
	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("validate")

		// Если тега нет, пропускаем поле
		if tag == "" {
			continue
		}

		// Разбиваем правила, если их несколько (например, min=18;max=65)
		rules := strings.Split(tag, ";")
		for _, rule := range rules {
			parts := strings.SplitN(rule, "=", 2)
			if len(parts) != 2 {
				continue
			}

			key := parts[0]
			param := parts[1]

			// Обработка конкретных правил
			switch key {
			case "min":
				if err := checkMin(fieldType.Name, fieldVal, param); err != nil {
					return err
				}
			case "max":
				if err := checkMax(fieldType.Name, fieldVal, param); err != nil {
					return err
				}
			case "regexp":
				if err := checkRegexp(fieldType.Name, fieldVal, param); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func checkMin(fieldName string, val reflect.Value, param string) error {
	limit, err := strconv.Atoi(param)
	if err != nil {
		return fmt.Errorf("invalid min param for field %s", fieldName)
	}

	switch val.Kind() {
	case reflect.String:
		if len(val.String()) < limit {
			return fmt.Errorf("field %s length is less than %d", fieldName, limit)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val.Int() < int64(limit) {
			return fmt.Errorf("field %s value is less than %d", fieldName, limit)
		}
	}
	return nil
}

func checkMax(fieldName string, val reflect.Value, param string) error {
	limit, err := strconv.Atoi(param)
	if err != nil {
		return fmt.Errorf("invalid max param for field %s", fieldName)
	}

	switch val.Kind() {
	case reflect.String:
		if len(val.String()) > limit {
			return fmt.Errorf("field %s length is greater than %d", fieldName, limit)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val.Int() > int64(limit) {
			return fmt.Errorf("field %s value is greater than %d", fieldName, limit)
		}
	}
	return nil
}

func checkRegexp(fieldName string, val reflect.Value, param string) error {
	if val.Kind() != reflect.String {
		return nil // Пропускаем, если regexp применили не к строке
	}

	re, err := regexp.Compile(param)
	if err != nil {
		return fmt.Errorf("invalid regexp pattern for field %s", fieldName)
	}

	if !re.MatchString(val.String()) {
		return fmt.Errorf("field %s does not match expression %s", fieldName, param)
	}
	return nil
}

type User struct {
	Name  string `validate:"min=3"`
	Age   int    `validate:"min=18;max=65"`
	Email string `validate:"regexp=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"`
}

func main() {
	u := User{
		Name:  "Alex",
		Age:   30,
		Email: "test@example.com",
	}

	if err := Validate(u); err != nil {
		fmt.Println("Ошибка валидации:", err)
	} else {
		fmt.Println("Структура валидна")
	}
}
