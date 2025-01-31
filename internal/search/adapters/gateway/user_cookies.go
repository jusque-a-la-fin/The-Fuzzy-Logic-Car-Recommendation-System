package gateway

import (
	"car-recommendation-service/internal/search/usecases/repository"
	"car-recommendation-service/internal/shared/adapters"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const cookieName string = "userID"

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

// SetUserIDInCookie сохраняет в cookie идентификатор пользователя
// Входные параметры: userID - идентификатор пользователя
func (usr userRepository) SetUserIDInCookie(ctx adapters.Context, userID string) {
	ctx.(*gin.Context).SetSameSite(http.SameSiteStrictMode)
	ctx.SetCookie(cookieName, userID, 0, "/", "localhost", false, false)
}

// GetUserIDFromCookie получает из cookie идентификатор пользователя
func (usr userRepository) GetUserIDFromCookie(ctx adapters.Context) (string, error) {
	userID, err := ctx.Cookie(cookieName)
	if err != nil {
		return "", fmt.Errorf("error from `Cookie` method, package `context`: %v", err)
	}
	return userID, nil
}
