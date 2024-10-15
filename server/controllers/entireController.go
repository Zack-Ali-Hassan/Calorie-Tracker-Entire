package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)
var collection *mongo.Collection
func SetCollection(c *mongo.Collection){
	collection =c
}
func GetAllEntire(c *gin.Context){

}