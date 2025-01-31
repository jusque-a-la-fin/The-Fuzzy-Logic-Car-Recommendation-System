package controller

import (
	"car-recommendation-service/internal/selection/entities"
	usecase "car-recommendation-service/internal/selection/usecases"
	"car-recommendation-service/internal/shared/adapters"
	"fmt"
)

// Selection содержит методы, которые обслуживают сервис,
// использующий нечеткий алгоритм для ранжирования автомобилей
type Selection interface {
	ChoosePriorities(ctx adapters.Context)
	ChoosePrice(ctx adapters.Context)
	ChooseManufacturers(ctx adapters.Context)
	GetSelection(ctx adapters.Context) error
	DisplaySelectionCarAd(ctx adapters.Context, sessionID, carID string) error
	TransferSelectionCarsData(ctx adapters.Context, sessionID string) error
}

type selectionController struct {
	selectionUseCase usecase.SelectionInput
}

func NewSelectionController(sli usecase.SelectionInput) Selection {
	return &selectionController{sli}
}

// ChoosePriorities ответственен за формирование веб-страницы, предлагающей пользователю
// расставить приоритеты: "Комфорт", "Экономичность", "Безопасность", "Динамика", "Управляемость", по
// которым будут ранжироваться автомобили
func (slc *selectionController) ChoosePriorities(ctx adapters.Context) {
	slc.selectionUseCase.PickPriorities(ctx)
}

// ChoosePrice ответственен за формирование веб-страницы, предлагающей пользователю задать
// минимальную цену и максимальную цену автомобилей
func (slc *selectionController) ChoosePrice(ctx adapters.Context) {
	slc.selectionUseCase.PickPrice(ctx)
}

// ChooseManufacturers ответственен за формирование веб-страницы, предлагающей пользователю
// выбрать страны-производители автомобилей
func (slc *selectionController) ChooseManufacturers(ctx adapters.Context) {
	slc.selectionUseCase.PickManufacturers(ctx)
}

// GetSelection ответственен за получение списка автомобилей, чьи данные собираются из базы данных этого сервиса
// под управлением PostgreSQL, ранжирование этого списка с помощью нечеткого алгоритма и
// сохранение этого списка в базе данных под управлением Redis
func (slc *selectionController) GetSelection(ctx adapters.Context) error {
	type selection struct {
		Priorities    []string `json:"priorities"`
		MinPrice      string   `json:"minPrice"`
		MaxPrice      string   `json:"maxPrice"`
		Manufacturers []string `json:"manufacturers"`
		SessionID     string   `json:"sessionID"`
	}

	var sln selection
	if err := ctx.BindJSON(&sln); err != nil {
		return fmt.Errorf("error from `BindJSON` method, package `gin`: %v", err)
	}

	var stn = &entities.Selection{
		Priorities:    sln.Priorities,
		MinPrice:      sln.MinPrice,
		MaxPrice:      sln.MaxPrice,
		Manufacturers: sln.Manufacturers,
	}

	err := slc.selectionUseCase.DoSelection(ctx, sln.SessionID, stn)
	if err != nil {
		return fmt.Errorf("error from `DoSelection` method, package `usecase`: %v", err)
	}
	return nil
}

// TransferSelectionCarsData ответственен за формирование веб-страницы, отображающей ранжированный с помощью
// нечеткого алгоритма список автомобилей
// Входные параметры: sessionID - идентификатор сессии
func (slc *selectionController) TransferSelectionCarsData(ctx adapters.Context, sessionID string) error {
	err := slc.selectionUseCase.PassSelectionCarsData(ctx, sessionID)
	if err != nil {
		return fmt.Errorf("error from `PassSelectionCarsData` method, package `usecase`: %v", err)
	}
	return nil
}

// DisplaySelectionCarAd ответственен за формирование веб-страницы конкретного автомобиля
// Входные параметры: sessionID - идентификатор сессии, carID - номер автомобиля в списке
func (slc *selectionController) DisplaySelectionCarAd(ctx adapters.Context, sessionID, carID string) error {
	err := slc.selectionUseCase.PresentSelectionCarAd(ctx, sessionID, carID)
	if err != nil {
		return fmt.Errorf("error from `PresentSelectionCarAd` method, package `usecase`: %v", err)
	}
	return nil
}
