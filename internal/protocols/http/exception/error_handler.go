package exception

import (
	"errors"
	"net/http"

	"toko-bangunan/internal/protocols/http/exception/message"
	"toko-bangunan/internal/protocols/http/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	validationException, okValidasi := err.(validator.ValidationErrors)
	badRequestException, okBadRequest := err.(BadRequestError)
	notFoundErrorException, okNotFound := err.(NotFoundError)
	UnauthorizeException, okUnauthorize := err.(Unauthorize)
	UnprocessableEntityException, okUnprocessableEntity := err.(UnprocessableEntity)
	forbiddenException, okForbidden := err.(Forbidden)

	if okValidasi {
		return ctx.Status(http.StatusBadRequest).JSON(response.NewResponse("BAD REQUEST", http.StatusBadRequest, message.ValidationMessage(validationException), nil, nil))
	} else if okBadRequest {
		return ctx.Status(http.StatusBadRequest).JSON(response.NewResponse("BAD REQUEST", http.StatusBadRequest, badRequestException.Message, nil, nil))
	} else if okNotFound {
		return ctx.Status(http.StatusNotFound).JSON(response.NewResponse("NOT FOUND", http.StatusNotFound, notFoundErrorException.Message, nil, nil))
	} else if okUnauthorize {
		return ctx.Status(http.StatusUnauthorized).JSON(response.NewResponse("UNAUTHORIZED", http.StatusUnauthorized, UnauthorizeException.Message, nil, nil))
	} else if okUnprocessableEntity {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(response.NewResponse("UNPROCESSABLE ENTITY", http.StatusUnprocessableEntity, UnprocessableEntityException.Message, nil, nil))
	} else if okForbidden {
		return ctx.Status(http.StatusForbidden).JSON(response.NewResponse("FORBIDDEN", http.StatusForbidden, forbiddenException.Message, nil, nil))
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
			status = "CLIENT ERROR"
		}
		ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		return ctx.Status(code).JSON(response.NewResponse(status, code, err.Error(), nil, nil))
	}
}
