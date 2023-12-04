package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Pool *sql.DB

// 执行给定的带参语句，并返回执行的结果。
// 注意：此时的结果是 sql.Result，仅包含了 LastInsertedId 和 RowsAffected 两项。
func Exec(execString string, values ...interface{}) (sql.Result, error) {

	stmt, stmtErr := Pool.Prepare(execString)

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
func Query(queryString string, values ...interface{}) (*sql.Rows, error) {
	stmt, stmtErr := Pool.Prepare(queryString)

	if stmtErr != nil {
		return nil, stmtErr
	}

	result, queryErr := stmt.Query(values...)
	if queryErr != nil {
		return nil, queryErr
	}

	return result, nil
}

// 获取一个语句结果的行数。如果过程中发生错误，返回 -1；正常结果 >=0
func Count(queryString string, values ...interface{}) int {
	var count int
	err := Pool.QueryRow(queryString, values...).Scan(&count)

	if err != nil {
		return -1
	}

	return count
}

// 判断一个语句是否有结果。
func HasResult(queryString string, values ...interface{}) bool {
	return Count(queryString, values...) > 0
}
