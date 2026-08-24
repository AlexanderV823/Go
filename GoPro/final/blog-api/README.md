# Blog API - Шаблон проектной работы

## Описание проекта

Разработать REST API для блог-платформы со следующим функционалом:

Управление пользователями:
· регистрация нового пользователя (POST /api/register);
· вход в систему (POST /api/login);
· получение JWT-токена при успешном входе.

Управление постами:
· создание поста (POST /api/posts) — только для авторизованных;
· получение списка постов (GET /api/posts) — доступно всем;
· получение поста по ID (GET /api/posts/{id}) — доступно всем;
· обновление поста (PUT /api/posts/{id}) — только автор;
· удаление поста (DELETE /api/posts/{id}) — только автор.

Управление комментариями:
· добавление комментария к посту (POST /api/posts/{postId}/comments) — только для авторизованных;
· получение комментариев к посту (GET /api/posts/{postId}/comments) — доступно всем.

Дополнительные требования:
· проверка состояния сервиса (GET /api/health);
· поддержка пагинации для списка постов (параметры limit и offset);
· валидация всех входных данных;
· корректные HTTP статус-коды;
· логирование запросов.
Технические требования:
· использовать PostgreSQL в качестве БД;
· реализовать слоевую архитектуру (Handler -> Service -> Repository);
· пароли должны храниться в хешированном виде (bcrypt);
· использовать JWT для аутентификации;
· все настройки через переменные окружения (.env файл);
· авторизация (только автор может редактировать/удалять свои посты и комментарии);
· автоматическое выполнение миграций при запуске.

## Структура проекта

```
blog-api
├─ cmd
│  └─ api
│     └─ main.go
├─ docker-compose.yml
├─ go.mod
├─ go.sum
├─ internal
│  ├─ handler
│  │  ├─ auth_handler.go
│  │  ├─ comment_handler.go
│  │  └─ post_handler.go
│  ├─ middleware
│  │  ├─ auth.go
│  │  └─ logging.go
│  ├─ model
│  │  └─ models.go
│  ├─ repository
│  │  ├─ comment_repo.go
│  │  ├─ interfaces.go
│  │  ├─ post_repo.go
│  │  └─ user_repo.go
│  └─ service
│     ├─ comment_service.go
│     ├─ post_service.go
│     └─ user_service.go
├─ Makefile
├─ migrations
│  └─ 001_init_schema.sql
├─ pkg
│  ├─ auth
│  │  ├─ jwt.go
│  │  └─ password.go
│  └─ database
│     └─ postgres.go
└─ README.md
```

## Начало работы


1. **pkg/database/postgres.go**
```
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Config содержит параметры подключения к PostgreSQL
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB создает новое подключение к PostgreSQL
func NewPostgresDB(cfg Config) (*sql.DB, error) {
	// TODO: Реализовать подключение к PostgreSQL
	// Шаги:
	// 1. Сформировать строку подключения (DSN) из параметров конфигурации
	// 2. Открыть соединение с БД используя sql.Open("postgres", dsn)
	// 3. Проверить соединение методом Ping()
	// 4. Настроить пул соединений (SetMaxOpenConns, SetMaxIdleConns)
	// 5. Вернуть подключение или ошибку

	return nil, fmt.Errorf("not implemented")
}

// Migrate выполняет миграции базы данных
func Migrate(db *sql.DB) error {
	// TODO: Реализовать применение миграций
	// Шаги:
	// 1. Создать таблицу users если не существует
	// 2. Создать таблицу posts если не существует
	// 3. Создать таблицу comments если не существует
	// 4. Создать необходимые индексы
	// 5. Вернуть ошибку если что-то пошло не так
	//
	// Структура таблиц:
	// - users: id, username, email, password_hash, created_at, updated_at
	// - posts: id, title, content, author_id, created_at, updated_at
	// - comments: id, content, post_id, author_id, created_at, updated_at

	queries := []string{
		// TODO: Добавить SQL запросы для создания таблиц
		// Пример:
		// `CREATE TABLE IF NOT EXISTS users (...)`,
		// `CREATE TABLE IF NOT EXISTS posts (...)`,
		// `CREATE TABLE IF NOT EXISTS comments (...)`,
		// `CREATE INDEX IF NOT EXISTS ...`,
	}

	// TODO: Выполнить каждый запрос в транзакции
	_ = queries // Удалить после реализации

	return fmt.Errorf("not implemented")
}

// CheckConnection проверяет соединение с базой данных
func CheckConnection(db *sql.DB) error {
	// TODO: Реализовать проверку соединения
	// Использовать db.Ping() для проверки

	return fmt.Errorf("not implemented")
}

// GetDSN формирует строку подключения к PostgreSQL
func GetDSN(cfg Config) string {
	// TODO: Сформировать DSN строку
	// Формат: "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"

	return ""
}

// Close закрывает соединение с базой данных
func Close(db *sql.DB) error {
	// TODO: Корректно закрыть соединение

	return fmt.Errorf("not implemented")
}

// TestConnection выполняет тестовый запрос к БД (опциональное задание)
func TestConnection(db *sql.DB) error {
	// TODO: Выполнить простой запрос для проверки работы БД
	// Например: SELECT 1

	return fmt.Errorf("not implemented")
}
```

2. **pkg/auth/password.go**
```
package auth

import (
	"errors"
)

var (
	ErrEmptyPassword    = errors.New("password cannot be empty")
	ErrPasswordTooShort = errors.New("password is too short")
)

// HashPassword хеширует пароль используя bcrypt
func HashPassword(password string) (string, error) {
	// TODO: Реализовать хеширование пароля
	// Шаги:
	// 1. Проверить что пароль не пустой
	// 2. Использовать bcrypt для хеширования
	// 3. Выбрать подходящий cost factor (например, 10-12)
	// 4. Вернуть хешированный пароль как строку
	//
	// Подсказка: используйте golang.org/x/crypto/bcrypt

	return "", errors.New("not implemented")
}

// CheckPassword проверяет соответствие пароля и его хеша
func CheckPassword(password, hash string) bool {
	// TODO: Реализовать проверку пароля
	// Шаги:
	// 1. Сравнить пароль с хешом используя bcrypt
	// 2. Вернуть true если пароль совпадает, false если нет
	// 3. При ошибке вернуть false
	//
	// Подсказка: bcrypt.CompareHashAndPassword

	return false
}

// ValidatePasswordStrength проверяет надежность пароля
func ValidatePasswordStrength(password string) error {
	// TODO: Реализовать проверку надежности пароля
	// Требования:
	// - Минимум 6 символов
	// - Опционально: содержит буквы и цифры
	// - Опционально: содержит заглавные и строчные буквы
	//
	// Вернуть соответствующую ошибку или nil

	return errors.New("not implemented")
}

// GenerateRandomPassword генерирует случайный пароль (опциональное задание)
func GenerateRandomPassword(length int) (string, error) {
	// TODO: Реализовать генерацию случайного пароля
	// Шаги:
	// 1. Создать набор допустимых символов
	// 2. Сгенерировать случайную последовательность заданной длины
	// 3. Вернуть пароль как строку
	//
	// Подсказка: используйте crypto/rand для криптографически стойкой генерации

	return "", errors.New("not implemented")
}
```

