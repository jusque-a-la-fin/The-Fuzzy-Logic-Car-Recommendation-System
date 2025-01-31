package handlers

import (
	"car-recommendation-service/internal/search/adapters/controller"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	SearchController   controller.Search
	QuestionController controller.Question
}

// SearchForCars ответственен за получение списка автомобилей, чьи данные
// собираются из базы данных (PostgreSQL), и сохранение этого списка в базу данных под управлением Redis
func (hnd *SearchHandler) SearchForCars(ctx *gin.Context) {
	err := hnd.SearchController.GetSearchCars(ctx)
	if err != nil {
		log.Printf("error from `GetSearchCars` method, package `controller`: %v", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}
