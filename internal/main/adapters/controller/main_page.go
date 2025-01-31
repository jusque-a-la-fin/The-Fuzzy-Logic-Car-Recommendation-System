package controller

import (
	usecases "car-recommendation-service/internal/main/usecases"
	"car-recommendation-service/internal/shared/adapters"
)

// MainPage содержит метод для формирования главной страницы
type MainPage interface {
	DisplayMainPage(ctx adapters.Context)
}

type mainPageController struct {
	mainPageUseCase usecases.MainPageInput
}

func NewMainPageController(mpi usecases.MainPageInput) MainPage {
	return &mainPageController{mpi}
}

// DisplayMainPage ответственен за формирование главной веб-страницы
func (mpc *mainPageController) DisplayMainPage(ctx adapters.Context) {
	mpc.mainPageUseCase.PresentMainPage(ctx)
}
