package http

import (
	"toko-bangunan/infrastructures/db"
	fileupload "toko-bangunan/internal/utils/file-upload"
	categoryproductrepo "toko-bangunan/src/modules/category-product/repositories"
	categoryproductservice "toko-bangunan/src/modules/category-product/services"
	productrepo "toko-bangunan/src/modules/product/repositories"
	productservice "toko-bangunan/src/modules/product/services"
	supplierrepo "toko-bangunan/src/modules/supplier/repositories"
	supplierservice "toko-bangunan/src/modules/supplier/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type HttpHandlerImpl struct {
	CategoryProductService categoryproductservice.CategoryProductService
	SupplierService        supplierservice.SupplierService
	ProductService         productservice.ProductService
}

func NewHttpHandlerImpl(
	categoryProductService categoryproductservice.CategoryProductService,
	supplierService supplierservice.SupplierService,
	productService productservice.ProductService,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		CategoryProductService: categoryProductService,
		SupplierService:        supplierService,
		ProductService:         productService,
	}
}

var (
	categoryProductRepo    = categoryproductrepo.NewRepositoriesImpl()
	categoryProductService = categoryproductservice.NewCategoryProductService(categoryProductRepo, db.NewMysqlConnection(), validator.New())

	supplierRepo    = supplierrepo.NewRepositoriesImpl()
	supplierService = supplierservice.NewSupplierServiceImpl(*supplierRepo, db.NewMysqlConnection(), validator.New())

	productRepo    = productrepo.NewRepositoriesImpl()
	productService = productservice.NewProductServiceImpl(*supplierRepo, *categoryProductRepo, *productRepo, db.NewMysqlConnection(), validator.New())

	handler = NewHttpHandlerImpl(categoryProductService, supplierService, productService)
)

func (h *HttpHandlerImpl) Router(router *fiber.App) {
	api := router.Group("/api")

	// router.Static("/public", "./public")

	api.Route("/category-product", func(r fiber.Router) {
		r.Get("/", handler.FindAllCategoryProduct)
		r.Get("/:id/detail", handler.FindByIdCategoryProduct)
		r.Post("/create", handler.CreateCategoryProduct)
		r.Put("/:id/update", handler.UpdateCategoryProduct)
		r.Delete("/:id/delete", handler.DeleteCategoryProduct)
	})

	// suplier
	api.Route("/supplier", func(router fiber.Router) {
		router.Get("/", handler.FindAllSupplier)
		router.Get("/:id/detail", handler.FindByIdSupplier)
		router.Post("/create", handler.CreateSupplier)
		router.Put("/:id/update", handler.UpdateSupplier)
		router.Delete("/:id/delete", handler.DeleteSupplier)
	})

	// product
	api.Route("/product", func(router fiber.Router) {
		router.Get("/", handler.FindAllProduct)
		router.Get("/:id/detail", handler.FindByIdProduct)
		router.Post("/create", fileupload.HandleSingleFileImg, handler.CreateProduct)
		router.Put("/:id/update", fileupload.HandleSingleFileImg, handler.UpdateProduct)
		router.Delete("/:id/delete", handler.DeleteProduct)
	})
}
