package routes

import (
	controllers "restraument-management/controllers"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(incomingRoutes *gin.Engine) {

	groupRoutes := incomingRoutes.Group("/note")

	groupRoutes.POST("/add", controllers.CreateNote())
	groupRoutes.GET("/:note_id", controllers.GetNoteById())
	groupRoutes.GET("", controllers.GetNotes())
	groupRoutes.PATCH("/:note_id", controllers.UpdateNote())

}
