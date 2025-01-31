package selectionservice

import (
	"car-recommendation-service/api/proto/generated/carsget"
	"car-recommendation-service/api/proto/generated/selection"
	"car-recommendation-service/api/proto/generated/storage"
	"car-recommendation-service/entities"
	"car-recommendation-service/internal/shared/format"
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
)

type selectionRepository struct {
	selection.UnimplementedSelectionServer
	carsClient    carsget.CarsClient
	storageClient storage.StorageClient
}

func NewSelectionRepository(csc carsget.CarsClient, stc storage.StorageClient) *selectionRepository {
	return &selectionRepository{
		selection.UnimplementedSelectionServer{},
		csc, stc,
	}
}

// PerformSelection ответственен за получение списка автомобилей, чьи данные собираются из базы данных этого сервиса
// под управлением PostgreSQL, ранжирование этого списка с помощью нечеткого алгоритма и
// сохранение этого списка в базе данных под управлением Redis
// Входные параметры: параметры поиска автомобилей, идентификатор сессии
func (slr *selectionRepository) PerformSelection(ctx context.Context, req *selection.SelectionRequest) (*emptypb.Empty, error) {
	var cars []*entities.Car
	resp, err := slr.carsClient.Select(context.Background(), req.Selection)
	if err != nil {
		return nil, fmt.Errorf("error from `Select` method, package `carsget`: %v", err)
	}
	cars = entities.GetCars(resp.Cars)

	sortedCars, err := generateResultOfFuzzyAlgorithm(cars, req.Selection.Priorities)
	if err != nil {
		return nil, fmt.Errorf("error from `generateResultOfFuzzyAlgorithm` function, package `selectionservice`: %v", err)
	}

	for _, car := range sortedCars {
		car.Offering.Price = fmt.Sprintf("%s ₽", format.FormatPrice(car.Offering.Price))
	}

	scrs := entities.GetCarsForStorage(sortedCars)
	req1 := &storage.LoadDataRequest{SessionID: req.SessionID, Cars: scrs, Survey: false}
	_, err = slr.storageClient.LoadData(context.Background(), req1)
	if err != nil {
		return nil, fmt.Errorf("error from `LoadData` method, package `storage`: %v", err)
	}
	return nil, nil
}
