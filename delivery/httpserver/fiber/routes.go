package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techerfan/2DCH7-20059/delivery/httpserver/fiber/middleware"
)

func (h *http) setupAPIRoutes(api fiber.Router) {
	// Setup Login
	api.Post("/users/login", h.HandleLogin(h.userService))
	api.Post("/users/register", h.HandleRegister(h.userValidator, h.userService))

	// Routes that are defined after this middleware, are protected by JWT token.
	api.Use(middleware.Authenticate(h.userService, h.tokenGenerator, h.tokenExpirationTime))

	user := api.Group("/users")
	table := api.Group("/tables")
	reservations := api.Group("/reservations")

	h.userRoutes(user)
	h.tableRoutes(table)
	h.reservationRoutes(reservations)
}

func (h *http) userRoutes(user fiber.Router) {
	user.Get("/logout", h.HandleLogout(h.userService))
}

func (h *http) tableRoutes(table fiber.Router) {
	table.Get("/all", h.HandleGetAllTables(h.tableService))
	table.Post("/add", h.HandleAddTable(h.tableValidator, h.tableService))
	table.Delete("/remove", h.HandleRemoveTable(h.tableValidator, h.tableService))
	table.Get("/timetable", h.HandleTimetable(h.tableValidator, h.tableService))
}

func (h *http) reservationRoutes(reservation fiber.Router) {
	reservation.Post("/book", h.HandleBook(h.reservationValidator, h.reservationService))
	reservation.Patch("/cancel", h.HandleCancelation(h.reservationValidator, h.reservationService))
}
