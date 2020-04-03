package store

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var dbm *sql.DB

func init() {
	var err error
	dbm, err = sql.Open("mysql", "go_user:root@tcp(127.0.0.1:3306)/emali_hh")
	if err != nil {
		panic(err.Error())
	}
}

func InsertOne(query string, args ...interface{}) (int64, error) {
	stmt, err := dbm.Prepare(query)
	defer closeStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(args)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

func Fetch(query string, args ...interface{}) (interface{}, error) {
	//check if there are any args
	//will return a slice of results
	var r *sql.Rows
	var err error
	if len(args) == 0 {
		r, err = dbm.Query(query)
	} else {
		r, err = dbm.Query(query, args)
	}
	for r.Next() {
		r.Scan()
	}
}

func closeRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}

func closeStmt(stmt *sql.Stmt) {
	if stmt != nil {
		stmt.Close()
	}
}