3. **pkg/auth/jwt.go**
```
package auth

import (
	"errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

// Claims представляет данные, хранимые в JWT токене
type Claims struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	// TODO: Добавить стандартные JWT claims
	// Подсказка: используйте jwt.RegisteredClaims или jwt.StandardClaims
}

// JWTManager управляет созданием и валидацией JWT токенов
type JWTManager struct {
	secretKey []byte
	ttl       time.Duration
}

// NewJWTManager создает новый экземпляр JWT менеджера
func NewJWTManager(secretKey string, ttlHours int) *JWTManager {
	// TODO: Инициализировать JWTManager
	// - Преобразовать secretKey в []byte
	// - Преобразовать ttlHours в time.Duration

	return &JWTManager{}
}

// GenerateToken создает новый JWT токен для пользователя
func (m *JWTManager) GenerateToken(userID int, email, username string) (string, time.Time, error) {
	// TODO: Реализовать генерацию JWT токена
	// Шаги:
	// 1. Создать Claims с данными пользователя
	// 2. Установить время истечения токена (текущее время + ttl)
	// 3. Создать токен используя алгоритм подписи (например, HS256)
	// 4. Подписать токен секретным ключом
	// 5. Вернуть подписанную строку токена и время истечения
	//
	// Подсказка: используйте библиотеку github.com/golang-jwt/jwt/v5

	return "", time.Time{}, errors.New("not implemented")
}

// ValidateToken проверяет и парсит JWT токен
func (m *JWTManager) ValidateToken(tokenString string) (*Claims, error) {
	// TODO: Реализовать валидацию и парсинг JWT токена
	// Шаги:
	// 1. Распарсить токен с проверкой подписи
	// 2. Извлечь claims из токена
	// 3. Проверить время истечения токена
	// 4. Вернуть claims если токен валидный
	//
	// Обработать ошибки:
	// - Невалидная подпись -> ErrInvalidToken
	// - Истекший токен -> ErrExpiredToken
	// - Другие ошибки -> ErrInvalidToken

	return nil, errors.New("not implemented")
}

// RefreshToken обновляет существующий токен (опциональное задание)
func (m *JWTManager) RefreshToken(tokenString string) (string, time.Time, error) {
	// TODO: Реализовать обновление токена (продвинутое задание)
	// Шаги:
	// 1. Валидировать существующий токен
	// 2. Извлечь данные пользователя из старого токена
	// 3. Сгенерировать новый токен с теми же данными
	// 4. Вернуть новый токен

	return "", time.Time{}, errors.New("not implemented")
}

// GetUserIDFromToken быстро извлекает ID пользователя из токена без полной валидации
func (m *JWTManager) GetUserIDFromToken(tokenString string) (int, error) {
	// TODO: Извлечь UserID из токена (опциональное задание)
	// Может быть полезно для быстрой проверки

	return 0, errors.New("not implemented")
}
```

4. **internal/repository/user_repo.go**
```
package repository

import (
	"blog-api/internal/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")
)

// UserRepo представляет репозиторий для работы с пользователями
type UserRepo struct {
	db *sql.DB
}

// NewUserRepo создает новый репозиторий пользователей
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Create создает нового пользователя
func (r *UserRepo) Create(ctx context.Context, user *model.User) error {
	// TODO: Реализовать создание пользователя
	// 1. Подготовить SQL запрос INSERT INTO users...
	// 2. Установить created_at и updated_at = time.Now()
	// 3. Выполнить запрос и получить ID созданной записи
	// 4. Установить ID в структуру user
	//
	// HINT: Используйте QueryRowContext с RETURNING id для получения ID
	// Пример запроса:
	// INSERT INTO users (username, email, password, created_at, updated_at)
	// VALUES ($1, $2, $3, $4, $5) RETURNING id

	query := `
		INSERT INTO users (username, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// TODO: Выполнить запрос и обработать результат

	return fmt.Errorf("not implemented")
}

