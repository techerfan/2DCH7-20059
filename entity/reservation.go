package entity

import "time"

type Reservation struct {
	ID            uint          `json:"id"`
	NumberOfSeats uint8         `json:"number_of_seats"`
	UserID        uint          `json:"user_id"`
	TableID       uint          `json:"table_id"`
	ReceiptID     uint          `json:"receipt_id"`
	StartDT       time.Time     `json:"start_dt"`
	EndDT         time.Duration `json:"end_dt"`
	// The cancelation is only valid when the client has payed the receipt (Otherwise, the reservation is not submitted)
	// * Also, if the client wants to cancel the reservation, they must do it 3 hours before the start time
	IsCanceled bool `json:"is_canceled"`
}
