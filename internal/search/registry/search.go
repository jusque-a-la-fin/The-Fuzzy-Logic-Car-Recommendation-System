package registry

import (
	"car-recommendation-service/api/proto/generated/carsget"
	"car-recommendation-service/api/proto/generated/storage"
	"car-recommendation-service/api/proto/generated/survey"
	"car-recommendation-service/internal/search/adapters/controller"
	"car-recommendation-service/internal/search/adapters/gateway"
	"car-recommendation-service/internal/search/adapters/presenter"
	usecase "car-recommendation-service/internal/search/usecases/usecases"
)

func NewSearchController(stc storage.StorageClient, svc survey.SurveyClient, csc carsget.CarsClient) controller.Search {
	nuu := usecase.NewUserUseCase(gateway.NewUserRepository())
	nsp := presenter.NewSearchPresenter()
	nsu := usecase.NewSearchUseCase(
		nuu,
		usecase.NewSurveyUseCase(
			nuu, nsp, stc, svc,
		),
		nsp,
		stc,
		csc,
	)
	return controller.NewSearchController(nsu)
}
