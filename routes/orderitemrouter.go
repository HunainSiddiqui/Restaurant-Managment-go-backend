package routes

import (
	controllers "restaurant-golang/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incommingroutes *gin.Engine) {
	incommingroutes.GET("/orderitem", controllers.GetOrderItems())
	incommingroutes.GET("/orderitem/:id", controllers.GetOrderItem())
	incommingroutes.GET("/orderitems/:order_id", controllers.GetOrderItemsByOrder())
	incommingroutes.POST("/orderitem", controllers.CreateOrderItem())
	incommingroutes.PATCH("/orderitem/:id", controllers.UpdateOrderItem())

}
