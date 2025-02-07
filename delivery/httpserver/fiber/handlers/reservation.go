package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/techerfan/2DCH7-20059/contract"
	"github.com/techerfan/2DCH7-20059/dto"
	"github.com/techerfan/2DCH7-20059/service/reservationservice"
)

// @Summary 			book a table
// @Description 	book a table
// @Tags 					reservation
// @Accept       	json
// @Produce      	json
// @Param 				"payload"													body 				dto.ReservationBookRequest 		true 	"payload"
// @Success 			201																{object}		dto.ReservationBookResponse
// @Failure				400																"bad request"
// @Failure 			401																"unauthorized"
// @Failure 			406																"not acceptable"
// @Failure 			500																"internal error"
// @Security 			BearerAuth
// @Router 				/reservations/book			[post]
func (h *Handler) HandleBook(reservationService contract.ReservationServcie) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := dto.ReservationBookRequest{}

		if err := c.BodyParser(&payload); err != nil {
			h.logger.Errorf("could not parse the payload: %v", err)
			c.Status(fiber.StatusBadRequest)
			return err
		}

		// TODO: validate payload

		resp, err := reservationService.Book(c.Context(), payload)
		if err != nil {
			h.logger.Errorf("could not reserve: %v", err)
			if errors.Is(err, reservationservice.ErrNoAvailableTable) {
				c.Status(fiber.StatusNotAcceptable)
				return err
			}
			c.Status(fiber.StatusInternalServerError)
			return err
		}

		c.Status(fiber.StatusCreated)
		return c.JSON(resp)
	}
}

// @Summary 			cancel a reservation
// @Description 	cancel a reservation
// @Tags 					reservation
// @Accept       	json
// @Produce      	json
// @Param 				"payload"													body 				dto.ReservationCancelRequest 		true 	"payload"
// @Success 			200																{object}		dto.ReservationCancelResponse
// @Failure				400																"bad request"
// @Failure 			401																"unauthorized"
// @Failure 			406																"not acceptable"
// @Failure 			500																"internal error"
// @Security 			BearerAuth
// @Router 				/reservations/cancel							[patch]
func (h *Handler) HandleCancelation(reservationService contract.ReservationServcie) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := dto.ReservationCancelRequest{}

		if err := c.BodyParser(&payload); err != nil {
			h.logger.Errorf("could not parse the payload: %v", err)
			c.Status(fiber.StatusBadRequest)
			return err
		}

		// TODO: validate payload

		resp, err := reservationService.Cancel(c.Context(), payload)
		if err != nil {
			h.logger.Errorf("could not cancel: %v", err)
			c.Status(fiber.StatusInternalServerError)
			return err
		}

		return c.JSON(resp)
	}
}
