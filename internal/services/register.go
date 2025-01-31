package services

import (
	"car-recommendation-service/api/proto/generated/carsget"
	"car-recommendation-service/api/proto/generated/selection"
	"car-recommendation-service/api/proto/generated/storage"
	"car-recommendation-service/api/proto/generated/survey"
	"fmt"

	"google.golang.org/grpc"
)

func RegisterService(server *grpc.Server, service string, repository interface{}) error {
	switch service {
	case "storage":
		if repo, ok := repository.(storage.StorageServer); ok {
			storage.RegisterStorageServer(server, repo)
		} else {
			return fmt.Errorf("error: StorageServer doesn't exist")
		}

	case "selection":
		if repo, ok := repository.(selection.SelectionServer); ok {
			selection.RegisterSelectionServer(server, repo)
		} else {
			return fmt.Errorf("error: SelectionServer doesn't exist")
		}

	case "carsget":
		if repo, ok := repository.(carsget.CarsServer); ok {
			carsget.RegisterCarsServer(server, repo)
		} else {
			return fmt.Errorf("error: VehiclesDBServer doesn't exist")
		}
	case "survey":
		if repo, ok := repository.(survey.SurveyServer); ok {
			survey.RegisterSurveyServer(server, repo)
		} else {
			return fmt.Errorf("error: SurveyServer doesn't exist")
		}
	}
	return nil
}
