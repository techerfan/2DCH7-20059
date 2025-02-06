package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techerfan/2DCH7-20059/contract"
	"github.com/techerfan/2DCH7-20059/dto"
)

func (h *Handler) HandleRegister(accountService contract.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := dto.UserRegisterRequest{}

		if err := c.BodyParser(&payload); err != nil {
			h.logger.Errorf("could not parse login request: %v", err)
			return err
		}

		resp, err := accountService.Register(c.Context(), payload)
		if err != nil {
			h.logger.Errorf("could not register: %v", err)
			return c.SendStatus(fiber.StatusNotAcceptable)
		}

		return c.JSON(resp)
	}
}

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

		return c.JSON(resp)
	}
}

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
