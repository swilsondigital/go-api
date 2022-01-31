package main

import (
	"fmt"
	controllers "goapi/controllers"
	"goapi/database"
	h "goapi/pages"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

/**
* Handle routing
 */
func initRoutes() {
	// new mux router
	router := mux.NewRouter().StrictSlash(true)

	// redirect root / to users
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "./users", http.StatusMovedPermanently)
	// })

	router.HandleFunc("/", h.ShowHomePage).Methods("GET")

	// get random user
	router.HandleFunc("/random-user", controllers.GetRandomUser).Methods("GET")

	// user subrouter paths
	userRouter := router.PathPrefix("/users").Subrouter()
	// create new user
	userRouter.HandleFunc("/", controllers.CreateNewUser).Methods("POST")
	// get all users
	userRouter.HandleFunc("/", controllers.GetAllUsers).Methods("GET")
	// get single user
	userRouter.HandleFunc("/{id}", controllers.GetUser).Methods("GET")
	// update single user
	userRouter.HandleFunc("/{id}", controllers.UpdateUser).Methods("PUT")
	// delete single user
	userRouter.HandleFunc("/{id}", controllers.DeleteUser).Methods("DELETE")
	// get single user
	userRouter.HandleFunc("/{id}/hello", controllers.IntroduceUser).Methods("GET")
	// delete all users
	userRouter.HandleFunc("/", controllers.DeleteAllUsers).Methods("POST", "DELETE")

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}
	// log and listen/serve
	log.Fatal(http.ListenAndServe(":"+port, router))
}

/**
* conver string to time for initial data
 */
// func convertStringToTime(s string) time.Time {
// 	t, _ := time.Parse(time.RFC3339, s)
// 	return t
// }

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
	database.InitDB(postgresURL)
	initRoutes()
}
