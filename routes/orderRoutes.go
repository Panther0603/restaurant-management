package routes

import (
	"restraument-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	groupRoutes := incomingRoutes.Group("/order")

	groupRoutes.POST("/add", controllers.CreateOrder())
	groupRoutes.GET("/:order_id", controllers.GetOrderById())
	groupRoutes.GET("", controllers.GetOrders())
	groupRoutes.PUT("/update", controllers.UpdateOrder())
	//groupRoutes.DELETE()

}
