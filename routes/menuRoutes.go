package routes

import (
	"restraument-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {

	groupRoutes := incomingRoutes.Group("/menu")

	groupRoutes.POST("/add", controllers.CreateMenu())
	groupRoutes.GET("/:menuId", controllers.GetMenuById())
	groupRoutes.GET("", controllers.GetMenus())
	groupRoutes.PUT("/", controllers.UpdateMenu())
	//groupRoutes.DELETE()

}
