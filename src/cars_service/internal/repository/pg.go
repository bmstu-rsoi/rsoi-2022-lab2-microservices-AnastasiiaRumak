package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"database/sql"
	"errors"
	"github.com/bmstu-rsoi/rsoi-2022-lab2-microservices-AnastasiiaRumak/src/cars_service/internal/models"

)
var (
	ErrNoPersonWithSuchID = errors.New("no car")
)

const (

	selectByIDQuery = `SELECT * FROM car WHERE id = $1`
	selectAll = `SELECT * FROM car`
)

type PG struct {
	db *sqlx.DB

}

func (p *PG) GetCarByID(ctx context.Context, id int64) (models.Car, error) {
	row := p.db.QueryRowxContext(ctx, selectByIDQuery, id)

	var car BDlist
	err := row.StructScan(&car)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Car{}, errors.New("no car")
		}
		return models.Car{}, err
	}

	return toModel(car), nil
}

func (p *PG) GetAll(ctx context.Context) (*[]models.Car, error) {
	row, err := p.db.QueryxContext(ctx, selectAll)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var cars []models.Car
	for row.Next() {
		var car BDlist
		err = row.StructScan(&car)
		if err != nil {
			return nil, err
		}
		cars = append(cars, toModel(car))
	}
	if err = row.Err(); err != nil {
		return nil, err
	}
	return &cars, nil
}

type BDlist struct {
	ID int `db:"id"`
	CarUID string `db:"carUID"`
	Brand string `db:"brand"`
	Model string `db:"model"`
	Registration_number string `db:"number"`
	Power int `db:"power"`
	Price int `db:"price"`
	Type string `db:"type"`
	Availability bool `db:"availability"`
}

func toModel(b BDlist) models.Car {
	return models.Car{
		ID: b.ID,
		CarUID: b.CarUID,
		Brand: b.Brand,
		Model: b.Model,
		Registration_number: b.Registration_number,
		Power: b.Power,
		Price: b.Price,
		Type: b.Type,
		Availability: b.Availability,
	}
}