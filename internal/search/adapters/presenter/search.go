package presenter

import (
	"car-recommendation-service/entities"
	usecase "car-recommendation-service/internal/search/usecases/usecases"
	"car-recommendation-service/internal/shared/adapters"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const carsFile string = "offer.html"
const carFile string = "car_card.html"

type searchPresenter struct{}

func NewSearchPresenter() usecase.SearchOutput {
	return &searchPresenter{}
}

type SearchData struct {
	Survey          bool
	Cars            []*entities.Car
	Quantity        int
	Indexes         []int
	Question        string
	QuestionID      string
	PossibleAnswers []string
	HomePageLink    string
	ErrorLink       string
}

// ShowCarsWithSurvey рендерит страницу, отображающую список автомобилей с вопросом для пользователя,
// ответ на который помогает нечеткому алгоритму "обучиться"
// Входные параметры: question - вопрос для пользователя, questionID - идентификатор вопроса, cars - автомобили,
// possibleAnswers - варианты ответа
func (s *searchPresenter) ShowCarsWithSurvey(ctx adapters.Context, question, questionID string, cars []*entities.Car,
	possibleAnswers []string) {

	indexes := make([]int, len(cars))
	for i := range cars {
		indexes[i] = i + 1
	}

	homePageLink := "http://localhost:8080/main"
	errorLink := "http://localhost:8081/search/error?code="

	data := SearchData{
		Survey:          true,
		Cars:            cars,
		Quantity:        len(cars),
		Indexes:         indexes,
		Question:        question,
		QuestionID:      questionID,
		PossibleAnswers: possibleAnswers,
		HomePageLink:    homePageLink,
		ErrorLink:       errorLink,
	}
	ctx.HTML(http.StatusOK, carsFile, data)
}

// ShowCars рендерит страницу, отображающую список автомобилей без вопроса для пользователя
// Входные параметры: sessionID - идентификатор сессии, cars - автомобили
func (s *searchPresenter) ShowCars(ctx adapters.Context, sessionID string, cars []*entities.Car) {
	indexes := make([]int, len(cars))
	for i := range cars {
		indexes[i] = i + 1
	}
	homePageLink := "http://localhost:8080/main"
	errorLink := "http://localhost:8081/search/error/code/"

	data := SearchData{
		Cars:         cars,
		Quantity:     len(cars),
		Indexes:      indexes,
		HomePageLink: homePageLink,
		ErrorLink:    errorLink,
	}

	ctx.HTML(http.StatusOK, carsFile, data)
}

// ShowSearchCarAd рендерит страницу конкретного автомобиля
// Входные параметры: car - автомобиль, path - путь из url
func (s *searchPresenter) ShowSearchCarAd(ctx adapters.Context, car *entities.Car, path string) {
	link := fmt.Sprintf("http://localhost:8081%s", path)
	ctx.HTML(http.StatusOK, carFile, gin.H{"Car": car, "Link": link})
}
