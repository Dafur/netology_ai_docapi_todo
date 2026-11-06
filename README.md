# Todo API Service

Учебный REST API-сервис для управления задачами (todo) и пользователями, написанный на Go. Сервис хранит данные в памяти, использует маршрутизатор [chi](https://github.com/go-chi/chi) и предоставляет интерактивную документацию через [Swagger UI](https://github.com/swaggo/http-swagger).

## Установка и запуск

### Требования

- Go 1.22 или новее
- [Swag CLI](https://github.com/swaggo/swag) (для генерации документации)

### Шаги

1. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/Dafur/ai-assist-it.git
   cd ai-assist-it
   ```

2. Установите зависимости:

   ```bash
   go mod download
   ```

3. Сгенерируйте Swagger-спецификацию:

   ```bash
   swag init -g main.go
   ```

   Спецификация будет сохранена в `docs/swagger.json`.

4. Запустите сервер:

   ```bash
   go run main.go
   ```

5. Сервис доступен по адресу:

   - API: `http://localhost:8080/api/v1`
   - Swagger UI: `http://localhost:8080/swagger/`
   - Спецификация OpenAPI: `http://localhost:8080/docs/swagger.json`

## Зависимости

| Пакет | Назначение |
|---|---|
| [github.com/go-chi/chi/v5](https://github.com/go-chi/chi) | Лёгкий маршрутизатор HTTP-запросов |
| [github.com/swaggo/swag](https://github.com/swaggo/swag) | Генерация Swagger/OpenAPI-документации из аннотаций |
| [github.com/swaggo/http-swagger/v2](https://github.com/swaggo/http-swagger) | Обслуживание Swagger UI |

Стандартная библиотека Go: `net/http`, `encoding/json`, `sync`, `log`, `time`.

## API

Интерактивная документация: **http://localhost:8080/swagger/**

### Список эндпоинтов

| Метод | Путь | Описание |
|---|---|---|
| GET | `/api/v1/users` | Список всех пользователей |
| GET | `/api/v1/users/{id}` | Получить пользователя по ID |
| POST | `/api/v1/users` | Создать пользователя |
| PUT | `/api/v1/users/{id}` | Обновить пользователя |
| DELETE | `/api/v1/users/{id}` | Удалить пользователя |
| GET | `/api/v1/todos` | Список всех задач |
| GET | `/api/v1/todos/{id}` | Получить задачу по ID |
| POST | `/api/v1/todos` | Создать задачу |
| PUT | `/api/v1/todos/{id}` | Обновить задачу |
| DELETE | `/api/v1/todos/{id}` | Удалить задачу |

> ⚠️ Данные хранятся в памяти — после перезапуска сервера все записи теряются.

## Примеры запросов и ответов

### Создать пользователя

**Запрос:**

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"email": "ivan@example.com", "name": "Иван Иванов"}'
```

**Ответ (`201 Created`):**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "ivan@example.com",
    "name": "Иван Иванов",
    "created_at": "2026-11-06T13:00:00Z"
  }
}
```

### Создать задачу

**Запрос:**

```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "title": "Купить продукты", "description": "Молоко, хлеб, яйца", "done": false}'
```

**Ответ (`201 Created`):**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "Купить продукты",
    "description": "Молоко, хлеб, яйца",
    "done": false,
    "created_at": "2026-11-06T13:05:00Z",
    "updated_at": "2026-11-06T13:05:00Z"
  }
}
```

### Обновить задачу

**Запрос:**

```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"done": true}'
```

**Ответ (`200 OK`):**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "Купить продукты",
    "description": "Молоко, хлеб, яйца",
    "done": true,
    "created_at": "2026-11-06T13:05:00Z",
    "updated_at": "2026-11-06T13:10:00Z"
  }
}
```

### Получить пользователя по ID

**Запрос:**

```bash
curl http://localhost:8080/api/v1/users/1
```

**Ответ (`200 OK`):**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "ivan@example.com",
    "name": "Иван Иванов",
    "created_at": "2026-11-06T13:00:00Z"
  }
}
```

### Ошибки

**Запрос несуществующего ресурса:**

```bash
curl http://localhost:8080/api/v1/todos/999
```

**Ответ (`404 Not Found`):**

```
Задача не найдена
```

**Невалидное тело запроса (`400 Bad Request`):** ответ — `Невалидный запрос`.

## Структура проекта

- [x] `main.go` — точка входа: настройка роутера chi, подключение middleware, Swagger UI
- [x] `go.mod` / `go.sum` — описание модуля и зависимостей
- [x] `internal/handlers/todo.go` — обработчики CRUD для задач (`/api/v1/todos`)
- [x] `internal/handlers/user.go` — обработчики CRUD для пользователей (`/api/v1/users`)
- [x] `internal/models/todo.go` — модели задач: `Todo`, `CreateTodoRequest`, `UpdateTodoRequest`
- [x] `internal/models/user.go` — модели пользователей и общие ответы: `User`, `CreateUserRequest`, `UpdateUserRequest`, `SuccessResponse`, `ErrorResponse`
- [ ] `docs/swagger.json` — сгенерированная OpenAPI-спецификация (создаётся командой `swag init -g main.go`)

## Лицензия

Проект распространяется под лицензией [MIT](LICENSE).

## Контакты

- Автор: **Dafur**
- GitHub: <https://github.com/Dafur>
