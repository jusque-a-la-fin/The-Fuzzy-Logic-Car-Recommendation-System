package main

import (
	"car-recommendation-service/api/proto/generated/carsget"
	"car-recommendation-service/api/proto/generated/storage"
	"car-recommendation-service/internal/services"
	selectionservice "car-recommendation-service/internal/services/selection"
	"car-recommendation-service/internal/shared/config"
	"log"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	configName := "selection_service"
	if err := config.SetupConfigForService(configName); err != nil {
		log.Fatalf("failed to load the config file %s", err.Error())
	}

	carsCon, err := grpc.NewClient(
		viper.GetString("target1"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("can't create a new gRPC `channel`", err)
	}
	defer carsCon.Close()

	storageCon, err := grpc.NewClient(
		viper.GetString("target2"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("can't create a new gRPC `channel`", err)
	}
	defer storageCon.Close()

	carsClient := carsget.NewCarsClient(carsCon)
	storageClient := storage.NewStorageClient(storageCon)

	repo := selectionservice.NewSelectionRepository(carsClient, storageClient)
	services.RunService("selection", repo)
}
