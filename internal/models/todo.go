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
	CreatedAt   time.Time `json:"created_at"`  // Дата и время создания (RFC 3339)
	UpdatedAt   time.Time `json:"updated_at"`  // Дата и время последнего обновления (RFC 3339)
}

// CreateTodoRequest — тело запроса на создание новой задачи (POST /todos).
//
// Обязательные поля: UserID и Title. Сервер проверяет только корректность
// JSON при декодировании — дополнительных проверок длины или формата нет.
type CreateTodoRequest struct {
	UserID      int    `json:"user_id" validate:"required"`       // Владелец задачи
	Title       string `json:"title" validate:"required,max=200"` // Краткое название задачи
	Description string `json:"description" validate:"max=1000"`   // Подробное описание задачи
	Done        bool   `json:"done"`                              // По умолчанию false (задача не выполнена)
}

// UpdateTodoRequest — тело запроса на частичное обновление задачи (PUT /todos/{id}).
//
// Все поля опциональны: обновляются только те, что переданы в запросе (не nil).
// Указатели используются для отличия «поле не задано» (nil) от нулевого значения.
// Сервер проверяет только корректность JSON при декодировании —
// дополнительных проверок длины или формата нет.
type UpdateTodoRequest struct {
	Title       *string `json:"title" validate:"omitempty,max=200"`        // Новое название; nil — не менять
	Description *string `json:"description" validate:"omitempty,max=1000"` // Новое описание; nil — не менять
	Done        *bool   `json:"done"`                                      // Новый статус; nil — не менять
}
