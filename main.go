package main

import (
	"fmt"
	"github.com/ernestomr87/go-productapi/controllers"
	"github.com/ernestomr87/go-productapi/models"
	"github.com/gin-gonic/gin"
)

var products = []models.Product{
	{
		Id: 1, Name: "Product 1", Stock: 10, Price: 12.5,
	},
	{
		Id: 2, Name: "Product 2", Stock: 12, Price: 2.5,
	},
	{
		Id: 3, Name: "Product 3", Stock: 13, Price: 8.5,
	},
}

func main() {
	r := gin.Default()
	controller := controllers.Init(&products)

	r.GET("/product", controller.ReadProducts)
	r.GET("/product/:id", controller.ReadProductById)
	r.POST("/product", controller.CreateProduct)
	r.PUT("/product/:id", controller.UpdateProductById)
	r.DELETE("/product/:id", controller.DeleteProductById)

	r.Run("localhost:8080")
	fmt.Println("Server is Running!!!")
}
