package handler

import(
	"net/http"
	"fmt"

	"github.com/uchiman/air_pollution_project/data-collector/internal/handler"
)

func SetRoutes(app *fiber.App) {
	// Add "/api" prefix for all routes
	apiRouter := app.Group("/api")
   
	// Products routes
	apiRouter.Get("/products", handler.GetAllProducts)
	apiRouter.Get("/products/:id", api.GetProductById)
	apiRouter.Post("/products", api.CreateProduct)
	apiRouter.Put("/products/:id", api.UpdateProduct)
	apiRouter.Delete("/products/:id", api.DeleteProduct)
}

func GetAllProducts(ctx *fiber.Ctx) error {
	var products []model.Product
	database.DB.Find(&products)
   
	if len(products) == 0 {
	 return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
	  "data":  nil,
	  "error": "Products not found",
	 })
	}
   
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
	 "data":  products,
	 "error": nil,
	})
   }
   
   
   func GetProductById(ctx *fiber.Ctx) error {
	var product model.Product
	var id = ctx.Params("id")
	database.DB.First(&product, id)
   
	if product.Id == 0 {
	 return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
	  "data":  nil,
	  "error": "Product not found",
	 })
	}
   
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
	 "data":  product,
	 "error": nil,
	})
   }
   
   
   func CreateProduct(ctx *fiber.Ctx) error {
	var product model.Product
	var err = ctx.BodyParser(&product)
   
	if err != nil {
	 return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	  "data":  nil,
	  "error": "Invalid request",
	 })
	}
   
	err = database.DB.Create(&product).Error
	if err != nil {
	 return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	  "data":  nil,
	  "error": "Create operation failed",
	 })
	}
   
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
	 "data":  product,
	 "error": nil,
	})
   }
   
   
   func UpdateProduct(ctx *fiber.Ctx) error {
	var updatedProduct model.Product
	var err = ctx.BodyParser(&updatedProduct)
   
	if err != nil {
	 return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	  "data":  nil,
	  "error": "Invalid request",
	 })
	}
   
	var id = ctx.Params("id")
	var oldProduct model.Product
	database.DB.First(&oldProduct, id)
   
	if oldProduct.Id == 0 {
	 return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
	  "data":  nil,
	  "error": "Product not found",
	 })
	}
   
	// Prevent id update
	err = database.DB.Model(&oldProduct).Select("name", "price").Updates(model.Product{
	 Name:  updatedProduct.Name,
	 Price: updatedProduct.Price,
	}).Error
   
	if err != nil {
	 return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	  "data":  nil,
	  "error": "Update operation failed",
	 })
	}
   
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
	 "data":  updatedProduct,
	 "error": nil,
	})
   }
   
   
   func DeleteProduct(ctx *fiber.Ctx) error {
	var product model.Product
	var id = ctx.Params("id")
	database.DB.First(&product, id)
   
	if product.Id == 0 {
	 return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
	  "data":  nil,
	  "error": "Product not found",
	 })
	}
   
	var err = database.DB.Delete(&product, id).Error
	if err != nil {
	 return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	  "data":  nil,
	  "error": "Delete operation failed",
	 })
	}
   
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
	 "data":  product,
	 "error": nil,
	})
   }