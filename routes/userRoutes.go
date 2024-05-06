package routes

import (
	"restraument-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	groupRoutes := incomingRoutes.Group("/user")

	groupRoutes.POST("/signup", controllers.Signup())
	groupRoutes.POST("/login", controllers.Login())
	groupRoutes.GET("/:user_id", controllers.GetUserById())
	groupRoutes.GET("", controllers.GetUsers())
}
