package dto

import "github.com/techerfan/2DCH7-20059/entity"

type (
	UserRegisterRequest struct {
		FirstName   string        `json:"first_name"`
		LastName    string        `json:"last_name"`
		Gender      entity.Gender `json:"gender"`
		Email       string        `json:"email"`
		PhoneNumber string        `json:"phone_number"`
		Password    string        `json:"password"`
	}

	UserRegisterResponse struct {
		UserID uint `json:"user_id"`
	}
)
