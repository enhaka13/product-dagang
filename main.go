package main

import (
	"log"
	"product-dagang/handler"
	"product-dagang/products"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/gin-contrib/cors"
)

func main() {
	dsn := "freedb_enhakanvas:3&cTwcnbWKnwbak@tcp(sql.freedb.tech:3306)/freedb_productdagangan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	productRepository := products.NewRepository(db)
	productService := products.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	//CRUD endpoint
	api.GET("/products", productHandler.GetAllProducts)
	api.GET("/products/:id", productHandler.GetProductByID)
	api.POST("/products", productHandler.CreateProduct)
	api.PUT("/products/:id", productHandler.UpdateProduct)
	api.DELETE("/products/:id", productHandler.DeleteProduct)

	//Pagination endpoint
	api.GET("/products/pages/:page", productHandler.ProductPagination)

	//Filtering endpoint
	api.GET("/products/brands/:brand", productHandler.FilterByBrand)
	api.GET("/products/types/:type", productHandler.FilterByType)
	api.GET("/products/colors/:color", productHandler.FilterByColor)
	api.GET("/products/sizes/:size", productHandler.FilterBySize)
	api.GET("/products/status/:status", productHandler.FilterByAvailability)

	router.Run()
}