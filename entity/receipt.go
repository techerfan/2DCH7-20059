package entity

type Receipt struct {
	ID            uint   `json:"id"`
	ReservationID uint   `json:"reservation_id"`
	Price         uint64 `json:"price"`
	// Reservation is only valid when the client has payed the receipt
	IsPayed bool `json:"is_payed"`
	// If the reservation is canceled, the receipt must be payed back to the client
	IsPayedBack bool `json:"is_payed_back"`
}
