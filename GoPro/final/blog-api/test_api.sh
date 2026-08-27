#!/bin/bash

# Настройки
API_URL="http://localhost:8080/api"
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Временный файл для хранения ответов curl
TMP_RESP="/tmp/api_resp.txt"

cleanup() {
    rm -f "$TMP_RESP"
}
trap cleanup EXIT

log_err_and_exit() {
    local step_name="$1"
    local expected_code="$2"
    local actual_code="$3"
    echo -e "${RED}ОШИБКА${NC}"
    echo "--------------------------------------------------"
    echo "Сбой на этапе: $step_name"
    echo "Ожидаемый статус-код: $expected_code"
    echo "Фактический статус-код: $actual_code"
    echo "Тело ответа сервера:"
    cat "$TMP_RESP"
    echo "--------------------------------------------------"
    exit 1
}

echo "=== Запуск сквозной строгой автопроверки Блог-API ==="

# ==============================================================================
# ЧАСТЬ 1: ПОЗИТИВНЫЕ СЦЕНАРИИ (HAPPY PATH)
# ==============================================================================

# 1. Проверка Health Check
echo -n "1. GET /health (Ожидается 200 OK): "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X GET "$API_URL/health")
if [ "$HTTP_CODE" -ne 200 ] || ! grep -q '"status":"ok"' "$TMP_RESP"; then
    log_err_and_exit "Health Check" "200" "$HTTP_CODE"
fi
echo -e "${GREEN}УСПЕШНО${NC}"

# Генерация уникальных пользователей для тестов
RANDOM_ID=$((1 + RANDOM % 10000))
USERNAME_A="user_a_$RANDOM_ID"
EMAIL_A="user_a_$RANDOM_ID@test.com"
PASSWORD_A="Password123!"

USERNAME_B="user_b_$RANDOM_ID"
EMAIL_B="user_b_$RANDOM_ID@test.com"
PASSWORD_B="Password123!"

# 2. Регистрация Пользователя А
echo -n "2. POST /register (Ожидается 201 Created): "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X POST "$API_URL/register" \
  -H "Content-Type: application/json" \
  -d "{\"username\":\"$USERNAME_A\",\"email\":\"$EMAIL_A\",\"password\":\"$PASSWORD_A\"}")

if [ "$HTTP_CODE" -ne 201 ]; then
    log_err_and_exit "Регистрация Пользователя А" "201" "$HTTP_CODE"
fi
echo -e "${GREEN}УСПЕШНО${NC}"

# 3. Аутентификация Пользователя А
echo -n "3. POST /login (Ожидается 200 OK): "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X POST "$API_URL/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$EMAIL_A\",\"password\":\"$PASSWORD_A\"}")

if [ "$HTTP_CODE" -ne 200 ]; then
    log_err_and_exit "Авторизация Пользователя А" "200" "$HTTP_CODE"
fi

TOKEN_A=$(grep -o '"token":"[^"]*' "$TMP_RESP" | grep -o '[^"]*$')
if [ -z "$TOKEN_A" ]; then
    log_err_and_exit "Парсинг JWT-токена" "Токен присутствует в JSON" "Токен не найден"
fi
echo -e "${GREEN}УСПЕШНО (Токен получен)${NC}"

# 4. Создание поста Пользователем А
echo -n "4. POST /posts (Ожидается 201 Created): "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X POST "$API_URL/posts" \
  -H "Authorization: Bearer $TOKEN_A" \
  -H "Content-Type: application/json" \
  -d '{"title":"Уникальный заголовок автотеста","content":"Корректное содержимое публикации на кириллице"}')

if [ "$HTTP_CODE" -ne 201 ]; then
    log_err_and_exit "Создание поста" "201" "$HTTP_CODE"
fi

POST_ID=$(grep -oP '^\{\s*"id":\s*\K[0-9]+' "$TMP_RESP" || grep -o '"id":[0-9]*' "$TMP_RESP" | head -n1 | grep -o '[0-9]*')
if [ -z "$POST_ID" ]; then
    log_err_and_exit "Извлечение ID нового поста" "Поле id в ответе" "Поле id отсутствует"
