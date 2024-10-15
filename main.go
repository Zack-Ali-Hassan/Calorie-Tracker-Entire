package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zack-ali-hassan/Calorie-Tracker-Entire/server/controllers"
	"github.com/zack-ali-hassan/Calorie-Tracker-Entire/server/db"
	"github.com/zack-ali-hassan/Calorie-Tracker-Entire/server/routes"
)

func main() {

	err := godotenv.Load(".env")
	if os.Getenv("ENV") != "production" {
		if err != nil {
			log.Fatal("Error loading environment file: ", err)
		}
	}
	db := db.ConnectDB()
	controllers.SetCollection(db.Collection("Entire"))
	app := gin.Default()

	app.RedirectTrailingSlash = false
	// Use Gin CORS middleware with the desired config
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},            // Allow specific origin
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},   // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"}, // Allowed headers
		AllowCredentials: true,                                         // Allow credentials like cookies, etc.
	}))
	app.Use(gin.Logger())
	routes.SetupRoutes(app)
	app.SetTrustedProxies([]string{"127.0.0.1"}) // Set trusted proxies as needed

	if os.Getenv("ENV") == "production" {
		app.Static("/", "./frontend/dist")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "6768"
	}
	log.Fatal(app.Run(":" + port))
}
