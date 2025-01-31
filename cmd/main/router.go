package main

import (
	"car-recommendation-service/internal/main/registry"
	"car-recommendation-service/internal/shared/errors"

	"github.com/gin-gonic/gin"
)

func MakeNewRouter(router *gin.Engine) *gin.Engine {
	main := router.Group("/main")
	{
		ProcessErrors(main)
		controller := registry.NewMainPageController()
		main.GET("", func(ctx *gin.Context) {
			controller.DisplayMainPage(ctx)
		})
	}
	return router
}

func ProcessErrors(main *gin.RouterGroup) {
	main.GET("error", func(ctx *gin.Context) {
		errors.ManageResponseStatusCodes(ctx)
	})

	main.POST("error", func(ctx *gin.Context) {
		errors.AcceptError(ctx, "/main/error")
	})
}
