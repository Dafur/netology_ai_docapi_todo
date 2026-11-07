package models

import "time"

// User представляет зарегистрированного пользователя системы.

type User struct {
	ID        int       `json:"id"`         // Уникальный идентификатор пользователя
	Email     string    `json:"email"`      // Email пользователя
	Name      string    `json:"name"`       // Отображаемое имя пользователя
	CreatedAt time.Time `json:"created_at"` // Дата и время регистрации (RFC 3339)
}

// CreateUserRequest — тело запроса на регистрацию нового пользователя (POST /users).
//
// Обязательные поля: Email и Name. Сервер проверяет только корректность
// JSON при декодировании — формат Email и длина Name не проверяются.
type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`  // Email пользователя
	Name  string `json:"name" validate:"required,max=100"` // Отображаемое имя пользователя
}

// UpdateUserRequest — тело запроса на частичное обновление профиля (PUT /users/{id}).
//
// Все поля опциональны: обновляются только те, что переданы в запросе (не nil).
// Указатели нужны, чтобы отличать «поле не задано» (nil) от пустой строки.
// Сервер проверяет только корректность JSON при декодировании —
// формат Email и длина Name не проверяются.
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
	Data    interface{} `json:"data,omitempty"` // Полезная нагрузка; может отсутствовать
}
