package main

import (
	"car-recommendation-service/internal/services"
	srv "car-recommendation-service/internal/services/survey"
	"car-recommendation-service/internal/shared/config"
	"log"
)

func main() {
	configName := "survey"
	if err := config.SetupConfigForService(configName); err != nil {
		log.Fatalf("failed to load the config file %s", err.Error())
	}

	surveyDB, err := CreateNewDBForSurvey()
	if err != nil {
		log.Fatalln("can't connect to the survey database", err)
	}

	repo := srv.NewSurveyRepository(surveyDB)
	services.RunService("survey", repo)
}
