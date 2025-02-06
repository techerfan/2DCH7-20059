package dto

type (
	ReservationCancelRequest struct {
		ReservationID uint `json:"reservation_id"`
	}

	ReservationCancelResponse struct{}
)
