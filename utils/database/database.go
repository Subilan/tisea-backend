package database

import (
	"database/sql"
	"fmt"
	"time"
	"tisea-backend/utils/config"

	_ "github.com/go-sql-driver/mysql"
)

// 使用 config.yml 中指定的数据，尝试从数据库创建连接。如果创建成功，会返回一个 *sql.DB 实例。
func GetConnection() (*sql.DB, error) {
	cfg := config.GetConfiguration()
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

// 执行给定的带参语句，并返回执行的结果。
// 注意：此时的结果是 sql.Result，仅包含了 LastInsertedId 和 RowsAffected 两项。
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

// 执行给定的带参语句，并返回此语句所选择的相关数据表行。
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

// 判断一个语句是否有结果。如果过程中发生错误，返回false
func HasResult(queryString string, values... interface{}) bool {
	db, err := GetConnection()
	if err != nil {
		return false
	}
	defer db.Close()

	stmt, stmtErr := db.Prepare(queryString)

	if stmtErr != nil {
		return false
	}

	result, queryErr := stmt.Query(values...)

	if queryErr != nil {
		return false
	}

	count := 0
	for result.Next() {
		count++
	}

	return count > 0
}