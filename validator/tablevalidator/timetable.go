package tablevalidator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/techerfan/2DCH7-20059/dto"
)

func (v Validator) ValidateTimetableRequest(req dto.TableTimetableRequest) (map[string]string, error) {

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.DT,
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
