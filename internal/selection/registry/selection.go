package registry

import (
	"car-recommendation-service/api/proto/generated/selection"
	"car-recommendation-service/api/proto/generated/storage"
	"car-recommendation-service/internal/selection/adapters/controller"
	"car-recommendation-service/internal/selection/adapters/presenter"
	usecase "car-recommendation-service/internal/selection/usecases"
)

func NewSelectionController(stc storage.StorageClient, slc selection.SelectionClient) controller.Selection {
	nsu := usecase.NewSelectionUseCase(
		presenter.NewSelectionPresenter(),
		stc,
		slc,
	)
	return controller.NewSelectionController(nsu)
}
