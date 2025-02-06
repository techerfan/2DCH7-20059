package contract

import (
	"context"

	"github.com/techerfan/2DCH7-20059/dto"
)

type ReservationServcie interface {
	AllReservations(context.Context, dto.ReservationGetAllRequest) (dto.ReservationGetAllResponse, error)
	Book(context.Context, dto.ReservationBookRequest) (dto.ReservationBookResponse, error)
	Cancel(context.Context, dto.ReservationCancelRequest) (dto.ReservationCancelResponse, error)
}