// GetByID получает пользователя по ID
func (r *UserRepo) GetByID(ctx context.Context, id int) (*model.User, error) {
	// TODO: Реализовать получение пользователя по ID
	// 1. Подготовить SQL запрос SELECT ... FROM users WHERE id = $1
	// 2. Выполнить запрос
	// 3. Просканировать результат в структуру User
	// 4. Обработать случай, когда пользователь не найден (sql.ErrNoRows)
	//
	// HINT: Используйте QueryRowContext и Scan

	query := `
		SELECT id, username, email, password, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user model.User
	// TODO: Выполнить запрос и просканировать результат
	// Не забудьте обработать sql.ErrNoRows и вернуть ErrUserNotFound

	_ = query // Удалите эту строку после реализации
	return nil, fmt.Errorf("not implemented")
}

// GetByEmail получает пользователя по email
func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	// TODO: Реализовать получение пользователя по email
	// Аналогично GetByID, но поиск по полю email

	return nil, fmt.Errorf("not implemented")
}

// GetByUsername получает пользователя по username
func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	// TODO: Реализовать получение пользователя по username
	// Аналогично GetByID, но поиск по полю username

	return nil, fmt.Errorf("not implemented")
}

// ExistsByEmail проверяет существование пользователя по email
func (r *UserRepo) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	// TODO: Реализовать проверку существования пользователя
	// HINT: Используйте SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	var exists bool
	// TODO: Выполнить запрос и просканировать результат в переменную exists

	_ = query // Удалите эту строку после реализации
	return false, fmt.Errorf("not implemented")
}

// ExistsByUsername проверяет существование пользователя по username
func (r *UserRepo) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	// TODO: Реализовать проверку существования пользователя по username
	// Аналогично ExistsByEmail

	return false, fmt.Errorf("not implemented")
}

// Update обновляет данные пользователя
func (r *UserRepo) Update(ctx context.Context, user *model.User) error {
	// TODO: (Опционально) Реализовать обновление пользователя
	// 1. Подготовить SQL запрос UPDATE users SET ... WHERE id = $X
	// 2. Обновить updated_at = time.Now()
	// 3. Выполнить запрос
	// 4. Проверить, что запись была обновлена (RowsAffected)

	return fmt.Errorf("not implemented")
}

// Delete удаляет пользователя
func (r *UserRepo) Delete(ctx context.Context, id int) error {
	// TODO: (Опционально) Реализовать удаление пользователя
	// 1. Подготовить SQL запрос DELETE FROM users WHERE id = $1
	// 2. Выполнить запрос
	// 3. Проверить, что запись была удалена (RowsAffected)

	return fmt.Errorf("not implemented")
}
```

5. **internal/repository/post_repo.go**
```
package repository

import (
	"blog-api/internal/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var (
	ErrPostNotFound = errors.New("post not found")
)

// PostRepo представляет репозиторий для работы с постами
type PostRepo struct {
	db *sql.DB
}

// NewPostRepo создает новый репозиторий постов
func NewPostRepo(db *sql.DB) *PostRepo {
	return &PostRepo{db: db}
}

// Create создает новый пост
func (r *PostRepo) Create(ctx context.Context, post *model.Post) error {
	// TODO: Реализовать создание поста
	// 1. Подготовить SQL запрос INSERT INTO posts...
	// 2. Установить created_at и updated_at = time.Now()
	// 3. Выполнить запрос и получить ID созданной записи
	// 4. Установить ID в структуру post
	//
	// HINT: Используйте QueryRowContext с RETURNING id
	// Пример запроса:
	// INSERT INTO posts (title, content, author_id, created_at, updated_at)
	// VALUES ($1, $2, $3, $4, $5) RETURNING id

	query := `
		INSERT INTO posts (title, content, author_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	now := time.Now()
	post.CreatedAt = now
	post.UpdatedAt = now

	// TODO: Выполнить запрос и обработать результат
	// err := r.db.QueryRowContext(ctx, query, ...).Scan(&post.ID)

	_ = query // Удалите эту строку после реализации
	return fmt.Errorf("not implemented")
}

// GetByID получает пост по ID
func (r *PostRepo) GetByID(ctx context.Context, id int) (*model.Post, error) {
	// TODO: Реализовать получение поста по ID
	// 1. Подготовить SQL запрос SELECT ... FROM posts WHERE id = $1
	// 2. Выполнить запрос
	// 3. Просканировать результат в структуру Post
	// 4. Обработать случай sql.ErrNoRows -> вернуть ErrPostNotFound

	query := `
		SELECT id, title, content, author_id, created_at, updated_at
		FROM posts
		WHERE id = $1
	`

	var post model.Post
	// TODO: Выполнить запрос и просканировать результат

	_ = query // Удалите эту строку после реализации
	return nil, fmt.Errorf("not implemented")
}

// GetAll получает все посты с пагинацией
func (r *PostRepo) GetAll(ctx context.Context, limit, offset int) ([]*model.Post, error) {
	// TODO: Реализовать получение всех постов с пагинацией
	// 1. Подготовить SQL запрос с ORDER BY created_at DESC
	// 2. Добавить LIMIT и OFFSET для пагинации
	// 3. Выполнить запрос и получить rows
	// 4. Итерировать по rows и собрать массив постов
	// 5. Не забудьте закрыть rows (defer rows.Close())
	//
	// HINT: Используйте QueryContext для получения множества записей
	// Пример запроса:
	// SELECT id, title, content, author_id, created_at, updated_at
	// FROM posts
	// ORDER BY created_at DESC
	// LIMIT $1 OFFSET $2

	query := `
		SELECT id, title, content, author_id, created_at, updated_at
		FROM posts
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	// TODO: Выполнить запрос
	// rows, err := r.db.QueryContext(ctx, query, limit, offset)
	// if err != nil { ... }
	// defer rows.Close()

	// TODO: Итерировать по результатам
	// var posts []*model.Post
	// for rows.Next() {
	//     var post model.Post
	//     err := rows.Scan(&post.ID, &post.Title, ...)
	//     posts = append(posts, &post)
	// }

	_ = query // Удалите эту строку после реализации
	return nil, fmt.Errorf("not implemented")
}

// GetTotalCount получает общее количество постов
func (r *PostRepo) GetTotalCount(ctx context.Context) (int, error) {
	// TODO: Реализовать подсчет общего количества постов
	// HINT: Используйте SELECT COUNT(*) FROM posts

	query := `SELECT COUNT(*) FROM posts`

	var count int
	// TODO: Выполнить запрос и получить количество

	_ = query // Удалите эту строку после реализации
	return 0, fmt.Errorf("not implemented")
}

// Update обновляет пост
func (r *PostRepo) Update(ctx context.Context, post *model.Post) error {
	// TODO: Реализовать обновление поста
	// 1. Подготовить SQL запрос UPDATE posts SET ... WHERE id = $X
	// 2. Обновить только title, content и updated_at
	// 3. Выполнить запрос с помощью ExecContext
	// 4. Проверить RowsAffected - если 0, вернуть ErrPostNotFound
	//
	// HINT:
	// UPDATE posts
	// SET title = $1, content = $2, updated_at = $3
	// WHERE id = $4

	query := `
		UPDATE posts
		SET title = $1, content = $2, updated_at = $3
		WHERE id = $4
	`

	post.UpdatedAt = time.Now()

	// TODO: Выполнить запрос
	// result, err := r.db.ExecContext(ctx, query, ...)
	// Проверить RowsAffected

	_ = query // Удалите эту строку после реализации
	return fmt.Errorf("not implemented")
}

// Delete удаляет пост
func (r *PostRepo) Delete(ctx context.Context, id int) error {
	// TODO: Реализовать удаление поста
	// 1. Подготовить SQL запрос DELETE FROM posts WHERE id = $1
	// 2. Выполнить запрос с помощью ExecContext
	// 3. Проверить RowsAffected - если 0, вернуть ErrPostNotFound

	query := `DELETE FROM posts WHERE id = $1`

	// TODO: Выполнить запрос и проверить результат

	_ = query // Удалите эту строку после реализации
	return fmt.Errorf("not implemented")
}

// Exists проверяет существование поста
func (r *PostRepo) Exists(ctx context.Context, id int) (bool, error) {
	// TODO: Реализовать проверку существования поста
	// HINT: SELECT EXISTS(SELECT 1 FROM posts WHERE id = $1)

	query := `SELECT EXISTS(SELECT 1 FROM posts WHERE id = $1)`

	var exists bool
	// TODO: Выполнить запрос и получить результат

	_ = query // Удалите эту строку после реализации
	return false, fmt.Errorf("not implemented")
}

// GetByAuthorID получает посты определенного автора
func (r *PostRepo) GetByAuthorID(ctx context.Context, authorID int, limit, offset int) ([]*model.Post, error) {
	// TODO: (Опционально) Реализовать получение постов автора
	// Аналогично GetAll, но с дополнительным условием WHERE author_id = $X

	return nil, fmt.Errorf("not implemented")
}
```

6. **internal/repository/comment_repo.go**
```
package repository

import (
	"blog-api/internal/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var (
	ErrCommentNotFound = errors.New("comment not found")
)

// CommentRepo представляет репозиторий для работы с комментариями
type CommentRepo struct {
	db *sql.DB
}

// NewCommentRepo создает новый репозиторий комментариев
func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

// Create создает новый комментарий
func (r *CommentRepo) Create(ctx context.Context, comment *model.Comment) error {
	// TODO: Реализовать создание комментария
	// 1. Подготовить SQL запрос INSERT INTO comments...
	// 2. Установить created_at и updated_at = time.Now()
	// 3. Выполнить запрос и получить ID созданной записи
	// 4. Установить ID в структуру comment
	//
	// HINT: Используйте QueryRowContext с RETURNING id
	// Пример запроса:
	// INSERT INTO comments (content, post_id, author_id, created_at, updated_at)
	// VALUES ($1, $2, $3, $4, $5) RETURNING id

	query := `
		INSERT INTO comments (content, post_id, author_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	now := time.Now()
	comment.CreatedAt = now
	comment.UpdatedAt = now

	// TODO: Выполнить запрос
	// err := r.db.QueryRowContext(ctx, query,
	//     comment.Content, comment.PostID, comment.AuthorID,
	//     comment.CreatedAt, comment.UpdatedAt,
	// ).Scan(&comment.ID)

	_ = query // Удалите эту строку после реализации
	return fmt.Errorf("not implemented")
}

// GetByID получает комментарий по ID
func (r *CommentRepo) GetByID(ctx context.Context, id int) (*model.Comment, error) {
	// TODO: Реализовать получение комментария по ID
	// 1. Подготовить SQL запрос SELECT ... FROM comments WHERE id = $1
	// 2. Выполнить запрос
	// 3. Просканировать результат в структуру Comment
	// 4. Обработать случай sql.ErrNoRows -> вернуть ErrCommentNotFound

	query := `
		SELECT id, content, post_id, author_id, created_at, updated_at
		FROM comments
		WHERE id = $1
	`

	var comment model.Comment
	// TODO: Выполнить запрос и просканировать результат
	// err := r.db.QueryRowContext(ctx, query, id).Scan(
	//     &comment.ID, &comment.Content, &comment.PostID,
	//     &comment.AuthorID, &comment.CreatedAt, &comment.UpdatedAt,
	// )
	// if err != nil {
	//     if err == sql.ErrNoRows {
	//         return nil, ErrCommentNotFound
	//     }
	//     return nil, fmt.Errorf("failed to get comment: %w", err)
	// }

	_ = query // Удалите эту строку после реализации
	return nil, fmt.Errorf("not implemented")
}

// GetByPostID получает комментарии к посту с пагинацией
func (r *CommentRepo) GetByPostID(ctx context.Context, postID int, limit, offset int) ([]*model.Comment, error) {
	// TODO: Реализовать получение комментариев к посту
	// 1. Подготовить SQL запрос с WHERE post_id = $1
	// 2. Добавить ORDER BY created_at ASC (комментарии по времени)
	// 3. Добавить LIMIT и OFFSET для пагинации
	// 4. Выполнить запрос и получить rows
	// 5. Итерировать по rows и собрать массив комментариев
	// 6. Не забудьте закрыть rows (defer rows.Close())
	//
	// HINT: Используйте QueryContext для получения множества записей
	// Пример запроса:
	// SELECT id, content, post_id, author_id, created_at, updated_at
	// FROM comments
	// WHERE post_id = $1
	// ORDER BY created_at ASC
	// LIMIT $2 OFFSET $3

	query := `
		SELECT id, content, post_id, author_id, created_at, updated_at
		FROM comments
		WHERE post_id = $1
		ORDER BY created_at ASC
		LIMIT $2 OFFSET $3
	`

	// TODO: Выполнить запрос
	// rows, err := r.db.QueryContext(ctx, query, postID, limit, offset)
	// if err != nil {
	//     return nil, fmt.Errorf("failed to get comments: %w", err)
	// }
	// defer rows.Close()

	// TODO: Итерировать по результатам
	// var comments []*model.Comment
	// for rows.Next() {
	//     var comment model.Comment
	//     err := rows.Scan(
	//         &comment.ID, &comment.Content, &comment.PostID,
	//         &comment.AuthorID, &comment.CreatedAt, &comment.UpdatedAt,
	//     )
	//     if err != nil {
	//         return nil, fmt.Errorf("failed to scan comment: %w", err)
	//     }
	//     comments = append(comments, &comment)
	// }
	//
	// if err = rows.Err(); err != nil {
	//     return nil, fmt.Errorf("failed to iterate comments: %w", err)
	// }
	//
	// return comments, nil

	_ = query // Удалите эту строку после реализации
	return nil, fmt.Errorf("not implemented")
}

// GetCountByPostID получает количество комментариев к посту
func (r *CommentRepo) GetCountByPostID(ctx context.Context, postID int) (int, error) {
	// TODO: Реализовать подсчет комментариев к посту
	// HINT: SELECT COUNT(*) FROM comments WHERE post_id = $1

	query := `SELECT COUNT(*) FROM comments WHERE post_id = $1`

	var count int
	// TODO: Выполнить запрос
	// err := r.db.QueryRowContext(ctx, query, postID).Scan(&count)

	_ = query // Удалите эту строку после реализации
	return 0, fmt.Errorf("not implemented")
}

// Update обновляет комментарий
func (r *CommentRepo) Update(ctx context.Context, comment *model.Comment) error {
	// TODO: (Опционально) Реализовать обновление комментария
	// 1. Обновить только content и updated_at
	// 2. Использовать UPDATE comments SET content = $1, updated_at = $2 WHERE id = $3
	// 3. Проверить RowsAffected

	query := `
		UPDATE comments
		SET content = $1, updated_at = $2
		WHERE id = $3
	`

	comment.UpdatedAt = time.Now()

	// TODO: Выполнить запрос

	_ = query // Удалите эту строку после реализации
	return fmt.Errorf("not implemented")
}

// Delete удаляет комментарий
func (r *CommentRepo) Delete(ctx context.Context, id int) error {
	// TODO: (Опционально) Реализовать удаление комментария
	// 1. DELETE FROM comments WHERE id = $1
	// 2. Проверить RowsAffected

	query := `DELETE FROM comments WHERE id = $1`

	// TODO: Выполнить запрос

	_ = query // Удалите эту строку после реализации
	return fmt.Errorf("not implemented")
}
```

7. **internal/service/user_service.go**
```
package service

import (
	"blog-api/internal/model"
	"blog-api/internal/repository"
	"blog-api/pkg/auth"
	"context"
	"errors"
	"fmt"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
)

type UserService struct {
	userRepo   repository.UserRepository
	jwtManager *auth.JWTManager
}

func NewUserService(userRepo repository.UserRepository, jwtManager *auth.JWTManager) *UserService {
	return &UserService{
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

func (s *UserService) Register(ctx context.Context, req *model.UserCreateRequest) (*model.TokenResponse, error) {
	// TODO: Реализовать регистрацию пользователя
	// Шаги:
	// 1. Валидация входных данных (username >= 3 символов, email валидный, пароль >= 6 символов)
	// 2. Проверить уникальность email через репозиторий
	// 3. Проверить уникальность username через репозиторий
	// 4. Захешировать пароль используя пакет auth
	// 5. Создать модель пользователя с хешированным паролем
	// 6. Сохранить пользователя через репозиторий
	// 7. Сгенерировать JWT токен для нового пользователя
	// 8. Вернуть TokenResponse с токеном и данными пользователя

	return nil, fmt.Errorf("not implemented")
}

func (s *UserService) Login(ctx context.Context, req *model.UserLoginRequest) (*model.TokenResponse, error) {
	// TODO: Реализовать вход пользователя
	// Шаги:
	// 1. Валидация входных данных
	// 2. Найти пользователя по email через репозиторий
	// 3. Проверить пароль используя функцию из пакета auth
	// 4. Сгенерировать JWT токен при успешной аутентификации
	// 5. Вернуть TokenResponse
	// ВАЖНО: При ошибке не раскрывать, что именно неправильно (email или пароль)

	return nil, fmt.Errorf("not implemented")
}

func (s *UserService) GetByID(ctx context.Context, id int) (*model.User, error) {
	// TODO: Получить пользователя по ID через репозиторий

	return nil, fmt.Errorf("not implemented")
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	// TODO: Получить пользователя по email через репозиторий

	return nil, fmt.Errorf("not implemented")
}

// validateUserCreateRequest проверяет корректность данных для регистрации
func validateUserCreateRequest(req *model.UserCreateRequest) error {
	// TODO: Реализовать проверку всех полей

	return nil
}

// validateUserLoginRequest проверяет корректность данных для входа
func validateUserLoginRequest(req *model.UserLoginRequest) error {
	// TODO: Реализовать проверку полей

	return nil
}
```

8. **internal/service/post_service.go**
```
package service

import (
	"blog-api/internal/model"
	"blog-api/internal/repository"
	"context"
	"errors"
	"fmt"
)

var (
	ErrPostNotFound = errors.New("post not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
)

type PostService struct {
	postRepo repository.PostRepository
	userRepo repository.UserRepository
}

func NewPostService(postRepo repository.PostRepository, userRepo repository.UserRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
		userRepo: userRepo,
	}
}

func (s *PostService) Create(ctx context.Context, userID int, req *model.PostCreateRequest) (*model.Post, error) {
	// TODO: Создать новый пост
	// Шаги:
	// 1. Валидация данных (title не пустой и <= 200 символов, content не пустой)
	// 2. Создать модель поста с данными из запроса и userID
	// 3. Сохранить через репозиторий
	// 4. Вернуть созданный пост

	return nil, fmt.Errorf("not implemented")
}

func (s *PostService) GetByID(ctx context.Context, id int) (*model.Post, error) {
	// TODO: Получить пост по ID
	// Шаги:
	// 1. Получить пост через репозиторий
	// 2. Опционально: загрузить информацию об авторе
	// 3. Вернуть пост

	return nil, fmt.Errorf("not implemented")
}

func (s *PostService) GetAll(ctx context.Context, limit, offset int) ([]*model.Post, int, error) {
	// TODO: Получить список постов с пагинацией
	// Шаги:
	// 1. Валидировать и нормализовать параметры пагинации (limit по умолчанию 10, максимум 100)
	// 2. Получить посты через репозиторий
	// 3. Получить общее количество для пагинации
	// 4. Опционально: обогатить данные информацией об авторах
	// 5. Вернуть посты и общее количество

	return nil, 0, fmt.Errorf("not implemented")
}

func (s *PostService) Update(ctx context.Context, id int, userID int, req *model.PostUpdateRequest) (*model.Post, error) {
	// TODO: Обновить пост
	// Шаги:
	// 1. Получить существующий пост
	// 2. Проверить что userID является автором (иначе ErrForbidden)
	// 3. Валидировать новые данные (если предоставлены)
	// 4. Обновить только измененные поля
	// 5. Сохранить через репозиторий
	// 6. Вернуть обновленный пост

	return nil, fmt.Errorf("not implemented")
}

func (s *PostService) Delete(ctx context.Context, id int, userID int) error {
	// TODO: Удалить пост
	// Шаги:
	// 1. Найти пост и проверить существование
	// 2. Проверить что userID является автором
	// 3. Удалить через репозиторий
	// 4. Вернуть соответствующую ошибку при неудаче

	return fmt.Errorf("not implemented")
}

func (s *PostService) GetByAuthor(ctx context.Context, authorID int, limit, offset int) ([]*model.Post, int, error) {
	// TODO: Получить посты конкретного автора
	// Шаги:
	// 1. Валидировать параметры пагинации
	// 2. Получить посты автора через репозиторий
	// 3. Получить общее количество постов автора
	// 4. Опционально: добавить информацию об авторе к постам
	// 5. Вернуть результат с общим количеством

	return nil, 0, fmt.Errorf("not implemented")
}

// validatePostCreateRequest проверяет корректность данных для создания поста
func validatePostCreateRequest(req *model.PostCreateRequest) error {
	// TODO: Реализовать валидацию title и content

	return nil
}

// validatePostUpdateRequest проверяет корректность данных для обновления поста
func validatePostUpdateRequest(req *model.PostUpdateRequest) error {
	// TODO: Реализовать валидацию опциональных полей

	return nil
}
```

9. **internal/service/comment_service.go**
```
package service

import (
	"blog-api/internal/model"
	"blog-api/internal/repository"
	"context"
	"errors"
	"fmt"
)

var (
	ErrCommentNotFound = errors.New("comment not found")
	ErrPostNotExists   = errors.New("post does not exist")
)

type CommentService struct {
	commentRepo repository.CommentRepository
	postRepo    repository.PostRepository
	userRepo    repository.UserRepository
}

func NewCommentService(
	commentRepo repository.CommentRepository,
	postRepo repository.PostRepository,
	userRepo repository.UserRepository,
) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
		userRepo:    userRepo,
	}
}

func (s *CommentService) Create(ctx context.Context, userID int, req *model.CommentCreateRequest) (*model.Comment, error) {
	// TODO: Создать новый комментарий
	// Шаги:
	// 1. Валидация данных (content не пустой и <= 1000 символов)
	// 2. Проверить что пост существует
	// 3. Создать модель комментария с userID как автором
	// 4. Сохранить через репозиторий
	// 5. Опционально: обогатить ответ информацией об авторе
	// 6. Вернуть созданный комментарий

	return nil, fmt.Errorf("not implemented")
}

func (s *CommentService) GetByID(ctx context.Context, id int) (*model.Comment, error) {
	// TODO: Получить комментарий по ID
	// Шаги:
	// 1. Получить комментарий через репозиторий
	// 2. Опционально: добавить информацию об авторе
	// 3. Вернуть результат или ErrCommentNotFound

	return nil, fmt.Errorf("not implemented")
}

func (s *CommentService) GetByPost(ctx context.Context, postID int, limit, offset int) ([]*model.Comment, int, error) {
	// TODO: Получить комментарии к посту с пагинацией
	// Шаги:
	// 1. Валидировать параметры пагинации (limit по умолчанию 20, максимум 100)
	// 2. Опционально: проверить существование поста
	// 3. Получить комментарии через репозиторий
	// 4. Получить общее количество для пагинации
	// 5. Опционально: обогатить данные информацией об авторах
	// 6. Вернуть комментарии и общее количество

	return nil, 0, fmt.Errorf("not implemented")
}

func (s *CommentService) Update(ctx context.Context, id int, userID int, req *model.CommentUpdateRequest) (*model.Comment, error) {
	// TODO: Обновить комментарий
	// Шаги:
	// 1. Найти существующий комментарий
	// 2. Проверить что userID является автором (иначе ErrForbidden)
	// 3. Валидировать новый content
	// 4. Обновить content и временную метку
	// 5. Сохранить через репозиторий
	// 6. Опционально: добавить информацию об авторе
	// 7. Вернуть обновленный комментарий

	return nil, fmt.Errorf("not implemented")
}

func (s *CommentService) Delete(ctx context.Context, id int, userID int) error {
	// TODO: Удалить комментарий
	// Шаги:
	// 1. Найти комментарий и проверить существование
	// 2. Проверить что userID является автором
	// 3. Удалить через репозиторий
	// 4. Вернуть соответствующую ошибку при неудаче

	return fmt.Errorf("not implemented")
}

func (s *CommentService) GetByAuthor(ctx context.Context, authorID int, limit, offset int) ([]*model.Comment, int, error) {
	// TODO: Получить комментарии конкретного автора
	// Шаги:
	// 1. Валидировать параметры пагинации
	// 2. Получить комментарии автора через репозиторий
	// 3. Получить общее количество комментариев автора
	// 4. Опционально: добавить информацию об авторе
	// 5. Вернуть результат с общим количеством

	return nil, 0, fmt.Errorf("not implemented")
}

// validateCommentCreateRequest проверяет корректность данных для создания комментария
func validateCommentCreateRequest(req *model.CommentCreateRequest) error {
	// TODO: Реализовать валидацию content и PostID

	return nil
}

// validateCommentUpdateRequest проверяет корректность данных для обновления комментария
func validateCommentUpdateRequest(req *model.CommentUpdateRequest) error {
	// TODO: Реализовать валидацию content

	return nil
}
```

10. **internal/handler/auth_handler.go**
```
package handler

import (
	"blog-api/internal/service"
	"context"
	"net/http"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

// Register обрабатывает запрос на регистрацию нового пользователя
// POST /api/register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать обработку регистрации
	// Шаги:
	// 1. Проверить метод запроса (должен быть POST)
	// 2. Декодировать JSON тело в UserCreateRequest
	// 3. Вызвать userService.Register
	// 4. Обработать ошибки (ErrUserAlreadyExists -> 409 Conflict)
	// 5. Вернуть JSON ответ с токеном (201 Created)

	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Login обрабатывает запрос на вход пользователя
// POST /api/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать обработку входа
	// Шаги:
	// 1. Проверить метод запроса (должен быть POST)
	// 2. Декодировать JSON тело в UserLoginRequest
	// 3. Вызвать userService.Login
	// 4. Обработать ошибки (ErrInvalidCredentials -> 401 Unauthorized)
	// 5. Вернуть JSON ответ с токеном (200 OK)

	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// GetProfile возвращает профиль текущего пользователя (опционально)
// Этот метод не используется в эталонной реализации
func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	// TODO: Опционально - реализовать получение профиля
	// Этот эндпоинт не обязателен для базовой реализации

	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// writeError отправляет JSON ответ с ошибкой
func writeError(w http.ResponseWriter, message string, statusCode int) {
	// TODO: Реализовать отправку ошибки в формате JSON
	// Создать структуру ErrorResponse и отправить как JSON

	http.Error(w, message, statusCode)
}

// getUserIDFromContext извлекает ID пользователя из контекста
func getUserIDFromContext(ctx context.Context) (int, bool) {
	// TODO: Извлечь userID из контекста
	// Ключ устанавливается в auth middleware

	return 0, false
}
```

11. **internal/handler/post_handler.go**
```
package handler

import (
	"blog-api/internal/service"
	"net/http"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

// Create обрабатывает создание нового поста
// POST /api/posts
// Требует аутентификации
func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать создание поста
	// Шаги:
	// 1. Проверить метод запроса (должен быть POST)
	// 2. Получить userID из контекста (установлен middleware)
	// 3. Декодировать JSON тело в PostCreateRequest
	// 4. Создать пост через postService.Create
	// 5. Вернуть созданный пост как JSON (201 Created)

	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// GetByID возвращает пост по ID
// GET /api/posts/{id}
// Не требует аутентификации
func (h *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать получение поста по ID
	// Шаги:
	// 1. Проверить метод запроса (должен быть GET)
	// 2. Извлечь ID из URL пути
	// 3. Получить пост через postService.GetByID
	// 4. Обработать ошибки (ErrPostNotFound -> 404)
	// 5. Вернуть пост как JSON (200 OK)

	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// GetAll возвращает список постов с пагинацией
// GET /api/posts?limit=10&offset=0
// Не требует аутентификации
func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать получение списка постов
	// Шаги:
	// 1. Проверить метод запроса (должен быть GET)
	// 2. Извлечь параметры пагинации из query string
	// 3. Получить посты через postService.GetAll
	// 4. Создать ответ с метаданными пагинации
	// 5. Вернуть список постов как JSON (200 OK)

	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Update обновляет пост
// PUT /api/posts/{id}
// Требует аутентификации, может обновить только автор
func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать обновление поста
	// Шаги:
	// 1. Проверить метод запроса (должен быть PUT)
	// 2. Получить userID из контекста
	// 3. Извлечь ID поста из URL
	// 4. Декодировать JSON тело в PostUpdateRequest
	// 5. Обновить через postService.Update
	// 6. Обработать ошибки (404 для не найден, 403 для чужого поста)
	// 7. Вернуть обновленный пост как JSON (200 OK)

	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Delete удаляет пост
// DELETE /api/posts/{id}
// Требует аутентификации, может удалить только автор
func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать удаление поста
	// Шаги:
	// 1. Проверить метод запроса (должен быть DELETE)
	// 2. Получить userID из контекста
	// 3. Извлечь ID поста из URL
	// 4. Удалить через postService.Delete
	// 5. Обработать ошибки (404 для не найден, 403 для чужого поста)
	// 6. Вернуть 204 No Content при успехе

	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// GetByAuthor возвращает посты конкретного автора
// GET /api/posts/author/{authorID}?limit=10&offset=0
// Не требует аутентификации
func (h *PostHandler) GetByAuthor(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать получение постов автора
	// Шаги:
	// 1. Проверить метод запроса (должен быть GET)
	// 2. Извлечь ID автора из URL
	// 3. Извлечь параметры пагинации из query string
	// 4. Получить посты через postService.GetByAuthor
	// 5. Создать ответ с метаданными и списком постов
	// 6. Вернуть как JSON (200 OK)

	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// extractIDFromPath извлекает ID из пути URL
func extractIDFromPath(path, prefix string) string {
	// TODO: Реализовать извлечение ID из пути
	// Пример: path = "/api/posts/123", prefix = "/api/posts/"
	// Должен вернуть "123"

	return ""
}
```

12. **internal/handler/comment_handler.go**
```
package handler

import (
	"blog-api/internal/service"
	"net/http"
)

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

// Create обрабатывает создание нового комментария
// POST /api/comments
// Требует аутентификации
func (h *CommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать создание комментария
	// HINT: Последовательность действий:

	// 1. Проверить метод запроса (должен быть POST)
	//    if r.Method != http.MethodPost {
	//        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	//        return
	//    }

	// 2. Получить ID пользователя из контекста
	//    userID, ok := getUserIDFromContext(r.Context())
	//    if !ok {
	//        writeError(w, "Unauthorized", http.StatusUnauthorized)
	//        return
	//    }

	// 3. Декодировать тело запроса
	//    var req model.CommentCreateRequest
	//    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	//        writeError(w, "Invalid request body", http.StatusBadRequest)
	//        return
	//    }

	// 4. Создать комментарий через сервис
	//    comment, err := h.commentService.Create(r.Context(), userID, &req)
	//    if err != nil {
	//        switch err {
	//        case service.ErrPostNotExists:
	//            writeError(w, "Post not found", http.StatusNotFound)
	//        default:
	//            writeError(w, "Failed to create comment", http.StatusInternalServerError)
	//        }
	//        return
	//    }

	// 5. Отправить успешный ответ
	//    w.Header().Set("Content-Type", "application/json")
	//    w.WriteHeader(http.StatusCreated)
	//    json.NewEncoder(w).Encode(comment)

	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// GetByID возвращает комментарий по ID
// GET /api/comments/{id}
// Не требует аутентификации
func (h *CommentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать получение комментария по ID
	// HINT: Последовательность действий:

	// 1. Проверить метод запроса (должен быть GET)
	//    if r.Method != http.MethodGet {
	//        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	//        return
	//    }

	// 2. Извлечь ID из URL
	//    // Примерный URL: /api/comments/123
	//    idStr := extractIDFromPath(r.URL.Path, "/api/comments/")
	//    id, err := strconv.Atoi(idStr)
	//    if err != nil {
	//        writeError(w, "Invalid comment ID", http.StatusBadRequest)
	//        return
	//    }

	// 3. Получить комментарий через сервис
	//    comment, err := h.commentService.GetByID(r.Context(), id)
	//    if err != nil {
	//        if err == service.ErrCommentNotFound {
	//            writeError(w, "Comment not found", http.StatusNotFound)
	//        } else {
	//            writeError(w, "Failed to get comment", http.StatusInternalServerError)
	//        }
	//        return
	//    }

	// 4. Отправить ответ
	//    w.Header().Set("Content-Type", "application/json")
	//    w.WriteHeader(http.StatusOK)
	//    json.NewEncoder(w).Encode(comment)

	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// GetByPost возвращает комментарии к посту
// GET /api/posts/{id}/comments?limit=20&offset=0
// Не требует аутентификации
func (h *CommentHandler) GetByPost(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать получение комментариев к посту
	// HINT: Последовательность действий:

	// 1. Проверить метод запроса (должен быть GET)
	//    if r.Method != http.MethodGet {
	//        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	//        return
	//    }

	// 2. Извлечь ID поста из URL
	//    // URL вида: /api/posts/123/comments
	//    // Нужно извлечь "123"
	//    path := r.URL.Path
	//    // Пример парсинга:
	//    // - убрать префикс "/api/posts/"
	//    // - взять часть до "/comments"
	//    idStr := extractPostIDFromCommentsPath(path)
	//    postID, err := strconv.Atoi(idStr)
	//    if err != nil {
	//        writeError(w, "Invalid post ID", http.StatusBadRequest)
	//        return
	//    }

	// 3. Извлечь параметры пагинации
	//    query := r.URL.Query()
	//    limit, _ := strconv.Atoi(query.Get("limit"))
	//    if limit <= 0 {
	//        limit = 20 // значение по умолчанию
	//    }
	//    offset, _ := strconv.Atoi(query.Get("offset"))
	//    if offset < 0 {
	//        offset = 0
	//    }

	// 4. Получить комментарии через сервис
	//    comments, total, err := h.commentService.GetByPost(r.Context(), postID, limit, offset)
	//    if err != nil {
	//        if err == service.ErrPostNotExists {
	//            writeError(w, "Post not found", http.StatusNotFound)
	//        } else {
	//            writeError(w, "Failed to get comments", http.StatusInternalServerError)
	//        }
	//        return
	//    }

	// 5. Создать ответ с метаданными
	//    type CommentsResponse struct {
	//        Comments []*model.Comment `json:"comments"`
	//        Total    int             `json:"total"`
	//        Limit    int             `json:"limit"`
	//        Offset   int             `json:"offset"`
	//        PostID   int             `json:"post_id"`
	//    }
	//
	//    resp := CommentsResponse{
	//        Comments: comments,
	//        Total:    total,
	//        Limit:    limit,
	//        Offset:   offset,
	//        PostID:   postID,
	//    }

	// 6. Отправить ответ
	//    w.Header().Set("Content-Type", "application/json")
	//    w.WriteHeader(http.StatusOK)
	//    json.NewEncoder(w).Encode(resp)

	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Update обновляет комментарий
// PUT /api/comments/{id}
// Требует аутентификации, может обновить только автор
func (h *CommentHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать обновление комментария
	// HINT: Последовательность действий:

	// 1. Проверить метод запроса (должен быть PUT)
	//    if r.Method != http.MethodPut {
	//        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	//        return
	//    }

	// 2. Получить ID пользователя из контекста
	//    userID, ok := getUserIDFromContext(r.Context())
	//    if !ok {
	//        writeError(w, "Unauthorized", http.StatusUnauthorized)
	//        return
	//    }

	// 3. Извлечь ID комментария из URL
	//    idStr := extractIDFromPath(r.URL.Path, "/api/comments/")
	//    id, err := strconv.Atoi(idStr)
	//    if err != nil {
	//        writeError(w, "Invalid comment ID", http.StatusBadRequest)
	//        return
	//    }

	// 4. Декодировать тело запроса
	//    var req model.CommentUpdateRequest
	//    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	//        writeError(w, "Invalid request body", http.StatusBadRequest)
	//        return
	//    }

	// 5. Обновить комментарий через сервис
	//    comment, err := h.commentService.Update(r.Context(), id, userID, &req)
	//    if err != nil {
	//        switch err {
	//        case service.ErrCommentNotFound:
	//            writeError(w, "Comment not found", http.StatusNotFound)
	//        case service.ErrForbidden:
	//            writeError(w, "You can only update your own comments", http.StatusForbidden)
	//        default:
	//            writeError(w, "Failed to update comment", http.StatusInternalServerError)
	//        }
	//        return
	//    }

	// 6. Отправить обновленный комментарий
	//    w.Header().Set("Content-Type", "application/json")
	//    w.WriteHeader(http.StatusOK)
	//    json.NewEncoder(w).Encode(comment)

	http.Error(w, "Not implemented", http.Status
```

13. **internal/middleware/auth.go**
```
package middleware

import (
	"blog-api/pkg/auth"
	"context"
	"net/http"
)

// contextKey is a custom type for context keys to avoid collisions
type contextKey string

const (
	// UserIDKey is the key for storing user ID in context
	UserIDKey contextKey = "userID"
	// UserEmailKey is the key for storing user email in context
	UserEmailKey contextKey = "userEmail"
	// UserNameKey is the key for storing username in context
	UserNameKey contextKey = "username"
)

// AuthMiddleware provides JWT authentication
type AuthMiddleware struct {
	jwtManager *auth.JWTManager
}

// NewAuthMiddleware creates a new auth middleware instance
func NewAuthMiddleware(jwtManager *auth.JWTManager) *AuthMiddleware {
	return &AuthMiddleware{
		jwtManager: jwtManager,
	}
}

// RequireAuth is a middleware that requires valid JWT token
func (m *AuthMiddleware) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Реализовать проверку JWT токена
		// Шаги:
		// 1. Извлечь токен из заголовка Authorization (Bearer токен)
		// 2. Валидировать токен через jwtManager
		// 3. Обработать ошибки валидации (истек, невалидный и т.д.)
		// 4. Добавить данные пользователя в контекст (UserIDKey, UserEmailKey, UserNameKey)
		// 5. Передать управление следующему handler

		// Временная заглушка - удалить после реализации
		http.Error(w, "Authentication not implemented", http.StatusNotImplemented)
	}
}

// OptionalAuth is a middleware that extracts JWT token if present, but doesn't require it
func (m *AuthMiddleware) OptionalAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Реализовать опциональную проверку JWT токена
		// Шаги:
		// 1. Попытаться извлечь токен из заголовка
		// 2. Если токен есть - валидировать его
		// 3. Если токен валидный - добавить данные в контекст
		// 4. Если токена нет или он невалидный - продолжить как анонимный
		// 5. В любом случае передать управление следующему handler

		// Временная реализация
		next(w, r)
	}
}

// extractToken извлекает JWT токен из заголовка Authorization
func extractToken(r *http.Request) string {
	// TODO: Извлечь JWT токен из заголовка Authorization
	// Формат: "Bearer <token>"

	return ""
}

// GetUserIDFromContext извлекает ID пользователя из контекста
func GetUserIDFromContext(ctx context.Context) (int, bool) {
	// TODO: Извлечь userID из контекста (ключ UserIDKey)

	return 0, false
}

// GetUserEmailFromContext извлекает email пользователя из контекста
func GetUserEmailFromContext(ctx context.Context) (string, bool) {
	// TODO: Извлечь email из контекста (ключ UserEmailKey)

	return "", false
}

// GetUsernameFromContext извлекает username из контекста
func GetUsernameFromContext(ctx context.Context) (string, bool) {
	// TODO: Извлечь username из контекста (ключ UserNameKey)

	return "", false
}

// writeJSONError отправляет ошибку в формате JSON
func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	// TODO: Отправить ошибку в формате JSON
	// Создать структуру ErrorResponse и отправить как JSON

	// Временная реализация
	http.Error(w, message, statusCode)
}

// Вспомогательные функции для упрощения использования middleware

// Chain позволяет объединить несколько middleware в цепочку
func Chain(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	// TODO: Реализовать объединение middleware в цепочку
	// Применить их в правильном порядке

	return handler
}
```

14. **internal/middleware/logging.go**
```
package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware provides request logging, CORS, recovery and other utility middleware
type LoggingMiddleware struct {
	logger *log.Logger
}

// NewLoggingMiddleware creates a new logging middleware instance
func NewLoggingMiddleware(logger *log.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{
		logger: logger,
	}
}

// Logger logs all HTTP requests
func (m *LoggingMiddleware) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Реализовать логирование запросов
		// Шаги:
		// 1. Засечь время начала запроса
		// 2. Создать wrapper для ResponseWriter чтобы захватить статус код
		// 3. Вызвать следующий handler с wrapped writer
		// 4. После выполнения залогировать: метод, путь, IP, статус, время выполнения

		// Временная реализация
		next(w, r)
	}
}

// Recovery восстанавливается после паник
func (m *LoggingMiddleware) Recovery(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Реализовать восстановление после паник
		// Шаги:
		// 1. Использовать defer с recover() для перехвата паник
		// 2. При панике залогировать ошибку
		// 3. Опционально: добавить stack trace
		// 4. Вернуть клиенту 500 Internal Server Error
		// 5. Вызвать следующий handler

		// Временная реализация
		next(w, r)
	}
}

// CORS добавляет CORS заголовки
func (m *LoggingMiddleware) CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Реализовать CORS заголовки
		// Шаги:
		// 1. Добавить необходимые CORS заголовки (Origin, Methods, Headers, Max-Age)
		// 2. Обработать preflight запросы (OPTIONS метод) - вернуть 204
		// 3. Для остальных методов вызвать следующий handler

		// Временная реализация
		next(w, r)
	}
}

// RequestID добавляет уникальный ID к каждому запросу
func (m *LoggingMiddleware) RequestID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Реализовать генерацию Request ID
		// Шаги:
		// 1. Сгенерировать уникальный ID (UUID или timestamp+random)
		// 2. Добавить ID в контекст запроса для использования в логах
		// 3. Добавить ID в заголовок ответа X-Request-ID
		// 4. Залогировать запрос с Request ID
		// 5. Вызвать следующий handler

		// Временная реализация
		next(w, r)
	}
}

