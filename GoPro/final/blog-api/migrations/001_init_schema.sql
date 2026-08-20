-- =========================================================================
-- 1. СОЗДАНИЕ ТАБЛИЦ С ОГРАНИЧЕНИЯМИ (CONSTRAINTS)
-- =========================================================================

-- 1. Таблица пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Проверка на то, что поля не пустые строки
    CONSTRAINT check_username_not_empty CHECK (LENGTH(TRIM(username)) > 0),
    CONSTRAINT check_email_not_empty CHECK (LENGTH(TRIM(email)) > 0)
);

-- 2. Таблица постов
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    content TEXT NOT NULL,
    author_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Ограничения данных
    CONSTRAINT check_title_not_empty CHECK (LENGTH(TRIM(title)) > 0),
    CONSTRAINT check_content_not_empty CHECK (LENGTH(TRIM(content)) > 0),

    -- Внешний ключ: при удалении пользователя удаляются и его посты
    CONSTRAINT fk_posts_author FOREIGN KEY (author_id)
        REFERENCES users(id) ON DELETE CASCADE
);

-- 3. Таблица комментариев
CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    post_id INTEGER NOT NULL,
    author_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Внешние ключи: каскадное удаление при удалении поста или автора
    CONSTRAINT fk_comments_post FOREIGN KEY (post_id)
        REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_comments_author FOREIGN KEY (author_id)
        REFERENCES users(id) ON DELETE CASCADE
);

-- =========================================================================
-- 2. СОЗДАНИЕ ИНДЕКСОВ
-- =========================================================================

-- Для быстрого поиска всех постов конкретного автора
CREATE INDEX IF NOT EXISTS idx_posts_author_id ON posts(author_id);

-- Для быстрой загрузки ленты постов (сортировка от новых к старым)
CREATE INDEX IF NOT EXISTS idx_posts_created_at_desc ON posts(created_at DESC);

-- Быстрая выборка всех комментариев к конкретному посту
CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);

-- Параллельный индекс для комментариев (чтобы знать, кто автор)
CREATE INDEX IF NOT EXISTS idx_comments_author_id ON comments(author_id);


-- =========================================================================
-- 3. АВТОМАТИЗАЦИЯ ПОЛЯ updated_at (ДЛЯ POSTGRESQL)
-- =========================================================================

-- Создаем функцию, которая меняет updated_at на текущее время
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Триггер для таблицы users
CREATE OR REPLACE TRIGGER update_users_modtime
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- Триггер для таблицы posts
CREATE OR REPLACE TRIGGER update_posts_modtime
    BEFORE UPDATE ON posts
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();
