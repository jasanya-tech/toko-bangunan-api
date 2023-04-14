package exception

import (
	"errors"
	"net/http"

	"toko-bangunan/helpers"

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
		code := 500

		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}
		status := "error"
		if code >= 500 {
			status = "INTERNAL SERVER ERROR"
		} else if code >= 400 {
			status = "NOT FOUND"
		}
		ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		return ctx.JSON(helpers.NewResponse(status, code, err.Error(), nil))
	}
}
