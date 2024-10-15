package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zack-ali-hassan/Calorie-Tracker-Entire/server/controllers"
)

func SetupRoutes(app *gin.Engine) {
	app.GET("/api/entire/", controllers.GetAllCalorieTracker)
	app.POST("/api/entire/", controllers.CreateCalorieTracker)
	app.GET("/api/entire/:id", controllers.GetCalorieTracker)
	app.PATCH("/api/entire/:id", controllers.UpdateCalorieTracker)
	app.DELETE("/api/entire/:id", controllers.DeleteCalorieTracker)
}
