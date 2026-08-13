package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

// RegisterHandler обрабатывает регистрацию нового пользователя
// @Summary      Регистрация нового пользователя
// @Description  Создает новый аккаунт пользователя в системе
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body      RegisterRequest  true  "Данные регистрации"
// @Success      201     {object}  AuthResponse
// @Failure      400     {object}  map[string]string "error: Invalid JSON or validation failed"
// @Failure      409     {object}  map[string]string "error: User already exists"
// @Router       /register [post]
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Очищаем пробелы по краям
	req.Email = strings.TrimSpace(req.Email)
	req.Username = strings.TrimSpace(req.Username)

	if req.Email == "" || req.Username == "" || req.Password == "" {
		http.Error(w, "All fields (email, username, password) are required", http.StatusBadRequest)
		return
	}
	if len(req.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters long", http.StatusBadRequest)
		return
	}

	exists, err := UserExistsByEmail(req.Email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "User with this email already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to process password", http.StatusInternalServerError)
		return
	}

	userPtr, err := CreateUser(req.Email, req.Username, hashedPassword)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	token, err := GenerateToken(*userPtr)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	resp := AuthResponse{
		Token: token,
		User:  *userPtr, // Здесь также передаем по значению
	}

	json.NewEncoder(w).Encode(resp)
}

// LoginHandler обрабатывает вход пользователя
// @Summary      Вход в систему
// @Description  Аутентификация пользователя по email и паролю
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body      LoginRequest  true  "Данные для входа"
// @Success      200     {object}  AuthResponse
// @Failure      401     {object}  map[string]string "error: Invalid email or password"
// @Router       /login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Очищаем пробелы по краям email
	req.Email = strings.TrimSpace(req.Email)

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	userPtr, err := GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if !CheckPassword(req.Password, userPtr.PasswordHash) {
		// Возвращаем точно такое же сообщение и статус 401
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := GenerateToken(*userPtr)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK

	resp := AuthResponse{
		Token: token,
		User:  *userPtr,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return
	}
}

// ProfileHandler возвращает профиль текущего пользователя
// @Summary      Получение профиля текущего пользователя
// @Description  Возвращает данные авторизованного пользователя. Требует JWT токен.
// @Tags         profile
// @Produce      json
// @Security     BearerAuth
// @Success      200     {object}  ProfileResponse
// @Failure      401     {object}  map[string]string "error: Unauthorized"
// @Failure      404     {object}  map[string]string "error: User not found"
// @Router       /profile [get]
func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := GetUserIDFromContext(r)
	if !ok {
		// Если ID нет в контексте, значит Middleware не отработал или настроен неверно
		http.Error(w, "Unauthorized: user ID missing from context", http.StatusUnauthorized)
		return
	}

	userPtr, err := GetUserByID(userID)
	if err != nil {
		// Если пользователь не найден в базе данных — возвращаем 404 Not Found
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK

	// Заполняем структуру ответа (разыменовываем указатель через *)
	// password_hash автоматически исключен тегом json:"-" в самой структуре User
	resp := ProfileResponse{
		User: *userPtr,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return
	}
}

// HealthHandler проверяет состояние сервиса
// @Summary      Проверка здоровья сервиса
// @Description  Проверяет работоспособность приложения и стабильность подключения к базе данных PostgreSQL
// @Tags         system
// @Produce      json
// @Success      200     {object}  map[string]string "status: ok, message: Service is running"
// @Failure      503     {string}  string            "Database connection failed"
// @Router       /health [get]
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем подключение к БД
	if db != nil {
		if err := db.Ping(); err != nil {
			http.Error(w, "Database connection failed", http.StatusServiceUnavailable)
			return
		}
	}

	// Возвращаем статус OK
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"status":  "ok",
		"message": "Service is running",
	}
	json.NewEncoder(w).Encode(response)
}

// sendJSONResponse отправляет JSON ответ (вспомогательная функция)
func sendJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// sendErrorResponse отправляет JSON ответ с ошибкой (вспомогательная функция)
func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := map[string]string{"error": message}
	json.NewEncoder(w).Encode(response)
}

// parseJSONRequest парсит JSON из тела запроса (вспомогательная функция)
func parseJSONRequest(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Строгая проверка полей

	return decoder.Decode(v)
}

// validateRegisterRequest валидирует данные регистрации
func validateRegisterRequest(req *RegisterRequest) error {

	// Срезаем лишние пробелы перед проверками
	req.Email = strings.TrimSpace(req.Email)
	req.Username = strings.TrimSpace(req.Username)

	// Базовая проверка на пустоту полей
	if req.Email == "" {
		return fmt.Errorf("email is required")
	}
	if req.Username == "" {
		return fmt.Errorf("username is required")
	}
	if req.Password == "" {
		return fmt.Errorf("password is required")
	}

	// Проверяем длину username (минимум 3 символа, максимум 30 для безопасности базы)
	if len(req.Username) < 3 {
		return fmt.Errorf("username must be at least 3 characters long")
	}
	if len(req.Username) > 30 {
		return fmt.Errorf("username is too long (maximum 30 characters)")
	}

	if !usernameRegex.MatchString(req.Username) {
		return fmt.Errorf("username can only contain alphanumeric characters and underscores")
	}

	if err := ValidateEmail(req.Email); err != nil {
		return fmt.Errorf("email validation failed: %w", err)
	}

	if err := ValidatePassword(req.Password); err != nil {
		return fmt.Errorf("password validation failed: %w", err)
	}

	return nil
}

// validateLoginRequest валидирует данные входа
func validateLoginRequest(req *LoginRequest) error {
	if req.Email == "" {
		return fmt.Errorf("email is required")
	}
	if req.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}
