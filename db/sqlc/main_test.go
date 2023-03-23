package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://rizkia:21204444@localhost:5432/bookskette?sslmode=disable" // copy saja dari migrate command
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("tidak bisa tersambung ke database : ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
