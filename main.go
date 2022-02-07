package main

import (
	"fmt"
	"goapi/database"
	"goapi/pages"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/**
* Handle routing
 */
func initRoutes(router *gin.Engine) {
	// homepage route
	router.GET("/", pages.ShowHomePage)
}

/**
* Main output
 */
func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No local env detected")
	}
	postgresURL := os.Getenv("DATABASE_URL")
	database.InitDB(postgresURL)
	database.Automigrate()

	// setup app
	app := gin.Default()
	// set cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))
	// set to listen
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	// initalize routes
	initRoutes(app)

	// serve
	app.Run(":" + port)

}