// RateLimiter ограничивает количество запросов от одного клиента
func (m *LoggingMiddleware) RateLimiter(maxRequests int, window time.Duration) func(http.HandlerFunc) http.HandlerFunc {
	// TODO: Реализовать rate limiting (продвинутое задание)
	// Шаги:
	// 1. Создать хранилище для отслеживания запросов по IP адресам
	// 2. Использовать mutex для безопасного доступа к хранилищу
	// 3. Для каждого запроса:
	//    - Получить IP клиента
	//    - Проверить количество запросов в текущем окне времени
	//    - Если превышен лимит - вернуть 429 Too Many Requests
	//    - Иначе увеличить счетчик и пропустить запрос

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			// Временная реализация
			next(w, r)
		}
	}
}

// ContentTypeJSON устанавливает Content-Type: application/json для всех ответов
func (m *LoggingMiddleware) ContentTypeJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Установить Content-Type: application/json для всех ответов

		// Временная реализация
		next(w, r)
	}
}

// getClientIP извлекает IP адрес клиента
func getClientIP(r *http.Request) string {
	// TODO: Извлечь реальный IP адрес клиента
	// Проверить заголовки: X-Forwarded-For, X-Real-IP, затем RemoteAddr
	// Учесть что X-Forwarded-For может содержать несколько IP

	return r.RemoteAddr
}

