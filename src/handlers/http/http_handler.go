package http

import (
	"toko-bangunan/infrastructures/db"
	"toko-bangunan/internal/protocols/http/middleware"
	"toko-bangunan/internal/utils/auth"
	categoryproductrepo "toko-bangunan/src/modules/category-product/repositories"
	categoryproductservice "toko-bangunan/src/modules/category-product/services"
	productrepo "toko-bangunan/src/modules/product/repositories"
	productservice "toko-bangunan/src/modules/product/services"
	purchaseproductrepo "toko-bangunan/src/modules/purchase-product/repositories"
	purchaseproductservice "toko-bangunan/src/modules/purchase-product/services"
	supplierrepo "toko-bangunan/src/modules/supplier/repositories"
	supplierservice "toko-bangunan/src/modules/supplier/services"
	userrepo "toko-bangunan/src/modules/user/repositories"
	userservice "toko-bangunan/src/modules/user/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type HttpHandlerImpl struct {
	CategoryProductService categoryproductservice.CategoryProductService
	SupplierService        supplierservice.SupplierService
	ProductService         productservice.ProductService
	PurchaseProductService purchaseproductservice.PurchaseProductService
	UserService            userservice.UserService
}

func NewHttpHandlerImpl(
	categoryProductService categoryproductservice.CategoryProductService,
	supplierService supplierservice.SupplierService,
	productService productservice.ProductService,
	purchaseProductService purchaseproductservice.PurchaseProductService,
	userService userservice.UserService,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		CategoryProductService: categoryProductService,
		SupplierService:        supplierService,
		ProductService:         productService,
		PurchaseProductService: purchaseProductService,
		UserService:            userService,
	}
}

func (h *HttpHandlerImpl) Router(router *fiber.App) {
	var (
		dbConnect              = db.NewMysqlConnection()
		categoryProductRepo    = categoryproductrepo.NewRepositoriesImpl()
		categoryProductService = categoryproductservice.NewCategoryProductService(categoryProductRepo, dbConnect, validator.New())

		supplierRepo    = supplierrepo.NewRepositoriesImpl()
		supplierService = supplierservice.NewSupplierServiceImpl(*supplierRepo, dbConnect, validator.New())

		productRepo    = productrepo.NewRepositoriesImpl()
		productService = productservice.NewProductServiceImpl(*supplierRepo, *categoryProductRepo, *productRepo, dbConnect, validator.New())

		purchaseProductRepo    = purchaseproductrepo.NewRepositoriesImpl()
		purchaseProductService = purchaseproductservice.NewPurchaseProductServiceImpl(*purchaseProductRepo, *productRepo, dbConnect, validator.New())

		userRepo      = userrepo.NewUserRepositoryImpl()
		jwtAuth       = auth.NewJwtTokenImpl(userRepo, dbConnect)
		userService   = userservice.NewUserServiceImpl(userRepo, dbConnect, validator.New(), jwtAuth)
		middlewareJwt = middleware.NewAuthMiddlewareImpl(userService)
		handler       = NewHttpHandlerImpl(categoryProductService, supplierService, productService, purchaseProductService, userService)
	)

	api := router.Group("/api")

	router.Static("/public", "./public")

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
		router.Post("/create", middleware.HandleSingleFileImg, handler.CreateProduct)
		router.Put("/:id/update", middleware.HandleSingleFileImg, handler.UpdateProduct)
		router.Delete("/:id/delete", handler.DeleteProduct)
	})

	// product
	api.Route("/purchase-product", func(router fiber.Router) {
		router.Get("/", handler.FindAllPurchaseProduct)
		router.Get("/:id/detail", handler.FindByIdPurchaseProduct)
		router.Post("/create", handler.CreatePurchaseProduct)
		router.Put("/:id/update", handler.UpdatePurchaseProduct)
		router.Delete("/:id/delete", handler.DeletePurchaseProduct)
	})

	// user
	api.Route("/auth", func(router fiber.Router) {
		router.Post("/create", middleware.HandleSingleFileImg, handler.CreateUser)
		router.Get("/:id/detail", middlewareJwt.JwtVerifyToken, handler.FindByIdUser)
		router.Post("/login", handler.UserLogin)
		router.Post("/refresh-token", middlewareJwt.JwtVerifyRefreshToken, handler.UserRefreshToken)
		router.Get("/logout", middlewareJwt.JwtVerifyToken, handler.UserLogout)
	})
}
