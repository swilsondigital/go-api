package mocks

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MockDatabase() (*gorm.DB, sqlmock.Sqlmock) {

	// get db and mock
	sqlDB, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp),
	)
	if err != nil {
		log.Fatalf("[sqlmock new] %s", err)
	}
	defer sqlDB.Close()

	// create dialector
	dialector := mysql.New(mysql.Config{
		Conn:       sqlDB,
		DriverName: "mysql",
	})

	// a SELECT VERSION() query will be run when gorm opens the database
	// so we need to expect that here
	columns := []string{"version"}
	mock.ExpectQuery("SELECT VERSION()").WithArgs().WillReturnRows(
		mock.NewRows(columns).FromCSVString("1"),
	)

	// open the database
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("[gorm open] %s", err)
	}

	return db, mock
}
