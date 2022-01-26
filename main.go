package main

import (
	"fmt"
	db "goapi/src"
	h "goapi/src/pages"
	u "goapi/src/users"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
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
	router.HandleFunc("/random-user", u.GetRandomUser).Methods("GET")

	// user subrouter paths
	userRouter := router.PathPrefix("/users").Subrouter()
	// get all users
	userRouter.HandleFunc("/", u.GetAllUsers).Methods("GET")
	// create new user
	userRouter.HandleFunc("/", u.CreateNewUser).Methods("POST")
	// get single user
	userRouter.HandleFunc("/{id}", u.GetUser).Methods("GET")
	// update single user
	userRouter.HandleFunc("/{id}", u.UpdateUser).Methods("PUT")
	// delete single user
	userRouter.HandleFunc("/{id}", u.DeleteUser).Methods("DELETE")
	// get single user
	userRouter.HandleFunc("/{id}/hello", u.IntroduceUser).Methods("GET")
	// delete all users
	userRouter.HandleFunc("/", u.DeleteAllUsers).Methods("POST", "DELETE")

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
func convertStringToTime(s string) time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return t
}

/**
* Main output
 */
func main() {
	// init message
	fmt.Println("Rest User API")
	// create db connection
	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	panic(err)
	// }
	// initialize users for instant data
	u.Users = []u.User{
		{
			ID:              1,
			FirstName:       "John",
			LastName:        "Smith",
			PreferredName:   "JJ",
			Email:           "john.smith@example.com",
			Skillset:        []string{"PHP", "GoLang", "JS"},
			YearsExperience: 15,
			MemberSince:     convertStringToTime("2001-02-22T00:00:00Z"),
		},
		{
			ID:              2,
			FirstName:       "Jeff",
			LastName:        "Goldblum",
			PreferredName:   "",
			Email:           "jeff.goldblum@example.com",
			Skillset:        []string{"MVC", "C#", "F"},
			YearsExperience: 32,
			MemberSince:     convertStringToTime("1984-08-30T00:00:00Z"),
		},
		{
			ID:              3,
			FirstName:       "Jack",
			LastName:        "Johnson",
			PreferredName:   "",
			Email:           "jack.johnson@example.com",
			Skillset:        []string{"Python", "NodeJS"},
			YearsExperience: 4,
			MemberSince:     convertStringToTime("2020-01-01T00:00:00Z"),
		},
		{
			ID:              4,
			FirstName:       "Sebastian",
			LastName:        "Cumberbundt",
			PreferredName:   "Sebas",
			Email:           "sebastian.cumberbundt@example.com",
			Skillset:        []string{"Things", "Stuff", "JS"},
			YearsExperience: 15,
			MemberSince:     convertStringToTime("2018-07-19T00:00:00Z"),
		},
	}
	db.InitDB()
	initRoutes()
}
