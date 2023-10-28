package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlDB(address string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", address)
	if err != nil {
		return nil, fmt.Errorf("GetPostgresDB - error on connecting to postgres database with address: %v error: %v", address, err)
	}

	// Note: need to set open max open connections to prevent large connection to database
	// the more we have connection the more asynchronous operations get done, but at the cost of memory usage of the database
	// and database's max_connection
	// by setting it to a value, a request will need to wait for a connection, if the connection's already at the limit.
	//
	// right now each instance in UFJT BE is set to 200, but it depends on the applications
	db.SetMaxOpenConns(200)

	// 200 connections, 200 used
	// 1 incoming request (A)
	// waiting for available request
	// 199 used
	// A is then handled

	// Note: Setting connection max life time to be longer makes services to become more responsive to request
	// but at the cost of memory usage, and sometimes a long running connection can becomes broken it it stays too long
	// so choose wisely
	db.SetConnMaxLifetime(10 * time.Minute)

	err = db.Ping()

	return db, err
}
