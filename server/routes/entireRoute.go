package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zack-ali-hassan/Calorie-Tracker-Entire/server/controllers"
)

func SetupRoutes(app *gin.Engine){
	app.GET("/api/entire/", controllers.GetAllEntire)
}