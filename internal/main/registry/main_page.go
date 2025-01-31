package registry

import (
	"car-recommendation-service/internal/main/adapters/controller"
	"car-recommendation-service/internal/main/adapters/presenter"
	usecases "car-recommendation-service/internal/main/usecases"
)

func NewMainPageController() controller.MainPage {
	nmp := presenter.NewMainPagePresenter()
	mpu := usecases.NewMainPageUseCase(nmp)
	return controller.NewMainPageController(mpu)
}
