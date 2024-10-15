package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func SetCollection(c *mongo.Collection) {
	collection = c
}
func GetAllCalorieTracker(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Get all calorie"})
}
func GetCalorieTracker(c *gin.Context) {

}
func CreateCalorieTracker(c *gin.Context) {

}
func UpdateCalorieTracker(c *gin.Context) {

}
func DeleteCalorieTracker(c *gin.Context) {

}
