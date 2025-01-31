package repository

import "car-recommendation-service/internal/shared/adapters"

type UserRepository interface {
	// SetUserIDInCookie сохраняет в cookie идентификатор пользователя
	// Входной параметр: userID - идентификатор пользователя
	SetUserIDInCookie(ctx adapters.Context, userID string)

	// GetUserIDFromCookie получает из cookie идентификатор пользователя
	GetUserIDFromCookie(ctx adapters.Context) (string, error)
}
