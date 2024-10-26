package routes

import (
	controllers "restaurant-golang/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incommingroutes *gin.Engine) {
	incommingroutes.GET("/table", controllers.GetTables())
	incommingroutes.GET("/table/:id", controllers.GetTable())
	incommingroutes.POST("/table", controllers.CreateTable())
	incommingroutes.PATCH("/table/:id", controllers.UpdateTable())

}
