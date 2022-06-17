package controllers

import (
	"net/http"
	"strconv"

	"github.com/ernestomr87/go-productapi/models"
	"github.com/gin-gonic/gin"
)

type ProductRepo struct {
	Products *[]models.Product
}

func Init(products *[]models.Product) *ProductRepo {
	return &ProductRepo{Products: products}
}

// CRUD OPERATIONS
func (r *ProductRepo) CreateProduct(c *gin.Context) {
	var product models.Product

	c.BindJSON(&product)
	err := models.CreateProduct(r.Products, &product)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (r *ProductRepo) ReadProducts(c *gin.Context) {
	c.JSON(http.StatusOK, r.Products)
}

func (r *ProductRepo) ReadProductById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)

	product := models.ReadProductById(r.Products, int(id))
	c.JSON(http.StatusOK, product)
}

func (r *ProductRepo) UpdateProductById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)

	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	var product models.Product
	c.BindJSON(&product)

	if product.Id != int(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	updateProduct := models.UpdateProductById(r.Products, &product)

	c.JSON(http.StatusOK, updateProduct)

}

func (r *ProductRepo) DeleteProductById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)

	product := models.DeleteProductById(r.Products, int(id))

	c.JSON(http.StatusOK, product)
}
