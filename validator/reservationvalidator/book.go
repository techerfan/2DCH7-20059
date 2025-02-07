package reservationvalidator

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/techerfan/2DCH7-20059/dto"
)

func (v Validator) ValidateBookRequest(req dto.ReservationBookRequest) (map[string]string, error) {

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.NumberOfSeats,
			validation.Required,
			validation.Min(uint8(1)),
			validation.Max(uint8(10))),

		validation.Field(&req.StartDT,
			validation.Required,
			validation.By(v.isStartDTValid(req.EndDT))),

		validation.Field(&req.EndDT,
			validation.Required),
	); err != nil {
		fieldErrors := make(map[string]string)

		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, err
	}

	return nil, nil
}

func (v Validator) isStartDTValid(endDT time.Time) validation.RuleFunc {
	return func(value interface{}) error {
		startDT := value.(time.Time)

		if startDT.After(endDT) {
			return fmt.Errorf("start datetime cannot be bigger than end datetime")
		}

		if startDT.Before(time.Now()) {
			return fmt.Errorf("you cannot reserve a table for a past time")
		}

		return nil
	}
}
