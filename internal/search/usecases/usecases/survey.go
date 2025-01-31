package usecase

import (
	"car-recommendation-service/api/proto/generated/storage"
	"car-recommendation-service/api/proto/generated/survey"
	"car-recommendation-service/internal/search/entities"
	"car-recommendation-service/internal/shared/adapters"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type SurveyInput interface {
	PickQuestion(ctx adapters.Context, isKnown bool) (string, string, []string, error)
	GetAnswer(ctx adapters.Context, questionID, answer string) error
}

type surveyUseCase struct {
	userUseCase   UserInput
	output        SearchOutput
	storageClient storage.StorageClient
	surveyClient  survey.SurveyClient
}

func NewSurveyUseCase(urp UserInput, sep SearchOutput, src storage.StorageClient, svc survey.SurveyClient) SurveyInput {
	return &surveyUseCase{urp, sep, src, svc}
}

// PickQuestion выбирает вопрос для пользователя
// Ответ на этот вопрос необходим для "обучения" нечеткого алгоритма
// Входной параметр: isKnown - переменная-флаг: true - пользователь уже отвечал на вопрос,
// false - пользователь новый или не отвечал ни на один вопрос
func (suu *surveyUseCase) PickQuestion(ctx adapters.Context, isKnown bool) (string, string, []string, error) {
	user := entities.User{}
	var err error
	if !isKnown {
		user.ID = uuid.New().String()
		suu.userUseCase.SetUserID(ctx, user.ID)

	} else {
		user.ID, err = suu.userUseCase.GetUserID(ctx)
		if err != nil {
			return "", "", nil, fmt.Errorf("error from `GetUserID` method, package `usecase`: %v", err)
		}
	}

	if user.ID != "" {
		req := &survey.ChooseQuestionRequest{UserID: user.ID}
		resp, err := suu.surveyClient.ChooseQuestion(context.Background(), req)
		if err != nil {
			return "", "", nil, fmt.Errorf("error from `ChooseQuestion` method, package `survey`: %v", err)
		}
		return resp.QuestionText, resp.QuestionID, resp.PossibleAnswers, nil
	}
	return "", "", nil, fmt.Errorf("user.ID is ''(empty string). Neither SetUserID nor GetUserID obtained non-empty valid user.ID")
}

// GetAnswer принимает ответ пользователя. Этот ответ необходим для "обучения" нечеткого алгоритма
// Входные параметры: questionID - идентификатор вопроса, answer - ответ
func (suu *surveyUseCase) GetAnswer(ctx adapters.Context, questionID, answer string) error {
	user := entities.User{}
	var err error
	user.ID, err = suu.userUseCase.GetUserID(ctx)
	if err != nil {
		return fmt.Errorf("error from `GetUserID` method, package `usecase`: %v", err)
	}

	qtn := entities.Question{ID: questionID, Answer: answer}

	req := &survey.InsertAnswerRequest{UserID: user.ID, QuestionID: qtn.ID, Answer: qtn.Answer}
	_, err = suu.surveyClient.InsertAnswer(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error from `InsertAnswer` method, package `survey`: %v", err)
	}
	return nil
}
