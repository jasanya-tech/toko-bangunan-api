package controllers

import (
	"fmt"

	"toko-bangunan/helpers"
	"toko-bangunan/models/dto"
	"toko-bangunan/services"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	FindALL(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	AuthLogin(c *fiber.Ctx) error
}

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserControllerImpl(userService services.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) FindALL(c *fiber.Ctx) error {
	// suppliers := controller.SupplierService.FindALL(c.Context())
	// return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "list data supplier", suppliers))
	return nil
}

func (controller *UserControllerImpl) FindById(c *fiber.Ctx) error {
	// supplier := controller.SupplierService.FindById(c.Context(), c.Params("id"))
	// return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "data supplier", supplier))
	return nil
}

func (controller *UserControllerImpl) Create(c *fiber.Ctx) error {
	user := new(dto.CreateUserReq)

	if err := c.BodyParser(user); err != nil {
		panic(err)
	}

	fileImg := c.Locals("filename")
	if fileImg == "" {
		user.Image = "default.png"
	} else {
		user.Image = fileImg.(string)
	}

	userService := controller.UserService.Create(c.Context(), *user)

	file, _ := c.FormFile("image")
	if file != nil {
		if err := c.SaveFile(file, fmt.Sprintf("%s/%s", "./public/images/user-img", fileImg)); err != nil {
			panic(err)
		}
	}

	return c.Status(fiber.StatusCreated).JSON(helpers.NewResponse("ok", fiber.StatusCreated, "created user successfully", userService))
}

func (controller *UserControllerImpl) Update(c *fiber.Ctx) error {
	// supplier := new(dto.UpdateSupplierReq)
	// if err := c.BodyParser(supplier); err != nil {
	// 	panic(err)
	// }
	// supplierSercice := controller.SupplierService.Update(c.Context(), *supplier, c.Params("id"))
	// return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "update supplier successfully", supplierSercice))
	return nil
}

func (controller *UserControllerImpl) Delete(c *fiber.Ctx) error {
	// controller.SupplierService.Delete(c.Context(), c.Params("id"))
	// return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "delete supplier successfully", nil))
	return nil
}

func (controller *UserControllerImpl) AuthLogin(c *fiber.Ctx) error {
	payload := new(dto.LoginReq)

	if err := c.BodyParser(payload); err != nil {
		panic(err)
	}

	response := controller.UserService.AuthLogin(c.Context(), *payload)
	return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "login successfully", response))
}
