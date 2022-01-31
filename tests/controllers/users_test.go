package userController

import (
	"database/sql"
	"goapi/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const prod_url = "https://fierce-taiga-67843.herokuapp.com/users/"

type Suite struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
	user []models.User
}

func MockDB(t testing.TB) {
	// init helper
	t.Helper()
	// setup suite
	s := &Suite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	s.db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}
}
