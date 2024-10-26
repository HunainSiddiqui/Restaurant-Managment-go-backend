package routes

import (
	controllers "restaurant-golang/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incommingroutes *gin.Engine) {
	incommingroutes.GET("/menu", controllers.GetMenus())
	incommingroutes.GET("/menu/:id", controllers.GetMenu())
	incommingroutes.POST("/menu", controllers.CreateMenu())
	incommingroutes.PATCH("/menu/:id", controllers.UpdateMenu())

}
