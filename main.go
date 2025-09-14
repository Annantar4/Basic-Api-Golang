package main

import (
	"annanta/backend-api/config"
	"annanta/backend-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.ConnectDatabase()
	// productStruct := &controllers.ProductModel{}

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Success",
		})
	})

	router.GET("/api/product", controllers.GetAllProduct)
	router.POST("/api/product", controllers.CreateProduct)
	// router.GET("/api/product/:id", productStruct.GetProductById)
	router.GET("/api/product/:id", controllers.GetProductById)
	router.PUT("/api/product/:id", controllers.UpdateProduct)
	router.DELETE("/api/product/:id", controllers.DeleteProduct)
	router.Run(":3000")
}
