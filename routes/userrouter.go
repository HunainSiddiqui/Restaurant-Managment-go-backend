package routes

import (
	controllers "restaurant-golang/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingroutes *gin.Engine) {
	incommingroutes.GET("/user", controllers.GetUsers())
	incommingroutes.GET("/user/:id", controllers.GetUser())
	incommingroutes.POST("/user/signup", controllers.CreateUser())
	incommingroutes.POST("/user/login", controllers.LoginUser())
	incommingroutes.PUT("/user/:id", controllers.UpdateUser())
	incommingroutes.DELETE("/user/:id", controllers.DeleteUser())

}
