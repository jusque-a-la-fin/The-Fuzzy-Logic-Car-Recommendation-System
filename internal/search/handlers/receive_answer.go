package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReceiveAnswer получает ответ от пользователя, необходимый для "обучения" нечеткого алгоритма, и
// записывает его в базу данных под управлением PostgreSQL
func (hnd *SearchHandler) ReceiveAnswer(ctx *gin.Context) {
	err := hnd.QuestionController.GetAnswer(ctx)
	if err != nil {
		log.Printf("error from `GetAnswer` method, package `controller`: %v", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}
