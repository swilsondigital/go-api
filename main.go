package main

import (
	"fmt"
	"goapi/database"
	"goapi/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
		port = "3030"
	}

	// initalize routes
	routes.Setup(app)

	// serve
	app.Run(":" + port)

}
