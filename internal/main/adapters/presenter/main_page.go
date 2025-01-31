package presenter

import (
	usecases "car-recommendation-service/internal/main/usecases"
	"car-recommendation-service/internal/shared/adapters"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type mainPagePresenter struct{}

func NewMainPagePresenter() usecases.MainPageOutput {
	return &mainPagePresenter{}
}

// ShowMainPage рендерит главную страницу
func (m *mainPagePresenter) ShowMainPage(ctx adapters.Context) {
	sessionID := uuid.New().String()
	htmlFileName := "main_page.html"
	ctx.HTML(http.StatusOK, htmlFileName, gin.H{"sessionID": sessionID})
}
