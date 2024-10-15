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
	calorieTrackerEntire := models.CalorieTracker{}
	id := c.Params.ByName("id")
	object_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid Id, please try again?"})
		return
	}
	filter := bson.M{"_id": object_id}
	err = collection.FindOne(context.Background(), filter).Decode(&calorieTrackerEntire)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error fetching calories tracker"})
		return
	}
	c.JSON(http.StatusOK, calorieTrackerEntire)
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
	calorieTrackerEntire := new(models.CalorieTracker)
	id := c.Params.ByName("id")
	updated_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid Id, please try again?"})
		return
	}
	if err := c.ShouldBindJSON(&calorieTrackerEntire); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Error parsing Calorie Tracker"})
		return
	}
	update_fields := bson.M{}
	if calorieTrackerEntire.Dish != "" {
		update_fields["dish"] = calorieTrackerEntire.Dish
	}
	if calorieTrackerEntire.Fat != 0 {
		update_fields["fat"] = calorieTrackerEntire.Fat
	}
	if calorieTrackerEntire.Ingredients != "" {
		update_fields["ingredients"] = calorieTrackerEntire.Ingredients
	}
	if calorieTrackerEntire.Calories != 0 {
		update_fields["calories"] = calorieTrackerEntire.Calories
	}
	filter := bson.M{"_id": updated_id}
	update := bson.M{"$set": update_fields}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error updating calorie tracking"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Updated successfully.."})
}
func DeleteCalorieTracker(c *gin.Context) {
	id := c.Params.ByName("id")
	deleted_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid Id, please try again?"})
		return
	}
	filter := bson.M{"_id": deleted_id}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error deleting calorie tracking"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Deleted successfully.."})
}
