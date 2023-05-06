package http

import (
	"fmt"

	"toko-bangunan/config"
	"toko-bangunan/internal/protocols/http/exception"
	"toko-bangunan/src/handlers/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

type HttpImpl struct {
	handlers *http.HttpHandlerImpl
}

func NewHttpProtocol(handlers *http.HttpHandlerImpl) *HttpImpl {
	return &HttpImpl{
		handlers: handlers,
	}
}

func (p HttpImpl) Listen() {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, id",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	p.handlers.Router(app)

	serverPort := fmt.Sprintf(":%s", config.Get().Application.Port)
	_ = app.Listen(serverPort)
	// if err != nil {
	// 	panic(err)
	// }
	log.Info().Msgf("Server started on Port %s ", serverPort)
}
