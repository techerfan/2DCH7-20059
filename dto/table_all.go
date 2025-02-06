package dto

type (
	Table struct {
		ID          uint  `json:"id"`
		TableNumber uint8 `json:"table_number"`
		Capacity    uint8 `json:"capacity"`
	}

	TableAllRequest struct{}

	TableAllResponse struct {
		Tables []Table `json:"tables"`
	}
)
