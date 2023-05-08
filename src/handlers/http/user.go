package http

import (
	"toko-bangunan/internal/protocols/http/exception"
	"toko-bangunan/internal/protocols/http/response"
	"toko-bangunan/src/modules/user/dto"

	"github.com/gofiber/fiber/v2"
)

func (http *HttpHandlerImpl) CreateUser(c *fiber.Ctx) error {
	user := new(dto.CreateUserReq)

	if err := c.BodyParser(user); err != nil {
		panic(err)
	}

	fileImg := c.Locals("filename")
	if fileImg == "" {
		user.Image = "default.jpg"
	} else {
		user.Image = fileImg.(string)
	}

	userService, err := http.UserService.Create(c.Context(), *user)
	if err != nil {
		panic(err)
	}

	return c.Status(fiber.StatusCreated).JSON(response.NewResponse("ok", 201, "created user successfully", userService, nil))
}

func (http *HttpHandlerImpl) FindByIdUser(c *fiber.Ctx) error {
	user, err := http.UserService.FindById(c.Context(), c.Params("id"))
	if err != nil {
		panic(err)
	}
	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", 200, "data user", user, nil))
}

func (http *HttpHandlerImpl) UserLogin(c *fiber.Ctx) error {
	user := new(dto.LoginReq)

	if err := c.BodyParser(user); err != nil {
		panic(err)
	}

	userLogin, err := http.UserService.UserLogin(c.Context(), *user)
	if err != nil {
		panic(err)
	}

	return c.Status(fiber.StatusCreated).JSON(response.NewResponse("ok", 201, "created user successfully", userLogin, nil))
}

func (http *HttpHandlerImpl) UserRefreshToken(c *fiber.Ctx) error {
	id := c.Get("id")

	tokenFind, err := http.UserService.FindTokenByUserId(c.Context(), id)
	if err != nil {
		panic(err)
	}

	if tokenFind.RefreshToken != c.Get("Authorization") {
		panic(exception.Unauthorize{Message: "token invalid"})
	}
	tokenNew, err := http.UserService.UserRefreshToken(c.Context(), id)
	if err != nil {
		panic(err)
	}

	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "create token successfully", tokenNew, nil))
}

func (http *HttpHandlerImpl) UserLogout(c *fiber.Ctx) error {
	id := c.Get("id")
	_, err := http.UserService.FindTokenByUserId(c.Context(), id)
	if err != nil {
		panic(err)
	}

	http.UserService.DeleteToken(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "logout successfully", nil, nil))
}
