package controller

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"strconv"

	"ruibao_pos_system/models"
	"ruibao_pos_system/service"
)

func GetAllProducts(c *gin.Context) {
	products, err := service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	parsedID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	productID := uint(parsedID)
	
	product, err := service.GetProductByID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func GetProductByBarcode(c *gin.Context) {
	barcode := c.Param("barcode")
	product, err := service.GetProductByBarcode(barcode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var input struct {
		Name     string  `json:"name" binding:"required"`
		Barcode  string  `json:"barcode" binding:"required"`
		Price    float64 `json:"price" binding:"required"`
		Stock    int     `json:"stock" binding:"required"`
		Category string  `json:"category" binding:"required"`
	}

	product, err := service.CreateProduct(input.Name, input.Barcode, input.Price, input.Stock, input.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	parsedID, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var input struct {
		ID       uint    `json:"id" binding:"required"`
		Name     string  `json:"name"`
		Barcode  string  `json:"barcode"`
		Price    float64 `json:"price"`
		Stock    int     `json:"stock"`
		Category string  `json:"category"`
	}

	product := &models.Product{
		ID:       uint(parsedID),
		Name:     input.Name,
		Barcode:  input.Barcode,
		Price:    input.Price,
		Stock:    input.Stock,
		Category: input.Category,
	}	

	err = service.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	
	err = service.DeleteProduct(uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
