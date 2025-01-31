package usecase

import "car-recommendation-service/internal/shared/adapters"

// MainPageInput содержит метод, который формирует главную страницу
type MainPageInput interface {
	PresentMainPage(ctx adapters.Context)
}

// MainPageOutput содержит метод, который рендерит html-шаблон
type MainPageOutput interface {
	ShowMainPage(ctx adapters.Context)
}

type mainPageUseCase struct {
	output MainPageOutput
}

func NewMainPageUseCase(mpo MainPageOutput) MainPageInput {
	return &mainPageUseCase{mpo}
}

// PresentMainPage ответственен за формирование главной веб-страницы
func (mpu *mainPageUseCase) PresentMainPage(ctx adapters.Context) {
	mpu.output.ShowMainPage(ctx)
}
