package reservationvalidator

import (
	"context"
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/techerfan/2DCH7-20059/dto"
)

func (v Validator) ValidateCancelationRequest(req dto.ReservationCancelRequest) (map[string]string, error) {

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.ReservationID,
			validation.Required,
			validation.By(v.checkReservationExistence),
			validation.By(v.checkCancelationIsAllowed)),

		validation.Field(&req.UserID,
			validation.Required,
			validation.By(v.checkReservationBelongsToUser)),
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

func (v Validator) checkReservationExistence(value interface{}) error {
	id := value.(uint)

	if doesExist, err := v.repo.DoesReservationExist(context.Background(), id); err != nil || !doesExist {
		if err != nil {
			return err
		}

		if !doesExist {
			return fmt.Errorf("reservation does not exist")
		}
	}

	return nil
}

func (v Validator) checkCancelationIsAllowed(value interface{}) error {
	id := value.(uint)

	reservation, err := v.repo.FindReservationByID(context.Background(), id)
	if err != nil {
		return err
	}

	if reservation.StartDT.Before(time.Now().Add(time.Hour * 3)) {
		return fmt.Errorf("cancelation is only allowed when more than 3 hours is left to the reservation")
	}

	return nil
}

func (v Validator) checkReservationBelongsToUser(value interface{}) error {
	id := value.(uint)

	reservation, err := v.repo.FindReservationByID(context.Background(), id)
	if err != nil {
		return err
	}

	if reservation.UserID != id {
		return fmt.Errorf("reservation does not belong to this user")
	}

	return nil
}
