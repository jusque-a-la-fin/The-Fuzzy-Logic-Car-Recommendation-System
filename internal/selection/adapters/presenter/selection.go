package presenter

import (
	"car-recommendation-service/entities"
	usecases "car-recommendation-service/internal/selection/usecases"
	"car-recommendation-service/internal/shared/adapters"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type selectionPresenter struct{}

func NewSelectionPresenter() usecases.SelectionOutput {
	return &selectionPresenter{}
}

type SelectionData struct {
	Survey       bool
	Cars         []*entities.Car
	Quantity     int
	Indexes      []int
	HomePageLink string
	ErrorLink    string
}

// следующие методы рендерят различные веб-страницы
func (slp *selectionPresenter) ShowPriorities(ctx adapters.Context) {
	ctx.HTML(http.StatusOK, "priorities.html", nil)
}

func (slp *selectionPresenter) ShowPrice(ctx adapters.Context) {
	ctx.HTML(http.StatusOK, "price.html", nil)
}

func (slp *selectionPresenter) ShowManufacturers(ctx adapters.Context) {
	ctx.HTML(http.StatusOK, "manufacturers.html", nil)
}

// ShowResultOfFuzzyAlgorithm рендерит страницу, отображающую ранжированный с помощью нечеткого алгоритма список автомобилей
// Входные параметры: sessionID - идентификатор сессии, cars - автомобили
func (slp *selectionPresenter) ShowResultOfFuzzyAlgorithm(ctx adapters.Context, sessionID string, cars []*entities.Car) {
	indexes := make([]int, len(cars))
	for i := range cars {
		indexes[i] = i + 1
	}

	homePageLink := "http://localhost:8080/main"
	errorLink := "http://localhost:8082/selection/error?code="
	data := SelectionData{
		Survey:       false,
		Cars:         cars,
		Quantity:     len(cars),
		Indexes:      indexes,
		HomePageLink: homePageLink,
		ErrorLink:    errorLink,
	}

	ctx.HTML(http.StatusOK, "offer.html", data)
}

// ShowSelectionCarAd рендерит страницу конкретного автомобиля
// Входные параметры: sessionID - идентификатор сессии, car - автомобиль
func (slp *selectionPresenter) ShowSelectionCarAd(ctx adapters.Context, sessionID string, car *entities.Car) {
	link := fmt.Sprintf("http://localhost:8082/selection/guest/%s", sessionID)
	ctx.HTML(http.StatusOK, "car_card.html", gin.H{"Car": car, "Link": link})
}
