package main

import (
	"github.com/SyaibanAhmadRamadhan/toko-bangunan/app"
	"github.com/SyaibanAhmadRamadhan/toko-bangunan/controllers"
	"github.com/SyaibanAhmadRamadhan/toko-bangunan/helpers/exception"
	"github.com/SyaibanAhmadRamadhan/toko-bangunan/repositories"
	"github.com/SyaibanAhmadRamadhan/toko-bangunan/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	db := app.ConnectDb()
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})
	app.Use(recover.New())
	categoryProductRepo := repositories.NewCategoryProductRepository()
	categoryProductService := services.NewCategoryProductService(categoryProductRepo, db, validator.New())
	categoryProductController := controllers.NewCategoryProductControllerImpl(categoryProductService)

	app.Post("/api/category-product", categoryProductController.Create)
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
