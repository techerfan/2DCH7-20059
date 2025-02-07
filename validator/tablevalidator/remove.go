package tablevalidator

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/techerfan/2DCH7-20059/dto"
)

func (v Validator) ValidateRemoveTableRequest(req dto.TableRemoveRequest) (map[string]string, error) {

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.TableID,
			validation.Required,
			validation.By(v.checkTableExistenceForDeletion)),
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

func (v Validator) checkTableExistenceForDeletion(value interface{}) error {
	id := value.(uint)

	if doesExist, err := v.repo.DoesTableExist(id); err != nil || !doesExist {
		if err != nil {
			return err
		}

		if !doesExist {
			return fmt.Errorf("table does not exist")
		}
	}

	return nil
}
