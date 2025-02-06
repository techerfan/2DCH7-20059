package dto

import "time"

type (
	ReservationGetAllRequest struct {
		DT time.Time `json:"dt"`
	}

	ReservationGetAllResponse struct {
		Reservations []Reservation `json:"reservations"`
	}
)
