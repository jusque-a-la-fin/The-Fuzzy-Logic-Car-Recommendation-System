package handlers

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var reg = regexp.MustCompile(`/carID/[^/?]+`)

// ReceiveCarsList ответственен за формирование веб-страницы, отображающей список автомобилей без вопроса/с вопросом,
// ответ на который помогает "нечеткому алгоритму" обучиться
func (hnd *SearchHandler) ReceiveCarsList(ctx *gin.Context) {
	surveyParamStr := ctx.Query("survey")

	survey := false
	if surveyParamStr == "true" {
		survey = true
	}

	sessionID := ctx.Param("guest")

	err := hnd.SearchController.TransferSearchCarsData(ctx, sessionID, survey)
	if err != nil {
		log.Printf("error from `TransferSearchCarsData` method, package `controller`: %v", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
}

// GetCarPage ответственен за формирование веб-страницы конкретного автомобиля
func (hnd *SearchHandler) GetCarPage(ctx *gin.Context) {
	sessionID := ctx.Param("guest")
	thisCarID := ctx.Param("carID")
	surveyParamStr := ctx.Query("survey")
	survey := false
	if surveyParamStr == "true" {
		survey = true
	}

	path := reg.ReplaceAllString(ctx.Request.URL.String(), "")
	err := hnd.SearchController.DisplaySearchCarAd(ctx, sessionID, path, thisCarID, survey)
	if err != nil {
		log.Printf("error from `DisplaySearchCarAd` method, package `controller`: %v", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
}
