package delivery

import (
	"net/http"
	"errors"
	//"fmt"
	"context"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/bmstu-rsoi/rsoi-2022-lab2-microservices-AnastasiiaRumak/src/cars_service/internal/models"

)
 
const (
	locationValueFormat = "/api/v3/cars/%s"
)


type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) Configure(e *echo.Echo) {
	
	e.GET("/api/v3/cars/:id", h.GetCarByID())
	e.GET("/api/v3/cars", h.GetAll())
	
}

type response struct {
	ID int `json:"id"`
	CarUID string `json:"carUID"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Registration_number string `json:"number"`
	Power int `json:"power"`
	Price int `json:"price"`
	Type string `json:"type"`
	Availability bool `json:"availability"`
}
type httpError struct {
	Message string `json:"message"`
}


func fromModel(m models.Car) response {
	return response{	
		ID: m.ID,
		CarUID: m.CarUID,
		Brand: m.Brand,
		Model: m.Model,
		Registration_number: m.Registration_number,
		Power: m.Power,
		Price: m.Price,
		Type: m.Type,
		Availability: m.Availability,
	}
}


func (h *Handler) GetCarByID() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, nil)
		}
		cars, err := h.usecase.GetCarByID(context.Background(), id)
		if err != nil {
			if errors.Is(err, errors.New("no car with such ID")) {
				return ctx.JSON(http.StatusNotFound, httpError{Message: ""})
			}
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.Response().Header().Set("Content-Type", "application/json")
		return ctx.JSON(http.StatusOK, fromModel(cars))

	}
}

func (h *Handler) GetAll() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cars, err := h.usecase.GetAll(context.Background())
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		var respons []response
		for _, p := range *cars {
			respons = append(respons, fromModel(p))
		}

		ctx.Response().Header().Set("Content-Type", "application/json")

		return ctx.JSON(http.StatusOK, respons)

	}

}