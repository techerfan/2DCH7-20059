package tablevalidator

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/techerfan/2DCH7-20059/dto"
)

func (v Validator) ValidateAddTableRequest(req dto.TableAddRequest) (map[string]string, error) {

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Capacity,
			validation.Required,
			validation.Min(uint8(4)),
			validation.Max(uint8(10))),

		validation.Field(&req.TableNumber,
			validation.Required,
			validation.Min(uint8(1)),
			validation.Max(uint8(10)),
			validation.By(v.checkTableExistence)),
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

func (v Validator) checkTableExistence(value interface{}) error {
	tableNum := value.(uint8)

	if doesExist, err := v.repo.DoesTableExistByTableNum(tableNum); err != nil || doesExist {
		if err != nil {
			return err
		}

		if doesExist {
			return fmt.Errorf("table number is taken")
		}
	}

	return nil
}
