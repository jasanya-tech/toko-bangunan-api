package routes

import (
	"toko-bangunan/app"
	"toko-bangunan/controllers"
	"toko-bangunan/repositories"
	"toko-bangunan/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	db = app.ConnectDb()

	categoryProductRepo       = repositories.NewCategoryProductRepository()
	categoryProductService    = services.NewCategoryProductService(categoryProductRepo, db, validator.New())
	categoryProductController = controllers.NewCategoryProductControllerImpl(categoryProductService)

	supplierRepo       = repositories.NewSupplierRepositoryImpl()
	suppluerService    = services.NewSupplierServiceImpl(supplierRepo, db, validator.New())
	supplierController = controllers.NewSupplierControllerImpl(suppluerService)
)

func IndexRoutes(r *fiber.App) {
	// category product
	r.Get("/api/category-product", categoryProductController.FindALL)
	r.Get("/api/category-product/:id/detail", categoryProductController.FindById)
	r.Post("/api/category-product", categoryProductController.Create)
	r.Put("/api/category-product/:id/update", categoryProductController.Update)
	r.Delete("/api/category-product/:id/delete", categoryProductController.Delete)

	// suplier
	r.Get("/api/supplier", supplierController.FindALL)
	r.Get("/api/supplier/:id/detail", supplierController.FindById)
	r.Post("/api/supplier", supplierController.Create)
	r.Put("/api/supplier/:id/update", supplierController.Update)
	r.Delete("/api/supplier/:id/delete", supplierController.Delete)
}