// responseWriter обертка для захвата статус кода
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	written    bool
}

// WriteHeader сохраняет статус код
func (rw *responseWriter) WriteHeader(code int) {
	if !rw.written {
		rw.statusCode = code
		rw.ResponseWriter.WriteHeader(code)
		rw.written = true
	}
}

// Write вызывает WriteHeader если еще не был вызван
func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.written {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(b)
}

// newResponseWriter создает новую обертку
func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		written:        false,
	}
}
```

15. **cmd/api/main.go**
```
package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем конфигурацию из .env файла
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using environment variables")
	}

	// TODO: Загрузить конфигурацию из переменных окружения
	// cfg := loadConfig()

	// TODO: Подключиться к базе данных
	// - Создать database.Config из параметров конфигурации
	// - Вызвать database.NewPostgresDB
	// - Обработать ошибки подключения
	// - Не забыть defer db.Close()

	// TODO: Выполнить миграции базы данных
	// - Вызвать database.Migrate(db)

	// TODO: Инициализировать JWT менеджер
	// - Создать jwtManager через auth.NewJWTManager

	// TODO: Создать слои приложения
	// 1. Репозитории (передать db)
	// 2. Сервисы (передать репозитории и jwtManager)
	// 3. Хендлеры (передать сервисы)
	// 4. Middleware (передать необходимые зависимости)

	// Настраиваем маршруты
	router := chi.NewRouter()

	// TODO: Настроить middleware
	// - Добавить глобальные middleware (logging, recovery, CORS)

	// TODO: Настроить маршруты
	// Публичные эндпоинты:
	// - POST /api/register
	// - POST /api/login
	// - GET /api/posts
	// - GET /api/posts/{id}
	// - GET /api/posts/{id}/comments
	//
	// Защищенные эндпоинты (требуют JWT):
	// - POST /api/posts
	// - PUT /api/posts/{id}
	// - DELETE /api/posts/{id}
	// - POST /api/posts/{id}/comments

	// Health check эндпоинт
	router.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","service":"blog-api"}`))
	})

	// TODO: Запустить HTTP сервер
	// - Сформировать адрес из конфигурации
	// - Вывести информацию о запуске
	// - Запустить сервер и обработать ошибки
}

