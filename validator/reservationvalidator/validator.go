package reservationvalidator

import (
	"context"

	"github.com/techerfan/2DCH7-20059/entity"
)

type Repository interface {
	DoesReservationExist(ctx context.Context, id uint) (bool, error)
	FindReservationByID(ctx context.Context, id uint) (entity.Reservation, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}
