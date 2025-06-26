package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/leoneIAguilera/simple_bank/internal/db"
	_ "github.com/lib/pq"
)

const (
	db_url    = "postgresql://root:mypwd@localhost:5432/simple_bank?sslmode=disable"
	db_driver = "postgres"
)

var (
	testDB      *sql.DB
	TestQueries *db.Queries
)

func connectDb(m *testing.M) {
	var err error
	testDB, err = sql.Open(db_driver, db_url)

	if err != nil {
		log.Fatal("cannot connect to database", err)
	}
	TestQueries = db.New(testDB)
	os.Exit(m.Run())
}

func TestMain(m *testing.M) {
	connectDb(m)
}
