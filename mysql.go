package store

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// MySQLMaxOpenConns will be used to set a MySQL drivers MaxOpenConns
	MySQLMaxOpenConns = 10
	// MySQLMaxIndleConns will be used to set a MySQL drivers MaxIndleConns
	MySQLMaxIndleConns = 10
)

// MySQLConn will read env and open a db connection
func MySQLConn() (*sql.DB, error) {
	dns := os.Getenv("MYSQL_DSN")
	db, err := sql.Open("mysql", dns)

	if err != nil {
		return db, err
	}

	err = db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
	if err != nil {
		return db, err
	}

	db.SetMaxOpenConns(MySQLMaxIndleConns)
	db.SetMaxIdleConns(MySQLMaxIndleConns)
	return db, err
}
