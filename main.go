package main

import (
	customLog "toko-bangunan/internal/logger"
	"toko-bangunan/internal/protocols/http"
)

func main() {
	customLog.InitLogger()
	var http http.HttpImpl
	http.Listen()
}

// func main() {
// 	customLog.InitLogger()

// 	// var (
// 	// categoryProductRepo       = repositories.NewRepositories()
// 	// categoryProductService    = services.NewCategoryProductService(categoryProductRepo, db.NewMysqlConnection(), validator.New())
// 	// categoryProductController = http.NewHttpHandlerImpl(categoryProductService)

// 	// supplierRepo       = repositories.NewSupplierRepositoryImpl()
// 	// suppluerService    = services.NewSupplierServiceImpl(supplierRepo, db, validator.New())
// 	// supplierController = controllers.NewSupplierControllerImpl(suppluerService)

// 	// productRepo       = repositories.NewProductRepositoryImpl()
// 	// productService    = services.NewProductServiceImpl(supplierRepo, categoryProductRepo, productRepo, db, validator.New())
// 	// productController = controllers.NewProductControllerImpl(productService)

// 	// userRepo       = repositories.NewUserRepositoryImpl()
// 	// userService    = services.NewUserServiceImpl(userRepo, db, validator.New())
// 	// userController = controllers.NewUserControllerImpl(userService)
// 	// auth           = middleware.NewAuthMiddlewareImpl(userService)
// 	// )

// 	var http http.HttpImpl
// 	http.Listen()

// 	// auth
// 	// api.Route("/auth", func(router fiber.Router) {
// 	// 	router.Post("/create", utils.HandleSingleFileImg, userController.Create)
// 	// 	router.Post("/login", userController.AuthLogin)
// 	// 	router.Post("/refresh-token", userController.RefreshToken)
// 	// 	router.Get("/logout", auth.DeserializeUser, userController.LogoutUser)
// 	// })
// 	// static file
// 	// app.Static("/public", "./public")

// 	// // category product
// 	// api.Route("/category-product", func(router fiber.Router) {
// 	// 	router.Get("/", categoryProductController.FindAllCategoryProduct)
// 	// 	router.Get("/:id/detail", categoryProductController.FindByIdCategoryProduct)
// 	// 	router.Post("/create", categoryProductController.CreateCategoryProduct)
// 	// 	router.Put("/:id/update", categoryProductController.UpdateCategoryProduct)
// 	// 	router.Delete("/:id/delete", categoryProductController.DeleteCategoryProduct)
// 	// })

// 	// // suplier
// 	// api.Route("/supplier", func(router fiber.Router) {
// 	// 	app.Get("/", auth.DeserializeUser, supplierController.FindALL)
// 	// 	app.Get("/:id/detail", auth.DeserializeUser, supplierController.FindById)
// 	// 	app.Post("/create", auth.DeserializeUser, supplierController.Create)
// 	// 	app.Put("/:id/update", auth.DeserializeUser, supplierController.Update)
// 	// 	app.Delete("/:id/delete", auth.DeserializeUser, supplierController.Delete)
// 	// })

// 	// // product
// 	// api.Route("/product", func(router fiber.Router) {
// 	// 	app.Get("/", productController.FindALL)
// 	// 	app.Get("/:id/detail", productController.FindById)
// 	// 	app.Post("/create", auth.DeserializeUser, utils.HandleSingleFileImg, productController.Create)
// 	// 	app.Put("/:id/update", auth.DeserializeUser, utils.HandleSingleFileImg, productController.Update)
// 	// 	app.Delete("/:id/delete", auth.DeserializeUser, productController.Delete)
// 	// })
// }
