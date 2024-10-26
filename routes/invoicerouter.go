package routes

import (
	controllers "restaurant-golang/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incommingroutes *gin.Engine) {
	incommingroutes.GET("/invoice", controllers.GetInvoice())
	incommingroutes.GET("/invoice/:id", controllers.GetInvoice())
	incommingroutes.POST("/invoice", controllers.CreateInvoice())
	incommingroutes.PATCH("/invoice/:id", controllers.UpdateInvoice())

}
