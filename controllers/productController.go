package controllers

import (
	"annanta/backend-api/config"
	"annanta/backend-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type ProductModel struct {
// 	product *models.Product
// }

// get all data
func GetAllProduct(c *gin.Context) {

	var products []models.Product

	config.DB.Find(&products)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    products,
	})

}

//create data

type ValidateProduct struct {
	Name  string `json:"name" binding:"required"`
	Price string `json:"price" binding:"required"`
}

func CreateProduct(c *gin.Context) {
	var input ValidateProduct

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var product models.Product

	product.Name = input.Name
	product.Price = input.Price

	config.DB.Create(&product)

	c.JSON(201, gin.H{
		"message": "success",
		"data":    product,
	})

}

// func (p *ProductModel) GetProductById(c *gin.Context) {
// 	// var product models.Product

// 	if err := config.DB.Where("id = ?", c.Param("id")).First(&p.product); err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": "Not Found",
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"message": "success",
// 		"data":    p.product,
// 	})

// }

func GetProductById(c *gin.Context) {
	var product models.Product

	result := config.DB.Where("id = ?", c.Param("id")).First(&product)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    product,
	})

}

func UpdateProduct(c *gin.Context) {

	var input ValidateProduct

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var product models.Product

	result := config.DB.Where("id = ?", c.Param("id")).First(&product)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
		return
	}
	config.DB.Model(&product).Updates(input)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    product,
	})
}

func DeleteProduct(c *gin.Context) {
	var product models.Product

	result := config.DB.Where("id = ?", c.Param("id")).First(&product)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
		return
	}

	config.DB.Delete(&product)

	c.JSON(200, gin.H{
		"message": "success",
	})
}
