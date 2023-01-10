package usecase

import (
	"context"

	"github.com/bmstu-rsoi/rsoi-2022-lab2-microservices-AnastasiiaRumak/src/cars_service/internal/models"

)

type repo interface {

	GetCarByID(ctx context.Context, id int64) (models.Car, error)
	GetAll(ctx context.Context) (*[]models.Car, error)

}