fi
echo -e "${GREEN}УСПЕШНО (ID поста: $POST_ID)${NC}"

# 5. Получение списка постов (Публичный эндпоинт)
echo -n "5. GET /posts (Ожидается 200 OK и метаданные пагинации): "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X GET "$API_URL/posts?limit=5&offset=0")
if [ "$HTTP_CODE" -ne 200 ] || ! grep -q '"limit":5' "$TMP_RESP" || ! grep -q '"offset":0' "$TMP_RESP"; then
    log_err_and_exit "Получение всех постов с пагинацией" "200 + честные limit/offset" "$HTTP_CODE"
fi
echo -e "${GREEN}УСПЕШНО${NC}"

# 6. Создание комментария Пользователем А
echo -n "6. POST /posts/$POST_ID/comments (Ожидается 201 Created): "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X POST "$API_URL/posts/$POST_ID/comments" \
  -H "Authorization: Bearer $TOKEN_A" \
  -H "Content-Type: application/json" \
  -d '{"content":"Первый содержательный комментарий автора"}')

if [ "$HTTP_CODE" -ne 201 ]; then
    log_err_and_exit "Создание комментария" "201" "$HTTP_CODE"
fi

COMMENT_ID=$(grep -oP '^\{\s*"id":\s*\K[0-9]+' "$TMP_RESP" || grep -o '"id":[0-9]*' "$TMP_RESP" | head -n1 | grep -o '[0-9]*')
if [ -z "$COMMENT_ID" ]; then
    log_err_and_exit "Извлечение ID комментария" "Поле id в ответе" "Поле id отсутствует"
fi
echo -e "${GREEN}УСПЕШНО (ID комментария: $COMMENT_ID)${NC}"

# 7. Обновление собственного комментария Пользователем А
echo -n "7. PUT /comments/$COMMENT_ID (Ожидается 200 OK): "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X PUT "$API_URL/comments/$COMMENT_ID" \
  -H "Authorization: Bearer $TOKEN_A" \
  -H "Content-Type: application/json" \
  -d '{"content":"Измененный текст комментария самим автором"}')

if [ "$HTTP_CODE" -ne 200 ]; then
    log_err_and_exit "Обновление комментария автором" "200" "$HTTP_CODE"
fi
echo -e "${GREEN}УСПЕШНО${NC}"


# ==============================================================================
# ЧАСТЬ 2: НЕГАТИВНЫЕ СЦЕНАРИИ И ПРОВЕРКА БЕЗОПАСНОСТИ (NEGATIVE TESTS)
# ==============================================================================

echo "--- Запуск негативных тестов безопасности и валидации ---"

# 8. Запрос без токена (Защищенный эндпоинт)
echo -n "8. Запрос без Authorization заголовка (Ожидается 401 Unauthorized): "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X POST "$API_URL/posts" \
  -H "Content-Type: application/json" \
  -d '{"title":"Заголовок","content":"Контент"}')

if [ "$HTTP_CODE" -ne 401 ] || ! grep -q '"error"' "$TMP_RESP"; then
    log_err_and_exit "Пропуск запроса без токена" "401 JSON" "$HTTP_CODE"
fi
echo -e "${GREEN}УСПЕШНО (Доступ заблокирован)${NC}"

# 9. Запрос с поврежденным/невалидным токеном
echo -n "9. Запрос с битым JWT-токеном (Ожидается 401 Unauthorized): "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X POST "$API_URL/posts" \
  -H "Authorization: Bearer невалидный_токен_123" \
  -H "Content-Type: application/json" \
  -d '{"title":"Заголовок","content":"Контент"}')

if [ "$HTTP_CODE" -ne 401 ] || ! grep -q '"error"' "$TMP_RESP"; then
    log_err_and_exit "Пропуск запроса с битым токеном" "401 JSON" "$HTTP_CODE"
fi
echo -e "${GREEN}УСПЕШНО (Доступ заблокирован)${NC}"

