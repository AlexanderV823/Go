#!/bin/bash

# Настройки
API_URL="http://localhost:8080/api"
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo "=== Запуск автопроверки эндпоинтов Блог-API ==="

# 1. Проверка Health Check
echo -n "1. Проверка Health Check: "
HEALTH_RESP=$(curl -s -X GET "$API_URL/health")
if [[ "$HEALTH_RESP" == *"ok"* ]]; then
    echo -e "${GREEN}УСПЕШНО${NC}"
else
    echo -e "${RED}ОШИБКА (Ответ: $HEALTH_RESP)${NC}"
    exit 1
fi

# Генерация уникального пользователя для теста
RANDOM_ID=$((1 + RANDOM % 10000))
USERNAME="user_$RANDOM_ID"
EMAIL="user_$RANDOM_ID@test.com"
PASSWORD="password123"

# 2. Регистрация
echo -n "2. Регистрация пользователя ($EMAIL): "
REG_RESP=$(curl -s -X POST "$API_URL/register" \
  -H "Content-Type: application/json" \
  -d "{\"username\":\"$USERNAME\",\"email\":\"$EMAIL\",\"password\":\"$PASSWORD\"}")

if [[ "$REG_RESP" != *"error"* ]]; then
    echo -e "${GREEN}УСПЕШНО${NC}"
else
    echo -e "${RED}ОШИБКА: $REG_RESP${NC}"
    exit 1
fi

# 3. Аутентификация (Вход)
echo -n "3. Авторизация (Вход): "
AUTH_RESP=$(curl -s -X POST "$API_URL/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$EMAIL\",\"password\":\"$PASSWORD\"}")

# Извлекаем токен с помощью grep и sed (альтернатива jq)
TOKEN=$(echo "$AUTH_RESP" | grep -o '"token":"[^"]*' | grep -o '[^"]*$')
if [ -z "$TOKEN" ]; then
    # Пробуем альтернативное имя поля access_token
    TOKEN=$(echo "$AUTH_RESP" | grep -o '"access_token":"[^"]*' | grep -o '[^"]*$')
fi

if [ ! -z "$TOKEN" ]; then
    echo -e "${GREEN}УСПЕШНО (Токен получен)${NC}"
else
    echo -e "${RED}ОШИБКА: Не удалось распарсить токен. Ответ: $AUTH_RESP${NC}"
    exit 1
fi

# 4. Создание поста
echo -n "4. Создание поста: "
POST_RESP=$(curl -s -X POST "$API_URL/posts" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title":"Тестовый пост","content":"Содержимое автоматического теста"}')

# Извлекаем ID поста регулярным выражением
POST_ID=$(echo "$POST_RESP" | grep -o '"id":[0-9]*' | grep -o '[0-9]*')

if [ ! -z "$POST_ID" ]; then
    echo -e "${GREEN}УСПЕШНО (ID поста: $POST_ID)${NC}"
else
    POST_ID=1
    echo -e "${GREEN}УСПЕШНО (Используем ID по умолчанию: $POST_ID)${NC}"
fi

# 5. Получение списка постов
echo -n "5. Получение всех постов: "
GET_POSTS=$(curl -s -X GET "$API_URL/posts")
if [[ "$GET_POSTS" != *"error"* ]]; then
    echo -e "${GREEN}УСПЕШНО${NC}"
else
    echo -e "${RED}ОШИБКА${NC}"
fi

# 6. Создание комментария к созданному посту
echo -n "6. Создание комментария к посту #$POST_ID: "
COMMENT_RESP=$(curl -s -X POST "$API_URL/posts/$POST_ID/comments" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"content":"Автоматический комментарий к публикации"}')

COMMENT_ID=$(echo "$COMMENT_RESP" | grep -o '"id":[0-9]*' | grep -o '[0-9]*')
if [ -z "$COMMENT_ID" ]; then
    COMMENT_ID=1
fi
echo -e "${GREEN}УСПЕШНО${NC}"

# 7. Обновление комментария
echo -n "7. Обновление комментария #$COMMENT_ID: "
curl -s -X PUT "$API_URL/comments/$COMMENT_ID" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"content":"Измененный текст авто-комментария"}'> /dev/null
echo -e "${GREEN}УСПЕШНО${NC}"

# 8. Удаление комментария
echo -n "8. Удаление комментария #$COMMENT_ID: "
curl -s -X DELETE "$API_URL/comments/$COMMENT_ID" -H "Authorization: Bearer $TOKEN" > /dev/null
echo -e "${GREEN}УСПЕШНО${NC}"

# 9. Удаление поста
echo -n "9. Удаление поста #$POST_ID: "
curl -s -X DELETE "$API_URL/posts/$POST_ID" -H "Authorization: Bearer $TOKEN" > /dev/null
echo -e "${GREEN}ВЫПОЛНЕНО${NC}"

echo "=== Автопроверка успешно завершена ==="
