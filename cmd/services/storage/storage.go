package storage

import (
	services "car-recommendation-service/internal/services"
	storageservice "car-recommendation-service/internal/services/storage"
	"car-recommendation-service/internal/shared/config"
	"log"
)

func MakeStorage(configName string) {
	if err := config.SetupConfigForService(configName); err != nil {
		log.Fatalf("failed to load the config file %s", err.Error())
	}

	rdb, err := CreateNewRDB()
	if err != nil {
		log.Fatalf("can't connect to the redis db %s", err.Error())
	}
	repo := storageservice.NewStorageRepository(rdb)
	services.RunService("storage", repo)
}
