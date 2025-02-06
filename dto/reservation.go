package dto

import "time"

type Reservation struct {
	ID            uint          `json:"id"`
	NumberOfSeats uint8         `json:"number_of_seats"`
	UserID        uint          `json:"user_id"`
	TableID       uint          `json:"table_id"`
	ReceiptID     uint          `json:"receipt_id"`
	StartDT       time.Time     `json:"start_dt"`
	EndDT         time.Duration `json:"end_dt"`
	IsCanceled    bool          `json:"is_canceled"`
}
