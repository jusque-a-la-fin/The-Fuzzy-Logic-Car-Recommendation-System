package usecase

import (
	"car-recommendation-service/api/proto/generated/carsget"
	"car-recommendation-service/api/proto/generated/selection"
	"car-recommendation-service/api/proto/generated/storage"
	"car-recommendation-service/entities"
	slen "car-recommendation-service/internal/selection/entities"
	"car-recommendation-service/internal/shared/adapters"
	"context"
	"fmt"
	"strconv"
)

// SelectionInput содержит методы, которые обслуживают сервис,
// использующий нечеткий алгоритм для ранжирования автомобилей
type SelectionInput interface {
	PickPriorities(ctx adapters.Context)
	PickPrice(ctx adapters.Context)
	PickManufacturers(ctx adapters.Context)
	DoSelection(ctx adapters.Context, sessionID string, sln *slen.Selection) error
	PassSelectionCarsData(ctx adapters.Context, sessionID string) error
	PresentSelectionCarAd(ctx adapters.Context, sessionID, carID string) error
}

// SelectionOutput содержит методы, которые рендерят html-шаблоны
type SelectionOutput interface {
	ShowPriorities(ctx adapters.Context)
	ShowPrice(ctx adapters.Context)
	ShowManufacturers(ctx adapters.Context)
	ShowResultOfFuzzyAlgorithm(ctx adapters.Context, sessionID string, cars []*entities.Car)
	ShowSelectionCarAd(ctx adapters.Context, sessionID string, car *entities.Car)
}

type selectionUseCase struct {
	output          SelectionOutput
	storageClient   storage.StorageClient
	selectionClient selection.SelectionClient
}

func NewSelectionUseCase(snp SelectionOutput, stc storage.StorageClient, slc selection.SelectionClient) SelectionInput {
	return &selectionUseCase{snp, stc, slc}
}

// PickPriorities ответственен за формирование веб-страницы, предлагающей пользователю
// расставить приоритеты: "Комфорт", "Экономичность", "Безопасность", "Динамика", "Управляемость", по
// которым будут ранжироваться автомобили
func (slu *selectionUseCase) PickPriorities(ctx adapters.Context) {
	slu.output.ShowPriorities(ctx)
}

// PickPrice ответственен за формирование веб-страницы, предлагающей пользователю задать
// минимальную цену и максимальную цену автомобилей
func (slu *selectionUseCase) PickPrice(ctx adapters.Context) {
	slu.output.ShowPrice(ctx)
}

// PickManufacturers ответственен за формирование веб-страницы, предлагающей пользователю
// выбрать страны-производители автомобилей
func (slu *selectionUseCase) PickManufacturers(ctx adapters.Context) {
	slu.output.ShowManufacturers(ctx)
}

// DoSelection ответственен за получение списка автомобилей, чьи данные собираются из базы данных этого сервиса
// под управлением PostgreSQL, ранжирование этого списка с помощью нечеткого алгоритма и
// сохранение этого списка в базе данных под управлением Redis
// Входные параметры: sessionID - идентификатор сессии, sln - параметры поиска автомобилей
func (slu *selectionUseCase) DoSelection(ctx adapters.Context, sessionID string, sln *slen.Selection) error {
	scn := &carsget.CarsSelectionRequest{
		MinPrice:      sln.MinPrice,
		MaxPrice:      sln.MaxPrice,
		Manufacturers: sln.Manufacturers,
	}

	req := &selection.SelectionRequest{Selection: scn, SessionID: sessionID}
	_, err := slu.selectionClient.PerformSelection(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error from `PerformSelection` method, package `selection`: %v", err)
	}
	return nil
}

// PassSelectionCarsData ответственен за формирование веб-страницы,
// отображающей ранжированный с помощью нечеткого алгоритма список автомобилей
// Входной параметр: sessionID - идентификатор сессии
func (slu *selectionUseCase) PassSelectionCarsData(ctx adapters.Context, sessionID string) error {
	req := &storage.GetDataRequest{SessionID: sessionID, Survey: false}
	resp, err := slu.storageClient.GetData(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error from `GetData` method, package `storage`: %v", err)
	}
	cars := entities.GetCars(resp.Cars)
	slu.output.ShowResultOfFuzzyAlgorithm(ctx, sessionID, cars)
	return nil
}

// PresentSelectionCarAd ответственен за формирование веб-страницы конкретного автомобиля
// Входные параметры: sessionID - идентификатор сессии, carID - номер автомобиля в списке
func (slu *selectionUseCase) PresentSelectionCarAd(ctx adapters.Context, sessionID, carID string) error {
	req := &storage.GetDataRequest{SessionID: sessionID, Survey: false}
	resp, err := slu.storageClient.GetData(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error from `GetData` method, package `storage`: %v", err)
	}
	cars := entities.GetCars(resp.Cars)

	carIDint, err := strconv.Atoi(carID)
	if err != nil {
		return fmt.Errorf("error from `Atoi` function, package `strconv`: %v", err)
	}

	slu.output.ShowSelectionCarAd(ctx, sessionID, cars[carIDint-1])
	return nil
}
