package models

import "time"

// User представляет зарегистрированного пользователя системы.

type User struct {
	ID        int       `json:"id"`         // Уникальный идентификатор пользователя
	Email     string    `json:"email"`      // Email пользователя (уникальный, в нижнем регистре)
	Name      string    `json:"name"`       // Отображаемое имя пользователя
	CreatedAt time.Time `json:"created_at"` // Дата и время регистрации (UTC)
}

// CreateUserRequest — тело запроса на регистрацию нового пользователя (POST /users).
//
// Поля проверяются валидатором (go-playground/validator):
// Email должен быть валидным адресом, Name ограничен 100 символами.
type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`  // Обязательно: валидный email
	Name  string `json:"name" validate:"required,max=100"` // Обязательно: до 100 символов
}

// UpdateUserRequest — тело запроса на частичное обновление профиля (PATCH /users/{id}).
//
// Все поля опциональны: обновляются только те, что переданы в запросе.
// Указатели нужны, чтобы отличать «поле не задано» (nil) от пустой строки.
type UpdateUserRequest struct {
	Email *string `json:"email" validate:"omitempty,email"`  // Новый email; nil — не менять
	Name  *string `json:"name" validate:"omitempty,max=100"` // Новое имя; nil — не менять
}

// ErrorResponse — единый формат ответа API при ошибке.
//
// Возвращается с HTTP-статусами 4xx и 5xx.
type ErrorResponse struct {
	Error   string `json:"error"`   // Машиночитаемый код ошибки ("not_found", "validation_error", "internal_error")
	Message string `json:"message"` // Человекочитаемое описание ошибки
}

// SuccessResponse — обёртка для успешных ответов API.
//
// Data используется, когда эндпоинт возвращает полезную нагрузку:
// объект пользователя, задачу или их массив.
type SuccessResponse struct {
	Success bool        `json:"success"`        // Всегда true в этом формате ответа
	Data    interface{} `json:"data,omitempty"` // Полезная нагрузка; может отсутствовать (например, при удалении)
}
