package storageservice

import (
	"car-recommendation-service/api/proto/generated/storage"
	"car-recommendation-service/entities"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StorageRepository struct {
	storage.UnimplementedStorageServer
	// rdb - клиент Redis для подключения к NoSQL базе данных, хранящей данные выбранных пользователем автомобилей
	rdb *redis.Client
}

func NewStorageRepository(rdb *redis.Client) *StorageRepository {
	return &StorageRepository{
		UnimplementedStorageServer: storage.UnimplementedStorageServer{},
		rdb:                        rdb,
	}
}

// LoadData загружает в базу данных под управлением Redis идентификатор сессии, данные об автомобилях,
// полученные из базы данных под управлением PostgreSQL, данные опроса пользователей, если необходимо
func (slr *StorageRepository) LoadData(ctx context.Context, req *storage.LoadDataRequest) (*emptypb.Empty, error) {
	cars := entities.GetCars(req.Cars)
	carsJSON, err := json.Marshal(cars)
	if err != nil {
		return nil, fmt.Errorf("error from `Marshal` function, package `json`: %v", err)
	}

	if req.Survey {
		possibleAnswersStr := strings.Join(req.PossibleAnswers, ",")
		err = slr.rdb.HSet(ctx, "session:"+req.SessionID, "cars", string(carsJSON), "question", req.Question,
			"questionID", req.QuestionID, "possibleAnswers", possibleAnswersStr).Err()
		if err != nil {
			return nil, fmt.Errorf("error from `HSet` method, package `redis`: %v", err)
		}
	} else {
		err = slr.rdb.HSet(ctx, "session:"+req.SessionID, "cars", string(carsJSON)).Err()
		if err != nil {
			return nil, fmt.Errorf("error from `HSet` method, package `redis`: %v", err)
		}
	}
	return nil, nil
}

// GetData получает из базы данных под управлением Redis ранее полученные данные автомобилей из базы данных
// под управлением PostgreSQL, а также данные опроса, если необходимо
// Входные параметры: req.Survey - переменная-флаг: true - необходимо получить данные опроса пользователей,
// false - данные опроса не нужно получать, req.SessionID - идентификатор сессии
func (slr *StorageRepository) GetData(ctx context.Context, req *storage.GetDataRequest) (*storage.GetDataResponse, error) {
	carsJSON, err := slr.rdb.HGet(ctx, "session:"+req.SessionID, "cars").Result()
	if err != nil {
		return nil, fmt.Errorf("error from `HGet` method, package `redis`: %v", err)
	}

	var cars []*entities.Car
	if err := json.Unmarshal([]byte(carsJSON), &cars); err != nil {
		return nil, fmt.Errorf("error from `Unmarshal` function, package `json`: %v", err)
	}

	storageCars := entities.GetCarsForStorage(cars)

	if req.Survey {
		question, err := slr.rdb.HGet(ctx, "session:"+req.SessionID, "question").Result()
		if err != nil {
			return nil, fmt.Errorf("error from `HGet` method, package `redis`: %v", err)
		}

		questionID, err := slr.rdb.HGet(ctx, "session:"+req.SessionID, "questionID").Result()
		if err != nil {
			return nil, fmt.Errorf("error from `HGet` method, package `redis`: %v", err)
		}

		possibleAnswersStr, err := slr.rdb.HGet(ctx, "session:"+req.SessionID, "possibleAnswers").Result()
		if err != nil {
			return nil, fmt.Errorf("error from `HGet` method, package `redis`: %v", err)
		}

		possibleAnswers := strings.Split(possibleAnswersStr, ",")
		resp := storage.GetDataResponse{Cars: storageCars, Question: question, QuestionID: questionID, PossibleAnswers: possibleAnswers}
		return &resp, nil
	}

	resp := storage.GetDataResponse{Cars: storageCars}
	return &resp, nil
}
