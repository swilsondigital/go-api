package main

import (
	"fmt"
	u "goapi/src/goapi/users"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/**
* Handle routing
 */
func initRoutes() {
	// new mux router
	router := mux.NewRouter().StrictSlash(true)

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

	// log and listen/serve
	log.Fatal(http.ListenAndServe(":10000", router))
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
	// initialize users for instant data
	u.Users = []u.User{
		{
			ID:              1,
			FirstName:       "John",
			LastName:        "Smith",
			PreferredName:   "JJ",
			Email:           "john.smith@sourcestrike.com",
			Skillset:        []string{"PHP", "GoLang", "JS"},
			YearsExperience: 15,
			MemberSince:     convertStringToTime("2001-02-22T00:00:00Z"),
		},
		{
			ID:              2,
			FirstName:       "Jeff",
			LastName:        "Goldblum",
			PreferredName:   "",
			Email:           "jeff.goldblum@sourcestrike.com",
			Skillset:        []string{"MVC", "C#", "F"},
			YearsExperience: 32,
			MemberSince:     convertStringToTime("1984-08-30T00:00:00Z"),
		},
		{
			ID:              3,
			FirstName:       "Jack",
			LastName:        "Johnson",
			PreferredName:   "",
			Email:           "jack.johnson@sourcestrike.com",
			Skillset:        []string{"Python", "NodeJS"},
			YearsExperience: 4,
			MemberSince:     convertStringToTime("2020-01-01T00:00:00Z"),
		},
		{
			ID:              4,
			FirstName:       "Sebastian",
			LastName:        "Cumberbundt",
			PreferredName:   "Sebas",
			Email:           "sebastian.cumberbundt@sourcestrike.com",
			Skillset:        []string{"Things", "Stuff", "JS"},
			YearsExperience: 15,
			MemberSince:     convertStringToTime("2018-07-19T00:00:00Z"),
		},
	}
	initRoutes()
}
