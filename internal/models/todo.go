package models

import "time"

// Todo представляет задачу пользователя в списке дел.
//
// Используется как основная модель хранения и как формат ответа API (JSON).
type Todo struct {
	ID          int       `json:"id"`          // Уникальный идентификатор задачи
	UserID      int       `json:"user_id"`     // Идентификатор владельца задачи
	Title       string    `json:"title"`       // Краткое название задачи
	Description string    `json:"description"` // Подробное описание задачи
	Done        bool      `json:"done"`        // Признак выполнения задачи
	CreatedAt   time.Time `json:"created_at"`  // Дата и время создания (UTC)
	UpdatedAt   time.Time `json:"updated_at"`  // Дата и время последнего обновления (UTC)
}

// CreateTodoRequest — тело запроса на создание новой задачи (POST /todos).
//
// Поля проверяются валидатором (go-playground/validator):
// UserID и Title обязательны, Title ограничен 200 символами,
// Description — 1000 символами.
type CreateTodoRequest struct {
	UserID      int    `json:"user_id" validate:"required"`       // Обязательно: владелец задачи
	Title       string `json:"title" validate:"required,max=200"` // Обязательно: до 200 символов
	Description string `json:"description" validate:"max=1000"`   // Опционально: до 1000 символов
	Done        bool   `json:"done"`                              // По умолчанию false (задача не выполнена)
}

// UpdateTodoRequest — тело запроса на частичное обновление задачи (PATCH /todos/{id}).
//
// Все поля опциональны: обновляются только те, что переданы в запросе.
// Указатели используются для отличия «поле не задано» (nil) от нулевого значения.
type UpdateTodoRequest struct {
	Title       *string `json:"title" validate:"omitempty,max=200"`        // Новое название; nil — не менять
	Description *string `json:"description" validate:"omitempty,max=1000"` // Новое описание; nil — не менять
	Done        *bool   `json:"done"`                                      // Новый статус; nil — не менять
}
