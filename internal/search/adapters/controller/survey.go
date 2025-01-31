package controller

import (
	usecase "car-recommendation-service/internal/search/usecases/usecases"
	"car-recommendation-service/internal/shared/adapters"
	"fmt"
)

type Question interface {
	GetAnswer(ctx adapters.Context) error
}

type questionController struct {
	questionUseCase usecase.SurveyInput
}

func NewQuestionController(qni usecase.SurveyInput) Question {
	return &questionController{qni}
}

// GetAnswer получает ответ от пользователя, необходимый для "обучения" нечеткого алгоритма, и
// записывает его в базу данных под управлением PostgreSQL.
// Пользователь дает этот ответ в опросе на странице, где показан список автомобилей
func (qnc *questionController) GetAnswer(ctx adapters.Context) error {
	type questionID struct {
		// Answer - ответ
		Answer string `json:"answer"`
		// QuestionID - идентификатор вопроса
		QuestionID string `json:"questionID"`
	}

	qtn := new(questionID)
	if err := ctx.Bind(&qtn); err != nil {
		return fmt.Errorf("error from `Bind` method, package `gin`: %v", err)
	}

	err := qnc.questionUseCase.GetAnswer(ctx, qtn.QuestionID, qtn.Answer)
	if err != nil {
		return fmt.Errorf("error from `GetAnswer` method, package `usecase`: %v", err)
	}
	return nil
}
