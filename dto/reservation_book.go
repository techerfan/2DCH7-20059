package dto

import "time"

type (
	ReservationBookRequest struct {
		NumberOfSeats uint8     `json:"number_of_seats"`
		UserID        uint      `json:"user_id"`
		StartDT       time.Time `json:"start_dt"`
		EndDT         time.Time `json:"end_dt"`
	}

	ReservationBookResponse struct {
		ReservationID uint   `json:"reservation_id"`
		TableNumber   uint8  `json:"table_number"`
		Seats         uint8  `json:"seats"`
		Price         uint64 `json:"price"`
	}
)
