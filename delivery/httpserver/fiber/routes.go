package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techerfan/2DCH7-20059/delivery/httpserver/fiber/middleware"
)

func (h *http) setupAPIRoutes(api fiber.Router) {
	// Setup Login
	api.Post("/users/login", h.HandleLogin(h.userService))
	api.Post("/users/register", h.HandleRegister(h.userService))

	// Other routes except for login must be protected by the token
	api.Use(middleware.Authenticate(h.userService, h.tokenGenerator, h.tokenExpirationTime))

	user := api.Group("/users")

	h.userRoutes(user)

}

func (h *http) userRoutes(user fiber.Router) {
	user.Get("/logout", h.HandleLogout(h.userService))
}
