package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zack-ali-hassan/Calorie-Tracker-Entire/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	calorieTrackerEntire := new(models.CalorieTracker)
	if err := c.ShouldBindJSON(calorieTrackerEntire); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error parsing Calorie Tracker"})
		return
	}
	if calorieTrackerEntire.Dish == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please provide the dish name.."})
		return
	}
	if calorieTrackerEntire.Fat == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please provide the fat content.."})
		return
	}
	if calorieTrackerEntire.Ingredients == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please provide the ingredients.."})
		return
	}
	if calorieTrackerEntire.Calories == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please provide the Calories.."})
		return
	}
	insertData, err := collection.InsertOne(context.Background(), calorieTrackerEntire)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error inserted Calorie Tracker"})
		return
	}
	calorieTrackerEntire.ID = insertData.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusOK, gin.H{"msg": "Inseted successfully.."})

}
func UpdateCalorieTracker(c *gin.Context) {

}
func DeleteCalorieTracker(c *gin.Context) {
	id := c.Params.ByName("id")
	deleted_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid calorie tracking.."})
		return
	}
	filter := bson.M{"_id": deleted_id}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error deleting calorie tracking"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Inseted successfully.."})
}
