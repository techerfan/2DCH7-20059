package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techerfan/2DCH7-20059/contract"
	"github.com/techerfan/2DCH7-20059/dto"
)

// @Summary 			Get all tables
// @Description 	Get all tables
// @Tags 					table
// @Accept       	json
// @Produce      	json
// @Success 			200																{object}		dto.TableAllResponse
// @Failure				400																"bad request"
// @Failure 			401																"unauthorized"
// @Failure 			406																"not acceptable"
// @Failure 			500																"internal error"
// @Security 			BearerAuth
// @Router 				/tables/all												[get]
func (h *Handler) HandleGetAllTables(tableService contract.TableService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := dto.TableAllRequest{}

		resp, err := tableService.All(c.Context(), payload)
		if err != nil {
			h.logger.Errorf("could not get all the tables: %v", err)
			c.Status(fiber.StatusInternalServerError)
			return err
		}

		return c.JSON(resp)
	}
}

// @Summary 			Add a new table
// @Description 	Add a new table
// @Tags 					table
// @Accept       	json
// @Produce      	json
// @Param 				"payload"													body 				dto.TableAddRequest 		true 	"payload"
// @Success 			201																{object}		dto.TableAddResponse
// @Failure				400																"bad request"
// @Failure 			401																"unauthorized"
// @Failure 			406																"not acceptable"
// @Failure 			500																"internal error"
// @Security 			BearerAuth
// @Router 				/tables/add												[post]
func (h *Handler) HandleAddTable(tableService contract.TableService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := dto.TableAddRequest{}

		if err := c.BodyParser(&payload); err != nil {
			h.logger.Errorf("could not parse the payload: %v", err)
			c.Status(fiber.StatusBadRequest)
			return err
		}

		// TODO: validate payload

		resp, err := tableService.AddTable(c.Context(), payload)
		if err != nil {
			h.logger.Errorf("could not add the table: %v", err)
			c.Status(fiber.StatusInternalServerError)
			return err
		}

		c.Status(fiber.StatusCreated)
		return c.JSON(resp)
	}
}

// @Summary 			Delete a table
// @Description 	Delete a table
// @Tags 					table
// @Accept       	json
// @Produce      	json
// @Param 				"payload"													body 				dto.TableRemoveRequest 		true 	"payload"
// @Success 			200																{object}		dto.TableRemoveResponse
// @Failure				400																"bad request"
// @Failure 			401																"unauthorized"
// @Failure 			406																"not acceptable"
// @Failure 			500																"internal error"
// @Security 			BearerAuth
// @Router 				/tables/remove												[delete]
func (h *Handler) HandleRemoveTable(tableService contract.TableService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := dto.TableRemoveRequest{}

		if err := c.BodyParser(&payload); err != nil {
			h.logger.Errorf("could not parse the payload: %v", err)
			c.Status(fiber.StatusBadRequest)
			return err
		}

		// TODO: validate payload

		resp, err := tableService.RemoveTable(c.Context(), payload)
		if err != nil {
			h.logger.Errorf("could not remove the table: %v", err)
			c.Status(fiber.StatusInternalServerError)
			return err
		}

		return c.JSON(resp)
	}
}

// @Summary 			Get the timetable
// @Description 	Get the timetable
// @Tags 					table
// @Accept       	json
// @Produce      	json
// @Param 				"payload"													query 				dto.TableTimetableRequest 		true 	"payload"
// @Success 			200																{object}		dto.TableTimetableResponse
// @Failure				400																"bad request"
// @Failure 			401																"unauthorized"
// @Failure 			406																"not acceptable"
// @Failure 			500																"internal error"
// @Security 			BearerAuth
// @Router 				/tables/timetable												[get]
func (h *Handler) HandleTimetable(tableService contract.TableService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := dto.TableTimetableRequest{}

		if err := c.QueryParser(&payload); err != nil {
			h.logger.Errorf("could not parse the payload: %v", err)
			c.Status(fiber.StatusBadRequest)
			return err
		}

		// TODO: validate payload

		resp, err := tableService.Timetable(c.Context(), payload)
		if err != nil {
			h.logger.Errorf("could not get timetable: %v", err)
			c.Status(fiber.StatusInternalServerError)
			return err
		}

		return c.JSON(resp)
	}
}
