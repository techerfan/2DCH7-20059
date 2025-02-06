package dto

type (
	TableAddRequest struct {
		TableNumber uint8 `json:"table_number"`
		Capacity    uint8 `json:"capacity"`
	}

	TableAddResponse struct {
		ID uint `json:"id"`
	}
)
