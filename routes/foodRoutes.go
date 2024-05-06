package routes

import (
	"restraument-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {

	groupRoutes := incomingRoutes.Group("/food")

	groupRoutes.POST("/add", controllers.CreateFood())
	groupRoutes.GET("/:food_id", controllers.GetFoodById())
	groupRoutes.GET("", controllers.GetFoods())
	groupRoutes.PATCH("/:food_id", controllers.UpdateFood())
	//groupRoutes.DELETE()

}
