package routes

import (
	controllers "restaurant-golang/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incommingroutes *gin.Engine) {
	incommingroutes.GET("/food", controllers.GetFoods())
	incommingroutes.GET("/food/:id", controllers.GetFood())
	incommingroutes.POST("/food", controllers.CreateFood())
	incommingroutes.PATCH("/food/:id", controllers.UpdateFood())

}
