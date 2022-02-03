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
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})

	handler := c.Handler(router)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}
	// log and listen/serve
	log.Fatal(http.ListenAndServe(":"+port, handler))
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
