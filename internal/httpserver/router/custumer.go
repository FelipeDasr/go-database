package router

import (
	"go-databases/internal/httpserver/controller"
	"go-databases/internal/service"

	"github.com/gin-gonic/gin"
)

func startCustomerRouter(router *gin.Engine) {
	customerServices := service.NewCustomerServices()
	customerController := controller.NewCustomerController(customerServices)

	customerRouter := router.Group("/customers")
	customerRouter.POST("/", customerController.CreateCustomer);
	customerRouter.GET("/", customerController.GetMany);
	customerRouter.GET("/:id", customerController.GetCustomerById);
}