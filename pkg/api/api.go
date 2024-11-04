package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yulaymusin/go-ecommerce/internal/services"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// User routes
	userSvc := services.NewUserService()
	router.POST("/users", userSvc.CreateUser)
	router.POST("/auth", userSvc.Authenticate)

	// Product routes
	productSvc := services.NewProductService()
	router.GET("/products", productSvc.ListProducts)
	router.GET("/products/:id", productSvc.GetProduct)
	router.POST("/products", productSvc.CreateProduct)
	router.PUT("/products/:id", productSvc.UpdateProduct)
	router.DELETE("/products/:id", productSvc.DeleteProduct)

	// Order routes
	orderSvc := services.NewOrderService()
	router.POST("/orders", orderSvc.CreateOrder)

	return router
}
