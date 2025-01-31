package survey

import (
	"car-recommendation-service/api/proto/generated/survey"
	context "context"
	"database/sql"
	"fmt"
	"strconv"

	"golang.org/x/exp/rand"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type surveyRepository struct {
	survey.UnimplementedSurveyServer
	// surveyDB - клиент для подключения к базе данных под управлением PostgreSQL,
	// хранящей информацию, связанную с опросом пользователей
	surveyDB *sql.DB
}

func NewSurveyRepository(surveyDB *sql.DB) *surveyRepository {
	return &surveyRepository{
		surveyDB: surveyDB,
	}
}

// ChooseQuestion выбирает из базы данных вопрос, на который пользователь еще не отвечал
// Входной параметр: req.UserID - идентификатор пользователя
func (srr *surveyRepository) ChooseQuestion(ctx context.Context, req *survey.ChooseQuestionRequest) (*survey.ChooseQuestionResponse, error) {
	questionIDs, err := srr.GetIdsOfUnansweredQuestions(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("error from `GetIdsOfUnansweredQuestions` method, package `survey`: %v", err)
	}

	randIndex := rand.Intn(len(questionIDs))
	questionID := questionIDs[randIndex]

	questionText, possibleAnswers, err := srr.GetQuestion(questionID)
	if err != nil {
		return nil, fmt.Errorf("error from `GetQuestion` method, package `survey`: %v", err)
	}

	resp := &survey.ChooseQuestionResponse{QuestionText: questionText, QuestionID: questionID, PossibleAnswers: possibleAnswers}
	return resp, nil
}

// GetIdsOfUnansweredQuestions получает из базы данных id вопросов, на которые пользователь ещё не отвечал
// Входной параметр: userID - идентификатор пользователя
func (srr *surveyRepository) GetIdsOfUnansweredQuestions(userID string) ([]string, error) {
	query := `
        SELECT id
        FROM questions
        WHERE id NOT IN (
          SELECT question_id
          FROM user_responses
          WHERE user_id = (
            SELECT id
            FROM users
            WHERE guest_id = $1
          )
        );
    `

	rows, err := srr.surveyDB.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error from `Query` method, package `sql`: %v", err)
	}
	defer rows.Close()

	var questionIDs []string
	var id string
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		questionIDs = append(questionIDs, id)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error from `Err` method, package `sql`: %v", err)
	}
	return questionIDs, nil
}

// GetQuestion получает вопрос для пользователя из базы данных
// Входной параметр: questionID - идентификатор вопроса
func (srr *surveyRepository) GetQuestion(questionID string) (string, []string, error) {
	query := `
        SELECT question
        FROM questions
        WHERE id = $1;
    `

	var questionText string
	err := srr.surveyDB.QueryRow(query, questionID).Scan(&questionText)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil, fmt.Errorf("error: sql.ErrNoRows from `QueryRow` method, package `sql`: %v", sql.ErrNoRows)
		} else {
			return "", nil, fmt.Errorf("error from `QueryRow` method, package `sql`: %v", err)
		}
	}

	query = `
        SELECT possible_answer
        FROM possible_answers
        WHERE question_id = $1;
    `

	rows, err := srr.surveyDB.Query(query, questionID)
	if err != nil {
		return "", nil, fmt.Errorf("error from `Query` method, package `sql`: %v", err)
	}
	defer rows.Close()

	var possibleAnswers []string
	var possibleAnswer string
	for rows.Next() {
		err := rows.Scan(&possibleAnswer)
		if err != nil {
			return "", nil, fmt.Errorf("error from `Scan` method, package `sql`: %v", err)
		}
		possibleAnswers = append(possibleAnswers, possibleAnswer)
	}
	if err := rows.Err(); err != nil {
		return "", nil, fmt.Errorf("error from `Err` method, package `sql`: %v", err)
	}

	return questionText, possibleAnswers, nil
}

// InsertAnswer записывает ответ пользователя в базу данных
// Входные параметры: req.UserID - идентификатор пользователя,
// req.QuestionID - идентификатор вопроса, req.Answer - ответ пользователя
func (srr *surveyRepository) InsertAnswer(ctx context.Context, req *survey.InsertAnswerRequest) (*emptypb.Empty, error) {
	var exists bool
	err := srr.surveyDB.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE guest_id = $1)", req.UserID).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("error from `QueryRow` method, package `sql`: %v", err)
	}

	if !exists {
		sqlQuery := `
        WITH new_user AS (
            INSERT INTO users (guest_id)
            SELECT CAST($1 AS VARCHAR(160))
            WHERE NOT EXISTS (
                SELECT 1 FROM users WHERE guest_id = $1
        )
            RETURNING id
        ),
        upsert_response AS (
            INSERT INTO user_responses (user_id, question_id, answer)
            SELECT new_user.id, $2, $3
            FROM new_user
		    ON CONFLICT ON CONSTRAINT unique_user_question_responses DO 
			UPDATE SET answer = EXCLUDED.answer
            RETURNING *
        )
        SELECT * FROM upsert_response`

		two, err := strconv.Atoi(req.QuestionID)
		if err != nil {
			return nil, fmt.Errorf("error from `Atoi` function, package `strconv`: %v", err)
		}

		_, err = srr.surveyDB.Exec(sqlQuery, req.UserID, two, req.Answer)
		if err != nil {
			return nil, fmt.Errorf("error while inserting answer for non-existing user, error from `Exec` method, package `sql`: %v", err)
		}

	} else {
		sqlQuery := `
          WITH existing_user AS (
	          SELECT id FROM users WHERE guest_id = $1
          ),
          upsert_response AS (
	          INSERT INTO user_responses (user_id, question_id, answer)
	          SELECT existing_user.id, $2, $3
	          FROM existing_user
	          ON CONFLICT ON CONSTRAINT unique_user_question_responses DO 
			  UPDATE SET answer = EXCLUDED.answer
	          RETURNING *
           ) 
           SELECT * FROM upsert_response`

		_, err := srr.surveyDB.Exec(sqlQuery, req.UserID, req.QuestionID, req.Answer)
		if err != nil {
			return nil, fmt.Errorf("error while inserting answer for existing user, error from `Exec` method, package `sql`: %v", err)
		}
	}
	return nil, nil
}
