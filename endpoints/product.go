package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"neurothrone/go-db-web-api/data"
	"neurothrone/go-db-web-api/models"
	"neurothrone/go-db-web-api/views"
	"strconv"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html", &views.PageView{Title: "Welcome!", Text: "There is no turning back now."})
}

func AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		return
	}

	product.Id = 0
	data.CreateNewProduct(&product)
	c.IndentedJSON(http.StatusCreated, product)
}

func GetAllProducts(c *gin.Context) {
	products := data.GetAllProducts()
	c.IndentedJSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product := data.GetProduct(productId)

	if product == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Product by id: %d not found", productId)})
	} else {
		c.IndentedJSON(http.StatusOK, product)
	}
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	existingProduct := data.GetProduct(productId)
	if existingProduct == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Product by id: %d not found", productId)})
		return
	}

	var productToUpdate models.Product
	if err := c.BindJSON(&productToUpdate); err != nil {
		return
	}

	productToUpdate.Id = productId
	data.UpdateProduct(productToUpdate)
	c.IndentedJSON(http.StatusOK, productToUpdate)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	existingProduct := data.GetProduct(productId)
	if existingProduct == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Product by id: %d not found", productId)})
		return
	}

	data.DeleteProduct(existingProduct)
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Product deleted"})
}
