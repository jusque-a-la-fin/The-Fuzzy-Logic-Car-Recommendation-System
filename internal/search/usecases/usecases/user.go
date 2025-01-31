package usecase

import (
	"car-recommendation-service/internal/search/entities"
	"car-recommendation-service/internal/search/usecases/repository"
	"car-recommendation-service/internal/shared/adapters"
	"fmt"
)

type UserInput interface {
	SetUserID(ctx adapters.Context, userID string)
	GetUserID(ctx adapters.Context) (string, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserInput {
	return &userUseCase{r}
}

// SetUserID сохраняет в cookie идентификатор пользователя
// Входные параметры: userID - идентификатор пользователя
func (uru *userUseCase) SetUserID(ctx adapters.Context, userID string) {
	user := entities.User{ID: userID}
	uru.userRepo.SetUserIDInCookie(ctx, user.ID)
}

// GetUserID получает идентификатор пользователя из cookie
func (uru *userUseCase) GetUserID(ctx adapters.Context) (string, error) {
	user := entities.User{}
	var err error
	user.ID, err = uru.userRepo.GetUserIDFromCookie(ctx)
	if err != nil {
		return "", fmt.Errorf("error from `GetUserIDFromCookie` method, package `repository`: %v", err)
	}
	return user.ID, nil
}
