package fiber

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/swagger"
	"github.com/techerfan/2DCH7-20059/contract"
	"github.com/techerfan/2DCH7-20059/delivery/httpserver"
	"github.com/techerfan/2DCH7-20059/delivery/httpserver/fiber/handlers"
	"github.com/techerfan/2DCH7-20059/pkg/logger"
	"github.com/techerfan/2DCH7-20059/pkg/myjwt"
)

type http struct {
	*handlers.Handler
	app                 *fiber.App
	tokenGenerator      myjwt.Myjwt
	tokenExpirationTime int64
	userService         contract.UserService
	tableService        contract.TableService
	reservationService  contract.ReservationServcie
	logger              logger.Logger
}

// @title 				Pars Tasmim
// @version 			1.0
// @description 	Code Challenge | 2DCH7-20059

// @contact.name 	Erfan Derakhshani
// @contact.url 	https://something.com
// @contact.email techerfan@gmail.com
// @BasePath 			/api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(
	tokenGenerator myjwt.Myjwt,
	logger logger.Logger,
	tokenExpirationTime int64,
	userService contract.UserService,
	tableService contract.TableService,
	reservationService contract.ReservationServcie,
) httpserver.HttpPort {
	return &http{
		Handler:             handlers.NewHandler(logger),
		tokenGenerator:      tokenGenerator,
		tokenExpirationTime: tokenExpirationTime,
		userService:         userService,
		tableService:        tableService,
		reservationService:  reservationService,
		logger:              logger,
	}
}

func (h *http) Start(port string) error {
	h.app = fiber.New(fiber.Config{
		ServerHeader: "Pars Tasmim",
	})

	h.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	h.app.Use(limiter.New(limiter.Config{
		Max:        1000,
		Expiration: 5 * time.Second,
	}))

	h.app.Get("/swagger/*", swagger.HandlerDefault)

	api := h.app.Group("/api")
	h.setupAPIRoutes(api)

	h.logger.Info("Api is up", "port", port)
	return h.app.Listen(":" + port)
}

func (h *http) Shutdown() error {
	return h.app.Shutdown()
}
