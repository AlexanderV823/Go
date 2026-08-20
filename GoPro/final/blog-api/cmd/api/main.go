package main

import (
	"blog-api/internal/handler"
	"blog-api/internal/middleware"
	"blog-api/internal/repository"
	"blog-api/internal/service"
	"blog-api/pkg/auth"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем конфигурацию из .env файла
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using environment variables")
	}

	// Загружаем конфигурацию из переменных окружения
	cfg := loadConfig()

	// Подключаемся к базе данных
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	// Настройка пула соединений
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Проверяем подключение
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}
	log.Println("Successfully connected to the database")

	// Инициализируем JWT менеджер
	jwtManager := auth.NewJWTManager(cfg.JWTSecret, cfg.JWTExpiryHours)

	// Инициализируем слои приложения (Репозитории)
	userRepo := repository.NewUserRepo(db)
	// Предполагаем существование NewPostRepo и NewCommentRepo в вашем пакете repository
	postRepo := repository.NewPostRepo(db)
	commentRepo := repository.NewCommentRepo(db)

	// Инициализируем слои приложения (Сервисы)
	userService := service.NewUserService(userRepo, jwtManager)
	commentService := service.NewCommentService(commentRepo, postRepo, userRepo)
	postService := service.NewPostService(postRepo, userRepo) // Предполагаем наличие NewPostService

	// Инициализируем слои приложения (Хендлеры)
	authHandler := handler.NewAuthHandler(userService)
	postHandler := handler.NewPostHandler(postService)
	commentHandler := handler.NewCommentHandler(commentService)

	// Инициализируем Middleware
	loggerInstance := log.New(os.Stdout, "BLOG_API: ", log.LstdFlags|log.Lshortfile)
	logMW := middleware.NewLoggingMiddleware(loggerInstance)
	authMW := middleware.NewAuthMiddleware(jwtManager)

	// Настраиваем роутер chi
	router := chi.NewRouter()

	// Настраиваем глобальные middleware
	router.Use(func(next http.Handler) http.Handler { return logMW.RequestID(next.ServeHTTP) })
	router.Use(func(next http.Handler) http.Handler { return logMW.Logger(next.ServeHTTP) })
	router.Use(func(next http.Handler) http.Handler { return logMW.Recovery(next.ServeHTTP) })
	router.Use(func(next http.Handler) http.Handler { return logMW.CORS(next.ServeHTTP) })
	router.Use(func(next http.Handler) http.Handler { return logMW.ContentTypeJSON(next.ServeHTTP) })

	// Health check эндпоинт
	router.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok","service":"blog-api"}`))
	})

	// --- МАРШРУТЫ АУТЕНТИФИКАЦИИ ---
	router.Post("/api/register", authHandler.Register)
	router.Post("/api/login", authHandler.Login)
	router.Get("/api/profile", authMW.RequireAuth(authHandler.GetProfile))

	// --- МАРШРУТЫ ПОСТОВ ---
	// Используем суффикс /* для совместимости с ручным парсингом extractIDFromPath внутри ваших хендлеров
	router.Get("/api/posts", postHandler.GetAll)
	router.Get("/api/posts/*", postHandler.GetByID)
	router.Get("/api/posts/author/*", postHandler.GetByAuthor)
	router.Post("/api/posts", authMW.RequireAuth(postHandler.Create))
	router.Put("/api/posts/*", authMW.RequireAuth(postHandler.Update))
	router.Delete("/api/posts/*", authMW.RequireAuth(postHandler.Delete))

	// --- МАРШРУТЫ КОММЕНТАРИЕВ ---
	router.Post("/api/comments", authMW.RequireAuth(commentHandler.Create))
	router.Get("/api/comments/*", commentHandler.GetByID)
	router.Put("/api/comments/*", authMW.RequireAuth(commentHandler.Update))
	router.Delete("/api/comments/*", authMW.RequireAuth(commentHandler.Delete))

	// Получение комментариев к конкретному посту (/api/posts/{id}/comments)
	router.Get("/api/posts/*/comments", commentHandler.GetByPost)

	// Запуск HTTP сервера
	serverAddr := fmt.Sprintf("%s:%d", cfg.ServerHost, cfg.ServerPort)
	log.Printf("Starting HTTP server on %s...", serverAddr)

	server := &http.Server{
		Addr:         serverAddr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v", serverAddr, err)
	}
}

// Config представляет конфигурацию приложения
type Config struct {
	ServerHost      string
	ServerPort      int
	DBHost          string
	DBPort          int
	DBUser          string
	DBPassword      string
	DBName          string
	DBSSLMode       string
	JWTSecret       string
	JWTExpiryHours  int
	CacheTTLMinutes int
}

// loadConfig загружает конфигурацию из переменных окружения
func loadConfig() *Config {
	return &Config{
		ServerHost:      getEnv("SERVER_HOST", "0.0.0.0"),
		ServerPort:      getEnvAsInt("SERVER_PORT", 8080),
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBPort:          getEnvAsInt("DB_PORT", 5432),
		DBUser:          getEnv("DB_USER", "postgres"),
		DBPassword:      getEnv("DB_PASSWORD", "postgres"),
		DBName:          getEnv("DB_NAME", "blog_db"),
		DBSSLMode:       getEnv("DB_SSLMODE", "disable"),
		JWTSecret:       getEnv("JWT_SECRET", "super_secret_key_change_me_in_production"),
		JWTExpiryHours:  getEnvAsInt("JWT_EXPIRY_HOURS", 24),
		CacheTTLMinutes: getEnvAsInt("CACHE_TTL_MINUTES", 15),
	}
}

// getEnv получает значение переменной окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt получает значение переменной окружения как int или возвращает значение по умолчанию
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
