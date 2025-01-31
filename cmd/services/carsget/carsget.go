package main

import (
	"car-recommendation-service/internal/services"
	carsget "car-recommendation-service/internal/services/carsget"
	"car-recommendation-service/internal/shared/config"
	"log"
)

func main() {
	configName := "carsget"
	if err := config.SetupConfigForService(configName); err != nil {
		log.Fatalf("failed to load the config file %s", err.Error())
	}

	carsDB, err := CreateNewDBForCars()
	if err != nil {
		log.Fatalln("can't connect to vehicles database", err)
	}
	repo := carsget.NewCarsRepository(carsDB)
	services.RunService("carsget", repo)
}
