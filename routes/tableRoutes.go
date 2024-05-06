package routes

import (
	"restraument-management/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {

	groupRoutes := incomingRoutes.Group("/table")

	groupRoutes.POST("/add", controllers.CreateNote())
	groupRoutes.GET("/:table_id", controllers.GetTableById())
	groupRoutes.GET("", controllers.GetTables())
	groupRoutes.PATCH("/:table_id", controllers.UpdateTable())
}
