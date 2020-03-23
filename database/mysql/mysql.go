package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

var (
	// DB is the connection handle
	DB *sql.DB
)

// Open opens database connection
func Open(user, password, host, port, name, timezone, charset string, maxOpenConns, maxIdleConns int, maxLifetime time.Duration) error {
	db, err := sql.Open(
		"mysql",
		user+":"+password+"@tcp("+host+":"+port+")"+"/"+name+
			"?parseTime=true&loc="+timezone+"&charset="+charset)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(maxLifetime * time.Minute)

	DB = db

	return nil
}

// prepareQuery prepares query
func prepareQuery(query string) (*sql.Stmt, error) {
	statement, err := DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	return statement, nil
}

// executeQuery executes request of type INSERT, UPDATE or DELETE
func executeQuery(query string, args ...interface{}) (sql.Result, error) {
	statement, err := prepareQuery(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	result, err := statement.Exec(args...)
	if err != nil {
		return nil, err
	}

	return result, err
}

// Select request
func Select(query string, args ...interface{}) (*sql.Rows, error) {
	statement, err := prepareQuery(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query(args...)
	if err != nil {
		return nil, err
	}

	return rows, err
}

// Insert request
func Insert(query string, args ...interface{}) (int64, error) {
	result, err := executeQuery(query, args...)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, err
}

// Update request
func Update(query string, args ...interface{}) (int64, error) {
	result, err := executeQuery(query, args...)
	if err != nil {
		return 0, err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affect, err
}

// Delete request
func Delete(query string, args ...interface{}) (int64, error) {
	result, err := executeQuery(query, args...)
	if err != nil {
		return 0, err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affect, err
}