// Config представляет конфигурацию приложения
type Config struct {
	// Server
	ServerHost string
	ServerPort int

	// Database
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// JWT
	JWTSecret      string
	JWTExpiryHours int

	// Cache
	CacheTTLMinutes int
}

// loadConfig загружает конфигурацию из переменных окружения
func loadConfig() *Config {
	// TODO: Реализовать загрузку всех параметров конфигурации
	// Использовать вспомогательные функции getEnv и getEnvAsInt
	// Установить разумные значения по умолчанию

	return nil // Заменить на правильную реализацию
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
```

## API Эндпоинты

### Публичные (без аутентификации)
- `POST /api/register` - регистрация пользователя
- `POST /api/login` - вход пользователя
- `GET /api/posts` - список постов
- `GET /api/posts/{id}` - получить пост
- `GET /api/posts/{id}/comments` - комментарии к посту

### Защищенные (требуют JWT токен)
- `POST /api/posts` - создать пост
- `PUT /api/posts/{id}` - обновить пост (только автор)
- `DELETE /api/posts/{id}` - удалить пост (только автор)
- `POST /api/posts/{id}/comments` - создать комментарий к посту

# Примеры запросов
# Регистрация
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'

# Вход
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# Создание поста (с токеном)
curl -X POST http://localhost:8080/api/posts \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title":"My Post","content":"Post content"}'
```


1. **Не забудьте обработку ошибок** - всегда проверяйте err != nil
2. **SQL injection** - используйте placeholder'ы ($1, $2) в SQL запросах
3. **Контекст** - передавайте context во все методы работы с БД
4. **Закрытие ресурсов** - используйте defer для rows.Close()
5. **Права доступа** - проверяйте что пользователь может редактировать только свои данные