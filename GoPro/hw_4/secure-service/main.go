package main

import (
	"log"
	"net/http"
	_ "secure-service/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/joho/godotenv"
)

// @title           Secure Service API
// @version         1.0
// @description     API микросервиса авторизации и управления профилями пользователей на Go.
// @host            localhost:8080
// @BasePath        /

// @securityDefinitions.apiKey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 Введите токен в формате: Bearer <ваш_jwt_токен>
func main() {
	// Загрузка переменных окружения из .env файла
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Инициализация JWT секретного ключа
	InitAuth()

	// Инициализация подключения к базе данных
	if err := InitDB(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer CloseDB()

	// Настройка HTTP маршрутов
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/profile", AuthMiddleware(http.HandlerFunc(ProfileHandler)))
	http.HandleFunc("/health", HealthHandler)
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Запуск сервера
	port := getEnv("SERVER_PORT", "8080")
	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📝 Register: POST http://localhost:%s/register", port)
	log.Printf("🔐 Login: POST http://localhost:%s/login", port)
	log.Printf("👤 Profile: GET http://localhost:%s/profile (requires token)", port)
	log.Printf("❤️  Health: GET http://localhost:%s/health", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
