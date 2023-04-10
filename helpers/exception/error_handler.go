package exception

import (
	"net/http"

	"github.com/SyaibanAhmadRamadhan/toko-bangunan/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	validationException, okValidasi := err.(validator.ValidationErrors)
	badRequestException, okBadRequest := err.(BadRequestError)
	notFoundErrorException, okNotFound := err.(NotFoundError)

	if okValidasi {
		return ctx.JSON(helpers.NewResponse("BAD REQUEST", http.StatusBadRequest, helpers.ValidationMessage(validationException), nil))
	} else if okBadRequest {
		return ctx.JSON(helpers.NewResponse("BAD REQUEST", http.StatusBadRequest, badRequestException.Message, nil))
	} else if okNotFound {
		return ctx.JSON(helpers.NewResponse("NOT FOUND", http.StatusNotFound, notFoundErrorException.Message, nil))
	} else {
		return ctx.JSON(helpers.NewResponse("INTERNAL SERVER ERROR", http.StatusInternalServerError, err, nil))
	}
}
