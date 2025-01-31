package registry

import (
	"car-recommendation-service/api/proto/generated/storage"
	"car-recommendation-service/api/proto/generated/survey"
	"car-recommendation-service/internal/search/adapters/controller"
	"car-recommendation-service/internal/search/adapters/gateway"
	"car-recommendation-service/internal/search/adapters/presenter"
	usecase "car-recommendation-service/internal/search/usecases/usecases"
)

func NewQuestionController(stc storage.StorageClient, svc survey.SurveyClient) controller.Question {
	nuu := usecase.NewUserUseCase(gateway.NewUserRepository())
	nsp := presenter.NewSearchPresenter()
	nsu := usecase.NewSurveyUseCase(nuu, nsp, stc, svc)
	return controller.NewQuestionController(nsu)
}
