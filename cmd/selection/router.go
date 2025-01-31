package main

import (
	"car-recommendation-service/api/proto/generated/selection"
	"car-recommendation-service/api/proto/generated/storage"
	"car-recommendation-service/internal/selection/handlers"
	"car-recommendation-service/internal/selection/registry"
	"car-recommendation-service/internal/shared/errors"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func MakeNewRouter(router *gin.Engine, storageCon, selectionCon *grpc.ClientConn) *gin.Engine {
	storageClient := storage.NewStorageClient(storageCon)
	selectionClient := selection.NewSelectionClient(selectionCon)
	cont := registry.NewSelectionController(storageClient, selectionClient)

	seln := &handlers.SelectionHandler{
		SelectionController: cont,
	}

	selection := router.Group("/selection")
	{
		ProcessErrors(selection)
		selection.GET("priorities/guest/:guest", func(ctx *gin.Context) {
			cont.ChoosePriorities(ctx)
		})

		selection.GET("price/guest/:guest", func(ctx *gin.Context) {
			cont.ChoosePrice(ctx)
		})

		selection.GET("manufacturers/guest/:guest", func(ctx *gin.Context) {
			cont.ChooseManufacturers(ctx)
		})

		selection.POST("", seln.SelectCars)
		selection.GET("guest/:guest", seln.GetCarsList)
		selection.GET("guest/:guest/carID/:carID", seln.GetCarPage)
	}
	return router
}

func ProcessErrors(selection *gin.RouterGroup) {
	selection.GET("error", func(ctx *gin.Context) {
		errors.ManageResponseStatusCodes(ctx)
	})

	selection.POST("error", func(ctx *gin.Context) {
		errors.AcceptError(ctx, "/selection/error")
	})
}
