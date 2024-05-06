package routes

import (
	"restraument-management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {

	groupRoutes := incomingRoutes.Group("/invoice")

	groupRoutes.POST("/add", controllers.CreateInvoice())
	groupRoutes.GET("/:invoice_id", controllers.GetInvoiceById())
	groupRoutes.GET("", controllers.GetInvoices())
	groupRoutes.PATCH("/:invoice_id", controllers.UpdateInvoice())
	//groupRoutes.DELETE()

}
