package controller

import (
	"car-recommendation-service/internal/search/entities"
	usecase "car-recommendation-service/internal/search/usecases/usecases"
	"car-recommendation-service/internal/shared/adapters"
	"fmt"
)

// Search содержит методы для обслуживания сервиса,
// осуществляющего обычный поиск с фильтрами без ранжирования и собирающего
// ответы пользователей для "обучения" нечеткого алгоритма
type Search interface {
	GetSearchCars(ctx adapters.Context) error
	TransferSearchCarsData(ctx adapters.Context, sessionID string, survey bool) error
	DisplaySearchCarAd(ctx adapters.Context, sessionID, path, carID string, survey bool) error
}

type searchController struct {
	searchUseCase usecase.SearchInput
}

func NewSearchController(shi usecase.SearchInput) Search {
	return &searchController{shi}
}

// GetSearchCars ответственен за получение списка автомобилей, чьи данные
// собираются из базы данных (PostgreSQL), и сохранение этого списка в базу данных под управлением Redis
func (src *searchController) GetSearchCars(ctx adapters.Context) error {
	type search struct {
		SessionID string          `json:"sessionID"`
		IsKnown   bool            `json:"isknown"`
		Form      entities.Search `json:"form"`
	}

	srch := new(search)
	if err := ctx.Bind(&srch); err != nil {
		return fmt.Errorf("error from `Bind` method, package `gin`: %v", err)
	}

	err := src.searchUseCase.GetCars(ctx, srch.SessionID, &srch.Form, srch.IsKnown)
	if err != nil {
		return fmt.Errorf("error from `GetCars` method, package `usecase`: %v", err)
	}
	return nil
}

// TransferSearchCarsData ответственен за формирование веб-страницы, отображающей список автомобилей без вопроса/с вопросом,
// ответ на который помогает "нечеткому алгоритму" обучиться.
// Входные параметры: sessionID - идентификатор сессии, survey - переменная-флаг: true - вопрос будет показан на странице,
// false - вопрос не будет показан.
func (src *searchController) TransferSearchCarsData(ctx adapters.Context, sessionID string, survey bool) error {
	err := src.searchUseCase.PassSearchCarsData(ctx, sessionID, survey)
	if err != nil {
		return fmt.Errorf("error from `PassSearchCarsData` method, package `usecase`: %v", err)
	}
	return nil
}

// DisplaySearchCarAd ответственен за формирование веб-страницы конкретного автомобиля
// Входные параметры: sessionID - идентификатор сессии, path - путь из url, carID - номер автомобиля в списке,
// survey - переменная-флаг: true - вопрос будет показан на странице, false - вопрос не будет показан.
func (src *searchController) DisplaySearchCarAd(ctx adapters.Context, sessionID, path, carID string, survey bool) error {
	err := src.searchUseCase.PresentSearchCarAd(ctx, sessionID, path, carID, survey)
	if err != nil {
		return fmt.Errorf("error from `PresentSearchCarAd` method, package `usecase`: %v", err)
	}
	return nil
}
