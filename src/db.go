package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// type User struct {
// 	// this should be auto incremented so probably based on the db
// 	ID              int64     `json:"id" validate:"unique=ID"`
// 	FirstName       string    `json:"fname"`
// 	LastName        string    `json:"lname"`
// 	PreferredName   string    `json:"pname"`
// 	Email           string    `json:"email" validate:"required,email"`           // required email
// 	Skillset        []string  `json:"skills" validate:"gt=0,dive,dive,required"` // greater than 0 [] & string is required
// 	YearsExperience int       `json:"experience" validate:"gt=0,lt=130"`         // between 0-130
// 	MemberSince     time.Time `json:"since" validate:"lte"`                      // less than today
// }

const pgURL = "postgresql://postgres@localhost:5432/samplepostgres"

/**
* Initialize DB tables
 */
func InitDB() {
	// setup connection to db
	// db, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, pgURL)
	if err != nil {
		panic(err)
	}

	// defer closing until after query is run
	defer db.Close()

	// setup database transaction
	tx, err := db.Begin(ctx)
	if err != nil {
		// error on transaction start
		panic(err)
	}

	// defer transaction finishing
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	// setup query for easier management
	createTableQuery := "CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, fname VARCHAR(50), lname VARCHAR(50), pname VARCHAR(50), email VARCHAR(50), skills TEXT[], experience INT, since TIMESTAMP)"

	// run query and panic on error
	if _, err := db.Exec(ctx, createTableQuery); err != nil {
		panic(err)
	}

	// success message in console
	fmt.Println("Successfully created users table")
}

/**
* Wrapper func for running insert/update/delete db queries
 */
func RunQuery(query string) bool {
	// setup connection to db
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		// error on connection
		panic(err)
	}

	// defer closing until after query is run
	defer db.Close()

	// setup database transaction
	tx, err := db.Begin(ctx)
	if err != nil {
		// error on transaction start
		panic(err)
	}

	// defer transaction finishing
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	// sql run query
	if _, err := db.Exec(ctx, query); err != nil {
		panic(err)
	}

	// return bool for confirmation
	if err != nil {
		return false
	} else {
		return true
	}
}

/**
* wrapper func for running select queries
 */
func SelectQuery(query string) {
	// setup connection to db
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		// error on connection
		panic(err)
	}

	// defer closing until after query is run
	defer db.Close()

	// run query for rows
	rows, err := db.Query(ctx, query)
	if err != nil {
		panic(err)
	}
	// defer closing row data
	defer rows.Close()

}
