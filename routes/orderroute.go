package routes

import (
	controllers "restaurant-golang/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incommingroutes *gin.Engine) {
	incommingroutes.GET("/order", controllers.GetOrders())
	incommingroutes.GET("/order/:id", controllers.GetOrder())
	incommingroutes.POST("/order", controllers.CreateOrder())
	incommingroutes.PATCH("/order/:id", controllers.UpdateOrder())

}