# 10. Регистрация второго пользователя (Пользователь Б) для проверки прав
echo -n "10. Создание Пользователя Б для тестов привилегий: "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X POST "$API_URL/register" \
  -H "Content-Type: application/json" \
  -d "{\"username\":\"$USERNAME_B\",\"email\":\"$EMAIL_B\",\"password\":\"$PASSWORD_B\"}")

HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X POST "$API_URL/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$EMAIL_B\",\"password\":\"$PASSWORD_B\"}")

TOKEN_B=$(grep -o '"token":"[^"]*' "$TMP_RESP" | grep -o '[^"]*$')
if [ -z "$TOKEN_B" ]; then
    log_err_and_exit "Создание/Вход Пользователя Б" "200 + Токен" "$HTTP_CODE"
fi
echo -e "${GREEN}УСПЕШНО (Пользователь Б авторизован)${NC}"

# 11. Попытка редактирования чужого поста (Пользователь Б пытается изменить пост Пользователя А)
echo -n "11. PUT чужого поста Пользователем Б (Ожидается 403 Forbidden): "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X PUT "$API_URL/posts/$POST_ID" \
  -H "Authorization: Bearer $TOKEN_B" \
  -H "Content-Type: application/json" \
  -d '{"title":"Хакерская атака","content":"Пытаюсь изменить чужую запись"}')

if [ "$HTTP_CODE" -ne 403 ] || ! grep -q '"error"' "$TMP_RESP"; then
    log_err_and_exit "Модификация чужого поста" "403 JSON" "$HTTP_CODE"
fi
echo -e "${GREEN}УСПЕШНО (В модификации отказано)${NC}"

# 12. Попытка комментирования несуществующего поста
echo -n "12. POST комментария к несуществующему посту #99999 (Ожидается 404 Not Found): "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X POST "$API_URL/posts/99999/comments" \
  -H "Authorization: Bearer $TOKEN_A" \
  -H "Content-Type: application/json" \
  -d '{"content":"Комментарий в пустоту"}')

if [ "$HTTP_CODE" -ne 404 ] || ! grep -q '"error"' "$TMP_RESP"; then
    log_err_and_exit "Комментирование фантомного поста" "404 JSON" "$HTTP_CODE"
fi
echo -e "${GREEN}УСПЕШНО (Пост не найден)${NC}"

# 13. Валидация некорректной пагинации (передача строк вместо чисел)
echo -n "13. GET /posts?limit=abc (Ожидается 400 Bad Request): "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X GET "$API_URL/posts?limit=abc")
if [ "$HTTP_CODE" -ne 400 ] || ! grep -q '"error"' "$TMP_RESP"; then
    log_err_and_exit "Передача невалидного лимита" "400 JSON" "$HTTP_CODE"
fi
echo -e "${GREEN}УСПЕШНО (Строка в пагинации отклонена)${NC}"


# ==============================================================================
# ЧАСТЬ 3: ОЧИСТКА РЕСУРСОВ (CLEANUP)
# ==============================================================================

echo "--- Удаление тестовых записей (Очистка) ---"

# 14. Удаление комментария автором
echo -n "14. DELETE /comments/$COMMENT_ID (Ожидается 204 No Content): "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X DELETE "$API_URL/comments/$COMMENT_ID" \
  -H "Authorization: Bearer $TOKEN_A")

if [ "$HTTP_CODE" -ne 204 ]; then
    log_err_and_exit "Удаление комментария автором" "204" "$HTTP_CODE"
fi
echo -e "${GREEN}УСПЕШНО${NC}"

# 15. Удаление поста автором
echo -n "15. DELETE /posts/$POST_ID (Ожидается 204 No Content): "
HTTP_CODE=$(curl -s -o "$TMP_RESP" -w "%{http_code}" -X DELETE "$API_URL/posts/$POST_ID" \
  -H "Authorization: Bearer $TOKEN_A")

if [ "$HTTP_CODE" -ne 204 ]; then
    log_err_and_exit "Удаление поста автором" "204" "$HTTP_CODE"
fi
echo -e "${GREEN}УСПЕШНО${NC}"

echo "=== Сквозная автопроверка полностью и успешно завершена ==="
exit 0
