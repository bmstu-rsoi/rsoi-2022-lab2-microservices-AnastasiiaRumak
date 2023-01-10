package usecase

import (
	"context"

	"github.com/bmstu-rsoi/rsoi-2022-lab2-microservices-AnastasiiaRumak/src/cars_service/internal/models"

)

type UseCase struct {
	repo repo
}


func (u *UseCase) GetCarByID(ctx context.Context, id int64) (models.Car, error) {
	return u.repo.GetCarByID(ctx, id)
}
func (u *UseCase) GetAll(ctx context.Context) (*[]models.Car, error) {
	return u.repo.GetAll(ctx)
}
