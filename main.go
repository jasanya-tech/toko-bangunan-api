package main

import (
	"toko-bangunan/helpers/exception"

	"toko-bangunan/app"
	"toko-bangunan/controllers"
	"toko-bangunan/repositories"
	"toko-bangunan/services"
	"toko-bangunan/utils"

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

	var (
		categoryProductRepo       = repositories.NewCategoryProductRepository()
		categoryProductService    = services.NewCategoryProductService(categoryProductRepo, db, validator.New())
		categoryProductController = controllers.NewCategoryProductControllerImpl(categoryProductService)

		supplierRepo       = repositories.NewSupplierRepositoryImpl()
		suppluerService    = services.NewSupplierServiceImpl(supplierRepo, db, validator.New())
		supplierController = controllers.NewSupplierControllerImpl(suppluerService)

		productRepo       = repositories.NewProductRepositoryImpl()
		productService    = services.NewProductServiceImpl(supplierRepo, categoryProductRepo, productRepo, db, validator.New())
		productController = controllers.NewProductControllerImpl(productService)

		userRepo       = repositories.NewUserRepositoryImpl()
		userService    = services.NewUserServiceImpl(userRepo, db, validator.New())
		userController = controllers.NewUserControllerImpl(userService)
	)
	app.Use(recover.New())

	app.Static("/public", "./public")
	// category product
	app.Get("/api/category-product", categoryProductController.FindALL)
	app.Get("/api/category-product/:id/detail", categoryProductController.FindById)
	app.Post("/api/category-product", categoryProductController.Create)
	app.Put("/api/category-product/:id/update", categoryProductController.Update)
	app.Delete("/api/category-product/:id/delete", categoryProductController.Delete)

	// suplier
	app.Get("/api/supplier", supplierController.FindALL)
	app.Get("/api/supplier/:id/detail", supplierController.FindById)
	app.Post("/api/supplier", supplierController.Create)
	app.Put("/api/supplier/:id/update", supplierController.Update)
	app.Delete("/api/supplier/:id/delete", supplierController.Delete)

	// product
	app.Get("/api/product", productController.FindALL)
	app.Get("/api/product/:id/detail", productController.FindById)
	app.Post("/api/product", utils.HandleSingleFileImg, productController.Create)
	app.Put("/api/product/:id/update", utils.HandleSingleFileImg, productController.Update)
	app.Delete("/api/product/:id/delete", productController.Delete)

	// user
	app.Get("/api/user", userController.FindALL)
	app.Get("/api/user/:id/detail", userController.FindById)
	app.Post("/api/user", utils.HandleSingleFileImg, userController.Create)
	app.Post("/api/user/login", userController.AuthLogin)
	app.Put("/api/user/:id/update", utils.HandleSingleFileImg, userController.Update)
	app.Delete("/api/user/:id/delete", userController.Delete)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
