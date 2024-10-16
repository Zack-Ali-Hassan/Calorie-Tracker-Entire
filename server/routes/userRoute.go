package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zack-ali-hassan/Calorie-Tracker-Entire/server/controllers"
)

func SetupUserRoutes(app *gin.Engine) {
	app.GET("/api/user/", controllers.GetAllUsers)
	app.POST("/api/user/", controllers.CreateUser)
	// app.GET("/api/user/:id", controllers.GetCalorieTracker)
	// app.PATCH("/api/user/:id", controllers.UpdateCalorieTracker)
	// app.DELETE("/api/user/:id", controllers.DeleteCalorieTracker)
}
