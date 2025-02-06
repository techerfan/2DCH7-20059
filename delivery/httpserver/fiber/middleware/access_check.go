package middleware

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/techerfan/2DCH7-20059/contract"
	"github.com/techerfan/2DCH7-20059/pkg/myjwt"
)

func Authenticate(userService contract.UserService, tokenGenerator myjwt.Myjwt, tokenExpirationTime int64) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := myjwt.FetchToken(c.Get("Authorization"))
		if token == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		// first fetch token claims and see if they are ok
		ok, claims := tokenGenerator.IsValid(token)
		if !ok {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		// then check expiration time to check if it's not passed
		exp, ok := (claims[1]).(float64)
		if !ok {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// extract the user id
		userID, ok := (claims[0]).(float64)
		if !ok {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		ctx := context.WithValue(c.Context(), contract.UserID, userID)
		c.SetUserContext(ctx)

		IsTokenValid := userService.IsTokenValid(ctx, token)

		if int64(exp) < time.Now().UnixMilli() && tokenExpirationTime > 0 || !IsTokenValid {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.Next()
	}
}
