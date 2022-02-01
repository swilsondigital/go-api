package userController

import (
	"fmt"
	"goapi/controllers"
	"goapi/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	mocket "github.com/selvatico/go-mocket"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const prod_users_url = "https://fierce-taiga-67843.herokuapp.com/users/"

func SetupTests() (c *controllers.UserController, err error) {
	// start mocket
	mocket.Catcher.Register()
	mocket.Catcher.Logging = false
	// setup dialector
	dialector := postgres.New(postgres.Config{
		DSN:                  "mockdb",
		DriverName:           mocket.DriverName,
		PreferSimpleProtocol: true,
	})
	// open gorm
	db, err := gorm.Open(dialector, &gorm.Config{})
	// check if open failed
	if err != nil {
		err = fmt.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}

	// check if db set
	if db == nil {
		err = fmt.Errorf("gorm db is nil")
	}

	c = &controllers.UserController{Router: mux.NewRouter(), DB: db}

	return
}

func parseUserTime(s string) time.Time {
	t, _ := time.Parse(s, s)
	return t
}

func seedUsers(db *gorm.DB) error {
	users := []models.User{
		{
			FirstName:       "Anthony",
			LastName:        "Stark",
			PreferredName:   "Tony",
			Email:           "tony@starkindustries.com",
			Skillset:        `["Building Robotic Suits", "Philanthropy", "Saving the World"]`,
			YearsExperience: 15,
			MemberSince:     parseUserTime("2001-02-22"),
		},
		{
			FirstName:       "Peter",
			LastName:        "Parker",
			PreferredName:   "Spiderman",
			Email:           "peter@thedailybugle.net",
			Skillset:        `["Photography", "Gene Splicing", "Web Development"]`,
			YearsExperience: 9,
			MemberSince:     parseUserTime("2004-07-30"),
		},
		{
			FirstName:       "Aegon",
			LastName:        "Targaryen",
			PreferredName:   "Jon Snow",
			Email:           "jonsnow@nightswatch.org",
			Skillset:        `["Brooding", "Direwolves", "Knowing Nothing"]`,
			YearsExperience: 24,
			MemberSince:     parseUserTime("1989-12-02"),
		},
	}

	for _, u := range users {
		if err := db.Create(&u).Error; err != nil {
			return err
		}
	}
	return nil
}

func executeRequest(req *http.Request, uc *controllers.UserController) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	uc.Router.ServeHTTP(rr, req)

	return rr
}

/**
* Test for get all users endpoint
 */
func TestGetAllUsers(t *testing.T) {
	// get new instance of db
	uc, err := SetupTests()
	if err != nil {
		t.Errorf(err.Error())
	}

	// initialize dummy request
	r, err := http.NewRequest(http.MethodGet, prod_users_url, nil)
	if err != nil {
		t.Fatal(err)
	}

	if err := seedUsers(uc.DB); err != nil {
		t.Fatal(err)
	}

	fmt.Println(r)

	// run request
	// response := executeRequest(r, uc)

	var users []models.User
	if err := uc.DB.Find(&users).Error; err != nil {
		t.Fatal(err)
	}

	fmt.Println(users)

	// // run tests
	// t.Run("Return Empty", func(t *testing.T) {
	// 	// run request
	// 	uc.GetAllUsers(rr, r)
	// 	// get result
	// 	rs := rr.Result()
	// 	if rs.StatusCode != http.StatusOK {
	// 		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	// 	}
	// 	// defer closing result body
	// 	defer rs.Body.Close()
	// 	body, err := ioutil.ReadAll(rs.Body)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	if string(body) != "[]" {
	// 		t.Errorf("want body to equal %q. got %q", "[]", string(body))
	// 	}
	// })

	// run tests
	// t.Run("Return Users", func(t *testing.T) {
	// 	if err := seedUsers(uc.DB); err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	// run request
	// 	uc.GetAllUsers(rr, r)
	// 	// get result
	// 	rs := rr.Result()
	// 	if rs.StatusCode != http.StatusOK {
	// 		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	// 	}
	// 	// defer closing result body
	// 	defer rs.Body.Close()
	// 	body, err := io.ReadAll(rs.Body)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	actual := []models.User{}
	// 	if err := json.Unmarshal(body, &actual); err != nil {
	// 		t.Error(err)
	// 	}

	// 	fmt.Println(actual)
	// 	if string(body) != "[]" {
	// 		t.Errorf("want body to equal %q. got %q", "[]", body)
	// 	}
	// })
}
