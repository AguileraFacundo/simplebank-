package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	db_url    = "postgresql://root:mypwd@localhost:5432/simple_bank?sslmode=disable"
	db_driver = "postgres"
)

var TestQueries *Queries

func TestMain(m *testing.M) {

	conn, err := sql.Open(db_driver, db_url)

	if err != nil {
		log.Fatal("cannot connect to database", err)
	}
	TestQueries = New(conn)
	os.Exit(m.Run())
}
