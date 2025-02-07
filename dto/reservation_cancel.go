package dto

type (
	ReservationCancelRequest struct {
		ReservationID uint `json:"reservation_id"`
		UserID        uint `json:"-"`
	}

	ReservationCancelResponse struct{}
)
