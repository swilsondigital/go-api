package main

import (
	"fmt"
	"goapi/controllers"
	"goapi/database"
	h "goapi/pages"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

/**
* Handle routing
 */
func initRoutes(router *mux.Router, DB *gorm.DB) {
	// homepage route
	router.HandleFunc("/", h.ShowHomePage).Methods("GET")

	// user routes
	u := controllers.UserController{Router: router, DB: DB}
	u.InitializeUserRoutes()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}
	// log and listen/serve
	log.Fatal(http.ListenAndServe(":"+port, router))
}

/**
* Main output
 */
func main() {
	// init message
	fmt.Println("Initializing Rest User API")
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No local env detected")
	}
	postgresURL := os.Getenv("DATABASE_URL")
	DB := database.InitDB(postgresURL)
	router := mux.NewRouter().StrictSlash(true)
	initRoutes(router, DB)
}
