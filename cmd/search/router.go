package main

import (
	"car-recommendation-service/api/proto/generated/carsget"
	"car-recommendation-service/api/proto/generated/storage"
	"car-recommendation-service/api/proto/generated/survey"
	"car-recommendation-service/internal/search/handlers"
	"car-recommendation-service/internal/search/registry"
	"car-recommendation-service/internal/shared/errors"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func MakeNewRouter(router *gin.Engine, storageCon, surveyCon, carsCon *grpc.ClientConn) *gin.Engine {
	storageClient := storage.NewStorageClient(storageCon)
	surveyClient := survey.NewSurveyClient(surveyCon)
	carsClient := carsget.NewCarsClient(carsCon)
	searchCont := registry.NewSearchController(storageClient, surveyClient, carsClient)
	questionCont := registry.NewQuestionController(storageClient, surveyClient)

	srch := &handlers.SearchHandler{
		SearchController:   searchCont,
		QuestionController: questionCont,
	}

	search := router.Group("/search")
	{
		ProcessErrors(search)
		search.POST("", srch.SearchForCars)
		search.POST("answer", srch.ReceiveAnswer)
		search.GET("guest/:guest", srch.ReceiveCarsList)
		search.GET("guest/:guest/carID/:carID", srch.GetCarPage)
	}
	return router
}

func ProcessErrors(search *gin.RouterGroup) {
	search.GET("error", func(ctx *gin.Context) {
		errors.ManageResponseStatusCodes(ctx)
	})

	search.POST("error", func(ctx *gin.Context) {
		errors.AcceptError(ctx, "/search/error")
	})
}
