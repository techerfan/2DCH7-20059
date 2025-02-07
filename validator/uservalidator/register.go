package uservalidator

import (
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/techerfan/2DCH7-20059/dto"
)

func (v Validator) ValidateRegisterRequest(req dto.UserRegisterRequest) (map[string]string, error) {

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.FirstName,
			validation.Required,
			validation.Length(3, 50)),

		validation.Field(&req.LastName,
			validation.Required,
			validation.Length(3, 50)),

		validation.Field(&req.Email,
			validation.Required,
			is.Email,
			validation.By(v.checkEmailUniqueness)),

		validation.Field(&req.Password,
			validation.Required,
			validation.Match(regexp.MustCompile(`^[A-Za-z0-9!@#%^&*]{8,}$`))),

		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error("phone number is not valid")),
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

func (v Validator) checkEmailUniqueness(value interface{}) error {
	email := value.(string)

	if isUnique, err := v.repo.IsEmailUnique(email); err != nil || !isUnique {
		if err != nil {
			return err
		}

		if !isUnique {
			return fmt.Errorf("email is not unique")
		}
	}

	return nil
}
