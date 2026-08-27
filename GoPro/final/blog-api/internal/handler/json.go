package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/mail"
	"strings"
	"unicode/utf8"
)

// decodeJSONStrict настраивает строгое и безопасное чтение JSON
func decodeJSONStrict(w http.ResponseWriter, r *http.Request, dst interface{}, maxBytes int64) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		return err
	}

	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		if err == nil {
			return errors.New("body must contain only a single JSON value")
		}
		return fmt.Errorf("unexpected data after JSON: %w", err)
	}
	return nil
}

// cleanAndValidateString удаляет пробелы и проверяет длину в рунах
func cleanAndValidateString(val string, minRunes, maxRunes int) (string, bool) {
	cleaned := strings.TrimSpace(val)
	runeCount := utf8.RuneCountInString(cleaned)
	if runeCount < minRunes || runeCount > maxRunes {
		return "", false
	}
	return cleaned, true
}

// cleanAndValidateEmail нормализует email и проверяет его стандартными средствами
func cleanAndValidateEmail(email string) (string, bool) {
	cleaned := strings.ToLower(strings.TrimSpace(email))
	addr, err := mail.ParseAddress(cleaned)
	if err != nil {
		return "", false
	}
	if addr.Address == "" || utf8.RuneCountInString(addr.Address) > 100 {
		return "", false
	}
	return addr.Address, true
}

// extractIDFromPath извлекает ID из пути URL
func extractIDFromPath(path, prefix string) string {
	cleaned := strings.TrimPrefix(path, prefix)
	segments := strings.Split(cleaned, "/")
	if len(segments) > 0 {
		return segments[0]
	}
	return ""
}