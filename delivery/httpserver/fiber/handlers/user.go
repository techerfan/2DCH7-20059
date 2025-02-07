package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techerfan/2DCH7-20059/contract"
	"github.com/techerfan/2DCH7-20059/dto"
	"github.com/techerfan/2DCH7-20059/validator/uservalidator"
)

// @Summary 			Register a new user
// @Description 	Register a new user
// @Tags 					user
// @Accept       	json
// @Produce      	json
// @Param 				"payload"													body 				dto.UserRegisterRequest 		true 	"payload"
// @Success 			200																{object}		dto.UserRegisterResponse
// @Failure				400																"bad request"
// @Failure 			401																"unauthorized"
// @Failure 			406																"not acceptable"
// @Failure 			500																"internal error"
// @Security 			BearerAuth
// @Router 				/users/register										[post]
func (h *Handler) HandleRegister(validator uservalidator.Validator, accountService contract.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := dto.UserRegisterRequest{}

		if err := c.BodyParser(&payload); err != nil {
			h.logger.Errorf("could not parse login request: %v", err)
			return err
		}

		if fieldErrors, err := validator.ValidateRegisterRequest(payload); err != nil {
			c.Status(fiber.StatusNotAcceptable)
			return c.JSON(fiber.Map{
				"message": err.Error(),
				"errors":  fieldErrors,
			})
		}

		resp, err := accountService.Register(c.Context(), payload)
		if err != nil {
			h.logger.Errorf("could not register: %v", err)
			return c.SendStatus(fiber.StatusNotAcceptable)
		}

		return c.JSON(resp)
	}
}

// @Summary 			login
// @Description 	login
// @Tags 					user
// @Accept       	json
// @Produce      	json
// @Param 				"payload"													body 				dto.UserLoginRequest 		true 	"payload"
// @Success 			200																{object}		dto.UserLoginResponse
// @Failure				400																"bad request"
// @Failure 			401																"unauthorized"
// @Failure 			406																"not acceptable"
// @Failure 			500																"internal error"
// @Security 			BearerAuth
// @Router 				/users/login											[post]
func (h *Handler) HandleLogin(accountService contract.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := dto.UserLoginRequest{}

		if err := c.BodyParser(&payload); err != nil {
			h.logger.Errorf("could not parse login request: %v", err)
			return err
		}

		resp, err := accountService.Login(c.Context(), payload)
		if err != nil {
			h.logger.Errorf("could not login: %v", err)
			return c.SendStatus(fiber.StatusNotAcceptable)
		}

		c.Status(fiber.StatusOK)
		return c.JSON(resp)
	}
}

// @Summary 			logout
// @Description 	logout
// @Tags 					user
// @Accept       	json
// @Produce      	json
// @Success 			200																{object}		dto.UserLogoutResponse
// @Failure				400																"bad request"
// @Failure 			401																"unauthorized"
// @Failure 			406																"not acceptable"
// @Failure 			500																"internal error"
// @Security 			BearerAuth
// @Router 				/users/logout											[get]
func (h *Handler) HandleLogout(accountService contract.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := dto.UserLogoutRequest{}

		// Extracting user id for validation purposes
		userID := uint(c.UserContext().Value(contract.UserID).(float64))

		payload.UserID = userID

		resp, err := accountService.Logout(c.Context(), payload)
		if err != nil {
			h.logger.Errorf("could not logout: %v", err)
			return c.SendStatus(fiber.StatusNotAcceptable)
		}

		return c.JSON(resp)
	}
}
