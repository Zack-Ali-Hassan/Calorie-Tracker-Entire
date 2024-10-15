package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zack-ali-hassan/Calorie-Tracker-Entire/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func SetCollection(c *mongo.Collection) {
	collection = c
}
func GetAllCalorieTracker(c *gin.Context) {
	calorieTrackerEntire := []models.CalorieTracker{}
	reslut, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error getting Calorie Tracker"})
		return
	}
	defer reslut.Close(context.Background())
	for reslut.Next(context.Background()) {
		calorieTracker := models.CalorieTracker{}
		if err := reslut.Decode(&calorieTracker); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error decoding  Calorie Tracker"})
			return
		}
		calorieTrackerEntire = append(calorieTrackerEntire, calorieTracker)

	}

	c.JSON(http.StatusOK, calorieTrackerEntire)
}
func GetCalorieTracker(c *gin.Context) {

}
func CreateCalorieTracker(c *gin.Context) {

}
func UpdateCalorieTracker(c *gin.Context) {

}
func DeleteCalorieTracker(c *gin.Context) {

}
