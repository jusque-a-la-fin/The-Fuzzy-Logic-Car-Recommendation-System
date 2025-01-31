package handlers

import (
	"car-recommendation-service/internal/selection/adapters/controller"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SelectionHandler struct {
	SelectionController controller.Selection
}

// SelectCars ответственен за получение списка автомобилей, чьи данные собираются из базы данных этого сервиса
// под управлением PostgreSQL, ранжирование этого списка с помощью нечеткого алгоритма и
// сохранение этого списка в базе данных под управлением Redis
func (hnd *SelectionHandler) SelectCars(ctx *gin.Context) {
	err := hnd.SelectionController.GetSelection(ctx)
	if err != nil {
		log.Printf("error from `GetSelection` method, package `controller`: %v", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

// GetCarsList ответственен за формирование веб-страницы, отображающей ранжированный с помощью
// нечеткого алгоритма список автомобилей
func (hnd *SelectionHandler) GetCarsList(ctx *gin.Context) {
	sessionID := ctx.Param("guest")
	if sessionID != "" {
		err := hnd.SelectionController.TransferSelectionCarsData(ctx, sessionID)
		if err != nil {
			log.Printf("error from `TransferSelectionCarsData` method, package `controller`: %v", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
	}
}

// GetCarPage ответственен за формирование веб-страницы конкретного автомобиля
func (hnd *SelectionHandler) GetCarPage(ctx *gin.Context) {
	sessionID := ctx.Param("guest")
	thisCarID := ctx.Param("carID")

	if sessionID != "" && thisCarID != "" {
		err := hnd.SelectionController.DisplaySelectionCarAd(ctx, sessionID, thisCarID)
		if err != nil {
			log.Printf("error from `DisplaySelectionCarAd` method, package `controller`: %v", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
	}
}
