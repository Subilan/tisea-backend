package utils

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() (*sql.DB, error) {
	cfg := GetConfiguration()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s:%d/%s", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Dbname))
	if err != nil {
		return nil, fmt.Errorf("Invalid connection arguments.")
	}
	connErr := db.Ping()
	if connErr != nil {
		return nil, fmt.Errorf("Cannot reach the database.")
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}

func Exec(execString string, values... interface{}) (sql.Result, error) {
	db, err := GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, stmtErr := db.Prepare(execString)

	if stmtErr != nil {
		return nil, stmtErr
	}

	effect, execErr := stmt.Exec(values...)
	if execErr != nil {
		return nil, execErr
	}

	return effect, nil
}

func Query(queryString string, values... interface{}) (*sql.Rows, error) {
	db, err := GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, stmtErr := db.Prepare(queryString)

	if stmtErr != nil {
		return nil, stmtErr
	}

	result, queryErr := stmt.Query(values...)
	if queryErr != nil {
		return nil, queryErr
	}

	return result, nil
}