package usecase

import (
	"car-recommendation-service/api/proto/generated/carsget"
	storage "car-recommendation-service/api/proto/generated/storage"
	"car-recommendation-service/internal/shared/adapters"
	"car-recommendation-service/internal/shared/format"
	"context"
	"fmt"
	"strconv"

	"car-recommendation-service/entities"
	scen "car-recommendation-service/internal/search/entities"
)

// SearchInput содержит методы для обслуживания сервиса,
// осуществляющего обычный поиск с фильтрами без ранжирования и собирающего
// ответы пользователей для "обучения" нечеткого алгоритма
type SearchInput interface {
	GetCars(ctx adapters.Context, sessionID string, search *scen.Search, isKnown bool) error
	PassSearchCarsData(ctx adapters.Context, sessionID string, survey bool) error
	PresentSearchCarAd(ctx adapters.Context, sessionID, path, carID string, survey bool) error
}

// SearchOutput содержит методы, которые рендерят html-шаблоны
type SearchOutput interface {
	ShowCarsWithSurvey(ctx adapters.Context, question, questionID string, cars []*entities.Car, possibleAnswers []string)
	ShowCars(ctx adapters.Context, sessionID string, cars []*entities.Car)
	ShowSearchCarAd(ctx adapters.Context, car *entities.Car, path string)
}

type searchUseCase struct {
	userUseCase   UserInput
	surveyUseCase SurveyInput
	output        SearchOutput
	storageClient storage.StorageClient
	carsClient    carsget.CarsClient
}

func NewSearchUseCase(usi UserInput, svi SurveyInput, sho SearchOutput, stc storage.StorageClient, csc carsget.CarsClient) SearchInput {
	return &searchUseCase{usi, svi, sho, stc, csc}
}

// GetCars ответственен за получение списка автомобилей, чьи данные
// собираются из базы данных (PostgreSQL), и сохранение этого списка в базу данных под управлением Redis
// Входной параметр: sessionID - идентификатор сессии, search - параметры поиска автомобилей,
// isKnown - переменная-флаг: true - пользователь уже отвечал на вопрос,
// false - пользователь новый или не отвечал ни на один вопрос
func (src *searchUseCase) GetCars(ctx adapters.Context, sessionID string, search *scen.Search, isKnown bool) error {
	scn := &carsget.CarsSearchRequest{
		Make:         search.Make,
		Model:        search.Model,
		Gearbox:      search.Gearbox,
		MinPrice:     search.LowPriceLimit,
		MaxPrice:     search.HighPriceLimit,
		Drive:        search.Drive,
		EarliestYear: search.EarliestYear,
		LatestYear:   search.LatestYear,
		Fuel:         search.Fuel,
		IsNewCar:     search.IsNewCar,
	}

	resp, err := src.carsClient.Search(context.Background(), scn)
	if err != nil {
		return fmt.Errorf("error from `Search` method, package `carsget`: %v", err)
	}

	question, questionID, possibleAnswers, err := src.surveyUseCase.PickQuestion(ctx, isKnown)
	if err != nil {
		return fmt.Errorf("error from `PickQuestion` method, package `usecase`: %v", err)
	}

	for _, car := range resp.Cars {
		car.Offering.Price = fmt.Sprintf("%s ₽", format.FormatPrice(car.Offering.Price))
	}

	req := &storage.LoadDataRequest{SessionID: sessionID, Cars: resp.Cars, Question: question,
		QuestionID: questionID, PossibleAnswers: possibleAnswers, Survey: true}
	_, err = src.storageClient.LoadData(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error from `LoadData` method, package `storage`: %v", err)
	}
	return nil
}

// PassSearchCarsData ответственен за формирование веб-страницы, отображающей список автомобилей с вопросом,
// ответ на который помогает "нечеткому алгоритму" обучиться
// Входные параметры: sessionID - идентификатор сессии, survey - переменная-флаг: true - на странице нужно показать опрос,
// false - на странице не нужно показывать опрос
func (src *searchUseCase) PassSearchCarsData(ctx adapters.Context, sessionID string, survey bool) error {
	var cars []*entities.Car
	req := &storage.GetDataRequest{SessionID: sessionID, Survey: survey}
	resp, err := src.storageClient.GetData(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error from `GetData` method, package `storage`: %v", err)
	}
	cars = entities.GetCars(resp.Cars)

	if survey {
		src.output.ShowCarsWithSurvey(ctx, resp.Question, resp.QuestionID, cars, resp.PossibleAnswers)
	} else {
		src.output.ShowCars(ctx, sessionID, cars)
	}
	return nil
}

// PresentSearchCarAd ответственен за формирование веб-страницы конкретного автомобиля
// Входные параметры: sessionID - идентификатор сессии, path - путь из url, carID - номер автомобиля в списке,
// survey - переменная-флаг: true - на странице всех автомобилей нужно показать опрос,
// false - на странице всех автомобилей не нужно показывать опрос
func (srс *searchUseCase) PresentSearchCarAd(ctx adapters.Context, sessionID, path, carID string, survey bool) error {
	var cars []*entities.Car
	req := &storage.GetDataRequest{SessionID: sessionID, Survey: survey}
	resp, err := srс.storageClient.GetData(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error from `GetData` method, package `storage`: %v", err)
	}
	cars = entities.GetCars(resp.Cars)

	carIDint, err := strconv.Atoi(carID)
	if err != nil {
		return fmt.Errorf("error from `Atoi` function, package `strconv`: %v", err)
	}

	srс.output.ShowSearchCarAd(ctx, cars[carIDint-1], path)
	return nil
}
