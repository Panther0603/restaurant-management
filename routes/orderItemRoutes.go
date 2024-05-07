package routes

import (
	"restraument-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {

	groupRoutes := incomingRoutes.Group("/orderItems")

	groupRoutes.POST("/add", controllers.CreateOrderItem())
	groupRoutes.GET("/:orderItems_id", controllers.GetOrderItemById())
	groupRoutes.GET("", controllers.GetOrderItems())
	groupRoutes.PUT("/update", controllers.UpdateOrderItem())
}
